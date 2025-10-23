package cli

import (
	"encoding/json"
	"fmt"
	"gin/common/base"
	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
	"os"
	"sort"
	"strings"
)

var (
	commands = make(map[string]base.Command)
)

func Register(cmd base.Command) {
	name := cmd.Name()
	if _, exists := commands[name]; exists {
		color.Yellow("⚠️  Command \"%s\" already registered, skipped.", name)
		os.Exit(1)
	}
	commands[name] = cmd
}

func Get(name string) (base.Command, bool) {
	cmd, exists := commands[name]
	return cmd, exists
}

func Execute() {
	args := os.Args[1:]
	if len(args) == 0 {
		printUsage("txt")
		return
	}

	// 全局选项
	switch args[0] {
	case "-h", "--help":
		printUsage("txt")
		return
	case "-v", "--version":
		color.Green("Gin CLI v1.0.0")
		return
	}

	if strings.HasPrefix(args[0], "-f") || strings.HasPrefix(args[0], "--format") {
		format := "txt"

		// 支持三种写法：
		//   -f json
		//   -f=json
		//   --format=json
		if len(args) > 1 && !strings.Contains(args[0], "=") {
			format = args[1]
		} else if strings.Contains(args[0], "=") {
			parts := strings.SplitN(args[0], "=", 2)
			if len(parts) == 2 {
				format = parts[1]
			}
		}
		printUsage(format)
		return
	}

	// 子命令名
	name := args[0]
	cmdArgs := args[1:]

	cmd, exists := Get(name)
	if !exists {
		color.Red("❌  Command \"%s\" is not defined.", name)
		printUsage("txt")
		os.Exit(1)
	}

	// --help 自动打印命令帮助
	for _, arg := range cmdArgs {
		if arg == "-h" || arg == "--help" {
			printCommandHelp(cmd)
			return
		}
	}

	// 交给命令自己解析参数
	cmd.Execute(cmdArgs)
}

// 打印命令列表
func printUsage(format string) {
	switch format {
	case "json":
		printJSON()
	default:
		printText()
	}
}

func printText() {
	color.Cyan("Usage: cli [command] [options]\n")
	color.Yellow("Available commands:")

	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		cmd := commands[name]
		fmt.Printf("  %s%s\n", color.GreenString(fmt.Sprintf("%-25s", name)), cmd.Description())
	}

	fmt.Println()
	color.Yellow("Options:")
	fmt.Println("  -f, --format        The output format (txt, json) [default: txt]")
	fmt.Println("  -h, --help          Display help for the given command")
	fmt.Println("  -v, --version       Display CLI version")
}

// 打印单个命令帮助
func printCommandHelp(cmd base.Command) {
	fmt.Printf("\n%s - %s\n\n", color.GreenString(cmd.Name()), cmd.Description())

	options := cmd.Help()
	if len(options) == 0 {
		fmt.Println("该命令暂无选项")
		return
	}

	color.Yellow("Options:")

	// 计算最大显示宽度（基于未上色的原始字符串）
	maxFlagWidth := 0
	maxDescWidth := 0
	for _, opt := range options {
		flagStr := fmt.Sprintf("-%s, --%s", opt.Flag.Short, opt.Flag.Long)
		if w := runewidth.StringWidth(flagStr); w > maxFlagWidth {
			maxFlagWidth = w
		}
		if w := runewidth.StringWidth(opt.Desc); w > maxDescWidth {
			maxDescWidth = w
		}
	}

	// 打印，每列手动追加空格(基于显示宽度)
	for _, opt := range options {
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
			spaces(flagPad),
			colDesc,
			spaces(descPad),
			required,
		)
	}
}

func spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

func printJSON() {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	var list []map[string]string
	for _, name := range names {
		cmd := commands[name]
		list = append(list, map[string]string{
			"name":        name,
			"description": cmd.Description(),
		})
	}

	data := map[string]interface{}{
		"version":  "Gin CLI v2.0.0",
		"commands": list,
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	color.Green(string(jsonData))
}
