package make

import (
	"fmt"
	"gin/common/base"
	"gin/config"
	"gin/utils"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type MakeModel struct {
	base.BaseCommand
}

func (m *MakeModel) Name() string {
	return "make:model"
}

func (m *MakeModel) Description() string {
	return "模型创建"
}

func (m *MakeModel) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			"-t, --table",
			"表名, 如: user 或 user,menu",
			true,
		},
		{
			"-p, --path",
			"输出目录, 如: app/model",
			false,
		},
		{
			"-c, --camel",
			"是否驼峰字段, 如: true",
			false,
		},
	}
}

func (m *MakeModel) Execute(args []string) {
	var (
		fs    = pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
		table = fs.StringP("table", "t", "", "表名, 如: user 或 user,menu")
		path  = fs.StringP("path", "p", "", "输出目录, 如: app/model")
		camel = fs.BoolP("camel", "c", true, "是否驼峰字段, 如: true")
	)

	if err := fs.Parse(args); err != nil {
		color.Red("参数解析失败: %s", err.Error())
	}
	// 去除前斜杠
	p := filepath.Join("app/model/", strings.TrimPrefix(*path, "/"))
	tables := strings.Split(*table, ",")
	for i := range tables {
		tables[i] = strings.TrimSpace(tables[i])
		color.Green("✅  创建模型: %s (表名: %s 是否使用驼峰: %v)\n", p+"/"+tables[i]+".gen.go", tables[i], *camel)
	}

	m.generateFiles(p, tables, *camel)
}

func init() {
	cli.Register(&MakeModel{})
}

// generateFiles 生成模型文件
func (m *MakeModel) generateFiles(path string, tables []string, camel bool) {
	var (
		root    = utils.GetRootPath()
		pkg     = filepath.Base(path)
		outPath = filepath.Join(root + "/app/temp")
	)

	config.InitConfig()
	g := gen.NewGenerator(gen.Config{
		OutPath:           outPath,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
		ModelPkgPath:      path,
	})

	g.UseDB(config.InitMysql())

	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"json": func(detailType gorm.ColumnType) (dataType string) {
			if pkg != "model" {
				return "*model.ArrayString"
			} else {
				return "*ArrayString"
			}
		},
		"datetime": func(detailType gorm.ColumnType) (dataType string) {
			// 针对 deleted_at 字段特殊处理
			if detailType.Name() == "deleted_at" {
				if pkg != "model" {
					return "*model.DeletedAt"
				} else {
					return "*DeletedAt"
				}
			}

			if pkg != "model" {
				return "*model.DateTime"
			} else {
				return "*DateTime"
			}
		},
		// "timestamp":  func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 timestamp 转换为 string
		// "date":       func(detailType gorm.ColumnType) (dataType string) { return "string" }, // 添加此行将 date 转换为 string
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义JSON tag
	if camel {
		g.WithJSONTagNameStrategy(func(columnName string) string {
			return utils.SnakeToLowerCamel(columnName)
		})
	}

	color.Cyan("🚀 开始生成模型, 共 %d 张表", len(tables))

	for _, table := range tables {
		color.Yellow("→ 正在生成表: %s", table)

		model := g.GenerateModel(table)
		g.ApplyBasic(model)
	}

	g.Execute()

	// 自动追加 swaggerignore:"true"
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		filePath := filepath.Join(path, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}
		text := string(content)

		re := regexp.MustCompile("(`[^`]*json:\"deletedAt\"[^`]*`)")

		text = re.ReplaceAllStringFunc(text, func(match string) string {
			if strings.Contains(match, "swaggerignore") {
				return match
			}
			return strings.TrimSuffix(match, "`") + " swaggerignore:\"true\"`"
		})

		if err = os.WriteFile(filePath, []byte(text), 0644); err != nil {
			color.Red(fmt.Sprintf("❌  为文件 %s 添加 swaggerignore 失败", file.Name()))
			os.Exit(1)
		}
	}

	color.Green(fmt.Sprintf("✅  模型生成成功! 输出目录: %s", path))

	_ = os.RemoveAll(outPath)
}
