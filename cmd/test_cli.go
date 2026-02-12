package main

import (
	"bytes"
	"fmt"
	"gin/config"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type Column struct {
	Name     string
	DataType string
	Nullable bool
	Comment  string
}

type Import struct {
	pkgs map[string]struct{}
}

func NewImport() *Import {
	return &Import{
		pkgs: make(map[string]struct{}),
	}
}

func (m *Import) Add(pkg string) {
	if pkg != "" {
		m.pkgs[pkg] = struct{}{}
	}
}

func (m *Import) Render() string {
	if len(m.pkgs) == 0 {
		return ""
	}

	var list []string
	for p := range m.pkgs {
		list = append(list, fmt.Sprintf("\t%q", p))
	}
	sort.Strings(list)

	return "import (\n" + strings.Join(list, "\n") + "\n)\n"
}

func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func snakeToLowerCamel(s string) string {
	camel := snakeToCamel(s)
	return strings.ToLower(camel[:1]) + camel[1:]
}

func goType(c Column, im *Import) string {
	switch c.DataType {
	case "tinyint", "smallint", "mediumint", "int", "bigint":
		return "int64"
	case "varchar", "text", "longtext":
		return "string"
	case "json":
		im.Add("gin/app/model")
		return "*model.JsonValue"
	case "datetime", "timestamp":
		im.Add("time")
		return "*time.Time"
	}
	return "string"
}

func buildGormTag(c Column) string {
	var tags []string
	tags = append(tags, "column:"+c.Name)
	if c.Comment != "" {
		tags = append(tags, "comment:"+c.Comment)
	}
	return "gorm:\"" + strings.Join(tags, ";") + "\""
}

func generateModel(db *gorm.DB, table string, outDir string) error {
	cols, err := db.Migrator().ColumnTypes(table)
	if err != nil {
		return err
	}

	im := NewImport()

	var columns []Column
	for _, col := range cols {
		name := col.Name()
		dt := strings.ToLower(col.DatabaseTypeName())
		nullable, _ := col.Nullable()
		comment, _ := col.Comment()

		columns = append(columns, Column{
			Name:     name,
			DataType: dt,
			Nullable: nullable,
			Comment:  comment,
		})
	}

	structName := snakeToCamel(table)
	tableConst := "TableName" + structName

	// 计算最大长度实现对齐
	maxNameLen := 0
	maxTypeLen := 0
	fieldLines := []string{}

	for _, c := range columns {
		name := snakeToCamel(c.Name)
		typ := goType(c, im)
		if len(name) > maxNameLen {
			maxNameLen = len(name)
		}
		if len(typ) > maxTypeLen {
			maxTypeLen = len(typ)
		}
	}

	for _, c := range columns {
		fieldName := snakeToCamel(c.Name)
		fieldType := goType(c, im)
		jsonName := snakeToLowerCamel(c.Name)
		tag := fmt.Sprintf(
			"`%s json:\"%s\" form:\"%s\"`",
			buildGormTag(c),
			jsonName,
			jsonName,
		)
		line := fmt.Sprintf("%-*s %-*s %s", maxNameLen, fieldName, maxTypeLen, fieldType, tag)
		fieldLines = append(fieldLines, line)
	}

	tpl := template.Must(template.New("model").Parse(modelTpl))

	data := struct {
		Imports    string
		Struct     string
		Table      string
		TableConst string
		Fields     []string
	}{
		Imports:    im.Render(),
		Struct:     structName,
		Table:      table,
		TableConst: tableConst,
		Fields:     fieldLines,
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, data); err != nil {
		return err
	}

	err = os.MkdirAll(outDir, 0755)
	if err != nil {
		return err
	}

	file := filepath.Join(outDir, table+".go")
	return os.WriteFile(file, buf.Bytes(), 0644)
}

const modelTpl = `package model

{{ .Imports }}
const {{ .TableConst }} = "{{ .Table }}"

type {{ .Struct }} struct {
{{- range .Fields }}
	{{ . }}
{{- end }}
}

func (*{{ .Struct }}) TableName() string {
	return {{ .TableConst }}
}
`

func main() {
	err := generateModel(config.Db{}.GetDB(), "user", "./models")
	if err != nil {
		color.Red("Generate Model Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Model generated success!")
}
