package cli

import (
	"encoding/json"
	"fmt"
	"gin/common/base"
	"github.com/fatih/color"
	"github.com/spf13/pflag"
	"os"
	"sort"
	"strings"
)

// 全局命令注册
var commands = map[string]base.Command{}

// Register 注册命令
func Register(cmd base.Command) {
	commands[cmd.Name()] = cmd
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
	for _, a := range cmdArgs {
		if a == "-h" || a == "--help" {
			printCommandHelp(cmd)
			return
		}
	}

	cmdFs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	cmdFs.SetOutput(os.Stderr) // 避免默认打印
	cmdFs.Usage = func() {}    // 禁止默认 help

	// 注册子命令选项
	for _, opt := range cmd.Help() {
		if strings.Contains(opt.Flag, ",") {
			parts := strings.Split(opt.Flag, ",")
			short := strings.TrimSpace(parts[0])
			long := strings.TrimSpace(parts[1])
			cmdFs.StringP(strings.TrimPrefix(long, "--"), strings.TrimPrefix(short, "-"), "", opt.Desc)
		} else {
			cmdFs.String(strings.TrimPrefix(opt.Flag, "--"), "", opt.Desc)
		}
	}

	// 解析子命令参数
	if err := cmdFs.Parse(cmdArgs); err != nil {
		// 如果是未知选项
		if strings.Contains(err.Error(), "unknown flag") {
			color.Red("❌  Option \"%s\" is not defined.", extractFlag(err.Error()))
			fmt.Println()
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

	// 执行子命令
	cmd.Execute(cmdArgs)
}

// AutoRegister 自动注册
func AutoRegister(cmd base.Command) {
	name := cmd.Name()
	if _, exists := commands[name]; exists {
		color.Yellow("⚠️  Command \"%s\" already registered, skipped.", name)
		return
	}
	commands[name] = cmd
	// color.Green("✅  Registered command: %s", name)
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
	// 对齐计算
	maxLen := 0
	for _, opt := range options {
		if len(opt.Flag) > maxLen {
			maxLen = len(opt.Flag)
		}
	}

	for _, opt := range options {
		padding := strings.Repeat(" ", maxLen-len(opt.Flag)+2)
		fmt.Printf("  %s%s%s\n", color.GreenString(opt.Flag), padding, opt.Desc)
	}
	fmt.Println()
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
		{"--format=txt", "The output format (txt, json) [default: \"txt\"]"},
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
	fmt.Println()
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
	fmt.Println(string(jsonData))
}
