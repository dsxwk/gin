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

type MakeRequest struct {
	base.BaseCommand
}

func (m *MakeRequest) Name() string {
	return "make:request"
}

func (m *MakeRequest) Description() string {
	return "验证请求创建"
}

func (m *MakeRequest) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			"-f, --file",
			"文件路径, 如: user",
			true,
		},
		{
			"-d, --desc",
			"描述, 如: 用户请求验证",
			false,
		},
	}
}

func (m *MakeRequest) Execute(args []string) {
	fs := pflag.NewFlagSet(m.Name(), pflag.ExitOnError)
	_make := strings.TrimPrefix(m.Name(), "make:")
	file := fs.StringP("file", "f", "", "文件路径, 如: user")
	desc := fs.StringP("desc", "d", "Validator", "描述")
	fmt.Printf("✅ 创建服务: %s (描述: %s)\n", *file, *desc)
	f := m.GetMakePath(*file, _make)

	m.generateFile(_make, f, *desc)
}

func init() {
	cli.Register(&MakeRequest{})
}

func (m *MakeRequest) generateFile(_make, file, desc string) {
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
		Description string // 如果为空,使用默认值
	}{
		Package:     packageName,
		Name:        utils.UcFirst(strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file)))),
		Description: desc,
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		return
	} else {
		fmt.Println("Template executed and content written to file.")
	}

	color.Green("✅ 验证请求文件: " + file + " 生成成功!")
	fmt.Println()
}
