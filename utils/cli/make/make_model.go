package make

import (
	"fmt"
	"gin/common/base"
	"gin/config"
	"gin/utils"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
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
		{"-t, --table", "表名, 如: user 或 user,menu (必填)"},
		{"-p, --path", "输出目录, 如: app/model"},
		{"-c, --camel", "是否驼峰字段, 如: true"},
	}
}

func (m *MakeModel) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	table := fs.StringP("table", "t", "", "表名, 如: user 或 user,menu (必填)")
	path := fs.StringP("path", "p", "", "输出目录, 如: app/model")
	camel := fs.BoolP("camel", "c", true, "是否驼峰字段, 如: true")

	if err := fs.Parse(args); err != nil {
		color.Red("❌ 参数解析失败: %v", err)
		return
	}

	if *table == "" {
		m.ExitError(`请使用 --table 指定表名
Example: 
  go run cli.go make:model --table=user
  
Helper: go run cli.go make:model --help
`)
		return
	}

	fullPath := filepath.Join(utils.GetRootPath(), "app", "model", *path)
	fmt.Printf("✅ 创建模型: %s (表名: %s 是否使用驼峰: %v)\n", fullPath+"/"+*table+".gen.go", *table, *camel)

	tables := strings.Split(*table, ",")
	for i := range tables {
		tables[i] = strings.TrimSpace(tables[i])
	}

	m.generateFiles(fullPath, tables, *camel)
}

func init() {
	cli.Register(&MakeModel{})
}

// generateFiles 生成模型文件
func (m *MakeModel) generateFiles(path string, tables []string, camel bool) {
	root := utils.GetRootPath()
	pkg := filepath.Base(path)
	outPath := filepath.Join(root + "/app/temp")

	config.Init()
	// 读取配置
	var b strings.Builder
	// 预分配容量
	b.Grow(128)

	b.WriteString(config.Conf.Mysql.Username)
	b.WriteString(":")
	b.WriteString(config.Conf.Mysql.Password)
	b.WriteString("@tcp(")
	b.WriteString(config.Conf.Mysql.Host)
	b.WriteString(":")
	b.WriteString(config.Conf.Mysql.Port)
	b.WriteString(")/")
	b.WriteString(config.Conf.Mysql.Database)
	b.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	dsn := b.String()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		color.Red("❌ 数据库连接失败: %v", err)
		log.Fatal()
	}

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

	g.UseDB(db)

	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"json":      func(detailType gorm.ColumnType) (dataType string) { return "JsonString" },
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
				return "*model.JsonTime"
			} else {
				return "*JsonTime"
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

	color.Cyan("🚀 开始生成模型，共 %d 张表", len(tables))

	for _, table := range tables {
		color.Yellow("→ 正在生成表: %s", table)

		model := g.GenerateModel(table)
		g.ApplyBasic(model)
	}

	g.Execute()
	color.Green("✅ 模型生成成功! 输出目录: %s", path)

	_ = os.RemoveAll(outPath)
}
