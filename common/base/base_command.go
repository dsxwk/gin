package base

import (
	"bufio"
	"fmt"
	"gin/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type BaseCommand struct{}

// CommandOption 用于命令选项描述
type CommandOption struct {
	Flag     string // -f, --file
	Desc     string // 描述
	Required bool   // 是否必填
}

type Command interface {
	Name() string          // 命令名称，如 "make:controller"
	Description() string   // 命令描述
	Execute(args []string) // 执行逻辑
	Help() []CommandOption // 获取命令帮助信息
}

// Help 默认返回nil
func (b *BaseCommand) Help() []CommandOption {
	return nil
}

func (b *BaseCommand) ExitError(msg string) {
	color.Red("❌  %s", msg)
	os.Exit(1)
}

// GetMakeFile 获取make文件
func (b *BaseCommand) GetMakeFile(file string, _make string) string {
	// 去除前斜杠
	file = strings.TrimPrefix(file, "/")

	switch _make {
	case "router":
		file = filepath.Join("router", file)
	default:
		file = filepath.Join("app", _make, file)
	}

	return file + ".go"
}

// GetTemplate 获取模版文件
func (b *BaseCommand) GetTemplate(_make string) string {
	var (
		templateFile string
	)

	switch _make {
	case "model":
	case "command", "controller", "service", "request", "middleware", "router":
		templateFile = filepath.Join(utils.GetRootPath(), "common", "template", _make+".tpl")
	default:
		b.ExitError("未找到 \" + _make + \" 模版文件")
	}

	return templateFile
}

// CheckDirAndFile 检查目录和文件
func (b *BaseCommand) CheckDirAndFile(file string) *os.File {
	// 如果目录不存在则创建
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0755); err != nil {
		color.Red("❌ Failed to create directory:", err)
		return nil
	}

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		fmt.Printf("%s 文件 %s 已存在,是否覆盖?(%s/%s): ",
			color.YellowString("⚠️"),
			color.CyanString(file),
			color.GreenString("Y"),
			color.RedString("N"),
		)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
			log.Println("操作已取消")
			return nil
		}
	}

	color.Green("📄 创建文件: %s\n", color.CyanString(file))
	f, err := os.Create(file)
	if err != nil {
		color.Red("❌ Failed to create file:", err.Error())
		return nil
	}
	return f
}
