package db

import (
	"gin/common/base"
	"gin/common/global"
	"gin/database"
	"gin/utils/cli"
	"github.com/fatih/color"
)

type Seed struct{}

func (s *Seed) Name() string {
	return "db:seed"
}

func (s *Seed) Description() string {
	return "数据初始化"
}

func (s *Seed) Help() []base.CommandOption {
	return []base.CommandOption{}
}

func (s *Seed) Execute(args []string) {
	color.Cyan("🚀  开始执行数据库 Seed...")
	manager := database.NewMigrationManager(global.DB)

	// 默认 seed 文件路径
	seedFile := "database/seeds/initial_data.sql"
	if len(args) > 0 {
		seedFile = args[0]
	}

	// 执行 seed
	if err := manager.Seed(seedFile); err != nil {
		color.Red("❌ Seed 执行失败: %v", err)
		return
	}

	color.Green("✅ 数据初始化完成!")
}

func init() {
	cli.Register(&Seed{})
}
