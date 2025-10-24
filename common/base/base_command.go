package base

import (
	"bufio"
	"fmt"
	"gin/utils"
	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
	"strings"
)

type BaseCommand struct{}

// Flag 定义短长参数名
type Flag struct {
	Short   string
	Long    string
	Default string
}

// CommandOption 用于命令选项描述
type CommandOption struct {
	Flag     Flag   // flag定义
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

// ParseFlags flag解析
func (b *BaseCommand) ParseFlags(name string, args []string, opts []CommandOption) map[string]string {
	// fs := pflag.NewFlagSet(name, pflag.ExitOnError)
	// ContinueOnError防止自动退出
	fs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	fs.SetOutput(nil) // 禁止默认输出Usage信息

	// 暂存flag引用
	flagRefs := make(map[string]*string)

	for _, opt := range opts {
		defVal := opt.Flag.Default
		flagRefs[opt.Flag.Long] = fs.StringP(opt.Flag.Long, opt.Flag.Short, defVal, opt.Desc)
	}

	// 解析命令参数
	err := fs.Parse(args)
	if err != nil {
		color.Red("❌  argument error, %s is not defined.", err.Error())
		color.Cyan("Usage: cli %s [args]", name)
		fmt.Println()
		color.Yellow("Available args:")
		PrintArgs(opts)
		os.Exit(1)
	}

	// 构建结果 map
	values := make(map[string]string)
	for key, ref := range flagRefs {
		values[key] = *ref
	}

	// 检查必填参数
	for _, opt := range opts {
		val := values[opt.Flag.Long]
		if opt.Required && val == "" {
			b.ExitError(fmt.Sprintf("参数 -%s 或 --%s 不能为空", opt.Flag.Short, opt.Flag.Long))
		}
	}

	return values
}

// PrintArgs 打印参数
func PrintArgs(opts []CommandOption) {
	// 计算最大显示宽度(基于未上色的原始字符串)
	maxFlagWidth := 0
	maxDescWidth := 0
	for _, opt := range opts {
		flagStr := fmt.Sprintf("-%s, --%s", opt.Flag.Short, opt.Flag.Long)
		if w := runewidth.StringWidth(flagStr); w > maxFlagWidth {
			maxFlagWidth = w
		}
		if w := runewidth.StringWidth(opt.Desc); w > maxDescWidth {
			maxDescWidth = w
		}
	}

	// 打印，每列手动追加空格(基于显示宽度)
	for _, opt := range opts {
		flagStr := fmt.Sprintf("-%s, --%s", opt.Flag.Short, opt.Flag.Long)
		descStr := opt.Desc

		// 颜色化显示内容(不要用于计算宽度)
		colFlag := color.GreenString(flagStr)
		colDesc := descStr // 不上色描述也行,若想上色可以color.YellowString(descStr)

		// 计算需要的空格数(基于显示宽度)
		flagPad := maxFlagWidth - runewidth.StringWidth(flagStr) + 2 // +2 列间距
		descPad := maxDescWidth - runewidth.StringWidth(descStr) + 2

		required := color.GreenString("required:false")
		if opt.Required {
			required = color.RedString("required:true")
		}

		// 输出：带颜色的 flag + 空格 + 描述 + 空格 + required
		fmt.Printf("  %s%s%s%s%s\n",
			colFlag,
			utils.Spaces(flagPad),
			colDesc,
			utils.Spaces(descPad),
			required,
		)
	}
}

// FormatArgs 格式化参数
func (b *BaseCommand) FormatArgs(args map[string]string) string {
	str := ""
	for arg, value := range args {
		str += fmt.Sprintf("--%s=%s", arg, value)
	}

	return str
}

// StringToBool 将字符串安全地转换为布尔值
func (b *BaseCommand) StringToBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "1", "true", "yes", "y", "on":
		return true
	case "0", "false", "no", "n", "off":
		return false
	default:
		return false // 默认返回false,防止解析异常
	}
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
		b.ExitError("未找到 " + _make + " 模版文件")
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

		if !b.StringToBool(input) {
			fmt.Println("操作已取消")
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
