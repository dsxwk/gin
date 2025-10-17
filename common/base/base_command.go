package base

import (
	"gin/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type Command interface {
	Name() string          // 命令名称，如 "make:controller"
	Description() string   // 命令描述
	Execute(args []string) // 执行逻辑
	Help() []CommandOption // 获取命令帮助信息
}

type BaseCommand struct{}

// CommandOption 用于命令选项描述
type CommandOption struct {
	Flag string // -f, --file
	Desc string // 描述
}

// Help 默认返回nil
func (b *BaseCommand) Help() []CommandOption {
	return nil
}

func (b *BaseCommand) ExitError(msg string) {
	color.Red(`❌  ` + msg)
	os.Exit(1)
}

// GetMakePath 获取make文件路径
func (b *BaseCommand) GetMakePath(file string, _make string) string {
	if !strings.HasSuffix(file, "/") {
		file = "/" + file
	}

	switch _make {
	case "router":
		file = filepath.Join("/", "router", file)
	default:
		file = filepath.Join("/app", _make, file)
	}

	return file
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
		b.ExitError("未找到【\" + _make + \"】模版文件")
	}

	return templateFile
}
