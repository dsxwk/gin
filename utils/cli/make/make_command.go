package make

import (
	"fmt"
	"gin/common/base"
	"gin/utils"
	"gin/utils/cli"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type MakeCommand struct {
	base.BaseCommand
}

func (m *MakeCommand) Name() string {
	return "make:command"
}

func (m *MakeCommand) Description() string {
	return "服务创建"
}

func (m *MakeCommand) Help() []base.CommandOption {
	return []base.CommandOption{
		{"-f, --file", "文件路径, 如: cronjob/demo", true},
		{"-m, --name", "命令名称, 如: demo.command", false},
		{"-d, --desc", "描述, 如: demo-desc", false},
	}
}

func (m *MakeCommand) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	_make := strings.TrimPrefix(m.Name(), "make:")
	file := fs.StringP("file", "f", "", "文件路径, 如: cronjob/demo")
	command := fs.StringP("name", "m", "", "demo.command")
	desc := fs.StringP("desc", "d", "", "command-desc")

	if err := fs.Parse(args); err != nil {
		fmt.Println("解析参数失败:", err.Error())
		return
	}

	if *file == "" {
		m.ExitError(`请使用 --file 指定文件路径
Example: go run cli.go make:command --file=demo --desc=command-desc

Helper: go run cli.go make:command --help
`)
		return
	}

	fmt.Printf("✅ 创建命令: %s (命令: %s 描述: %s)\n", *file, *command, *desc)
	f := m.GetMakePath(*file, _make)

	m.generateFile(_make, f, *command, *desc)
}

func init() {
	cli.Register(&MakeCommand{})
}

func (m *MakeCommand) generateFile(_make, file, command, desc string) {
	var (
		rootPath = utils.GetRootPath()
	)

	templateFile := m.GetTemplate(_make)
	color.Green(fmt.Sprintf("Loading template file: %s\n", templateFile))
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error parsing template:", err.Error())
		return
	}
	fmt.Printf("解析模板文件: %v\n", tmpl.Name())

	// 提取包名 (文件路径中的最后一个目录作为包名)
	packageName := filepath.Base(filepath.Dir(file))
	fmt.Printf("Detected package name: %s\n", packageName)

	// 确保目录存在
	dir := filepath.Dir(rootPath + file)
	fmt.Printf("Checking if directory exists: %s\n", dir)

	// 使用 os.Stat 检查目录是否存在
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist. Creating: %s\n", dir)
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Failed to create directory:", err.Error())
			return
		}
	} else {
		fmt.Println("Directory already exists.")
	}

	// 创建文件
	file = filepath.Join(rootPath, file+".go")
	fmt.Printf("Creating file: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Failed to create file:", err.Error())
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			fmt.Println("Failed to close file:", err.Error())
		}
	}(f)

	data := struct {
		Package     string // 提取的包名
		Name        string // 模块名称(首字母大写)
		Command     string // 如果为空,使用默认值
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Command:     command,
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		return
	} else {
		fmt.Println("Template executed and content written to file.")
	}

	color.Green("✅ 命令行文件: " + file + " 生成成功!")
	fmt.Println()
}
