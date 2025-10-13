package make

import (
	"fmt"
	"gin/common/base"
	"gin/utils"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type MakeRouter struct {
	base.BaseCommand
}

func (m *MakeRouter) Name() string {
	return "make:router"
}

func (m *MakeRouter) Description() string {
	return "路由创建"
}

func (m *MakeRouter) Help() []base.CommandOption {
	return []base.CommandOption{
		{"-f, --file", "文件路径, 如: user (必填)"},
	}
}

func (m *MakeRouter) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	_make := strings.TrimPrefix(m.Name(), "make:")
	file := fs.StringP("file", "f", "", "文件路径, 如: user (必填)")

	if err := fs.Parse(args); err != nil {
		fmt.Println("解析参数失败:", err.Error())
		return
	}

	if *file == "" {
		m.ExitError("请使用 --file 指定文件路径\nExample: go run cli.go make:router --file=user \nHelper: go run cli.go make:router --help")
		return
	}

	fmt.Printf("✅ 创建路由: %s \n", *file)
	f := m.GetMakePath(*file, _make)

	m.generateFile(_make, f)
}

func (m *MakeRouter) generateFile(_make, file string) {
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
		Package string // 提取的包名
		Name    string // 模块名称(首字母大写)
	}{
		Package: packageName,
		Name:    utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		return
	} else {
		fmt.Println("Template executed and content written to file.")
	}

	color.Green("✅ 路由文件: " + file + " 生成成功!")
	fmt.Println()
}
