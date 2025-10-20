package cli

import (
	"encoding/json"
	"fmt"
	"gin/common/base"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"log"
	"os"
	"sort"
	"strings"
)

// 全局命令注册
var commands = map[string]base.Command{}

// Register 注册命令
func Register(cmd base.Command) {
	name := cmd.Name()
	if _, exists := commands[name]; exists {
		color.Yellow("⚠️  Command \"%s\" already registered, skipped.", name)
		os.Exit(1)
	}

	commands[name] = cmd
}

// Get 获取命令
func Get(name string) (base.Command, bool) {
	cmd, exists := commands[name]
	return cmd, exists
}

// Execute 执行命令
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
		fmt.Println()
		printUsage("txt")
		os.Exit(1)
	}

	// 如果子命令参数中包含 -h 或 --help，直接打印子命令帮助
	for _, arg := range cmdArgs {
		if arg == "-h" || arg == "--help" {
			printCommandHelp(cmd)
			return
		}
	}

	cmdFs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	cmdFs.SetOutput(os.Stderr) // 避免默认打印
	cmdFs.Usage = func() {}    // 禁止默认 help

	// 注册子命令选项
	optionMap := map[string]base.CommandOption{} // 保存对应的 option
	for _, opt := range cmd.Help() {
		if strings.Contains(opt.Flag, ",") {
			parts := strings.Split(opt.Flag, ",")
			short := strings.TrimSpace(parts[0])
			long := strings.TrimSpace(parts[1])
			cmdFs.StringP(strings.TrimPrefix(long, "--"), strings.TrimPrefix(short, "-"), "", opt.Desc)
			optionMap[strings.TrimPrefix(long, "--")] = opt
		} else {
			cmdFs.String(strings.TrimPrefix(opt.Flag, "--"), "", opt.Desc)
			optionMap[strings.TrimPrefix(opt.Flag, "--")] = opt
		}
	}

	// 解析子命令参数
	if err := cmdFs.Parse(cmdArgs); err != nil {
		// 如果是未知选项
		if strings.Contains(err.Error(), "unknown flag") {
			color.Red("❌  Option \"%s\" is not defined.", extractFlag(err.Error()))
			log.Println()
			printCommandHelp(cmd)
			os.Exit(1)
		}
		color.Red("❌  %s", err.Error())
		printCommandHelp(cmd)
		os.Exit(2)
	}

	if err := cmdFs.Parse(cmdArgs); err != nil {
		color.Red("❌  %s", err.Error())
		printCommandHelp(cmd)
		os.Exit(2)
	}

	// 检查必填参数
	for key, opt := range optionMap {
		val := cmdFs.Lookup(key).Value.String()
		if opt.Required && val == "" {
			color.Red("❌  参数: --%s 不能为空", key)
			color.Cyan(fmt.Sprintf("Example: go run cli.go %s --%s", cmd.Name(), key))
			fmt.Println()
			color.Cyan(fmt.Sprintf("Helper: go run cli.go %s --help", cmd.Name()))
			printCommandHelp(cmd)
			os.Exit(3)
		}
	}

	// 执行子命令
	cmd.Execute(cmdArgs)
}

// extractFlag 提取未知flag
func extractFlag(msg string) string {
	parts := strings.Split(msg, " ")
	if len(parts) > 2 {
		return parts[2]
	}
	return ""
}

// 命令选项提示 printCommandHelp
func printCommandHelp(cmd base.Command) {
	fmt.Printf("\n%s - %s\n\n", color.GreenString(cmd.Name()), cmd.Description())

	options := cmd.Help()
	if len(options) == 0 {
		fmt.Println("该命令暂无选项")
		return
	}

	color.Yellow("Options:")

	// 计算flag最大长度用于对齐
	maxFlagLen := 0
	maxDescLen := 0
	for _, opt := range options {
		if len(opt.Flag) > maxFlagLen {
			maxFlagLen = len(opt.Flag)
		}
		if len(opt.Desc) > maxDescLen {
			maxDescLen = len(opt.Desc)
		}
	}

	for _, opt := range options {
		// flag与描述之间填充空格
		flagPadding := strings.Repeat(" ", maxFlagLen-len(opt.Flag)+2)
		descPadding := strings.Repeat(" ", maxDescLen-len(opt.Desc)+2)

		required := color.GreenString("required:false")
		if opt.Required {
			required = color.RedString("required:true")
		}

		fmt.Printf("  %s%s%s%s%s\n",
			color.GreenString(opt.Flag),
			flagPadding,
			opt.Desc,
			descPadding,
			required,
		)
	}
}

// printUsage 打印命令列表
func printUsage(format string) {
	switch format {
	case "json":
		printJSON()
	default:
		printText()
	}
}

// printText text格式输出
func printText() {
	color.Cyan("Usage: go run cli.go [command] [options]\n")
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

	options := []struct {
		flag string
		desc string
	}{
		{"-f, --format", "The output format (txt, json) [default: txt]"},
		{"-h, --help", "Display help for the given command. When no command is given display help for the list command"},
		{"-v, --version", "Display this application version"},
	}

	maxLen := 0
	for _, opt := range options {
		if len(opt.flag) > maxLen {
			maxLen = len(opt.flag)
		}
	}

	for _, opt := range options {
		padding := strings.Repeat(" ", maxLen-len(opt.flag)+2)
		fmt.Printf("  %s%s%s\n", color.GreenString(opt.flag), padding, opt.desc)
	}
}

// printJSON json格式输出
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
		"version":  "Gin CLI v1.0.0",
		"commands": list,
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")
	color.Green(string(jsonData))
}
