package db

import (
	"gin/common/base"
	"gin/common/global"
	"gin/database"
	"gin/utils/cli"
	"github.com/fatih/color"
)

type Seed struct {
	base.BaseCommand
}

func (s *Seed) Name() string {
	return "db:seed"
}

func (s *Seed) Description() string {
	return "数据初始化"
}

func (s *Seed) Help() []base.CommandOption {
	return []base.CommandOption{
		{
			base.Flag{
				Short:   "f",
				Long:    "file",
				Default: "database/seeds/init_user1_data.sql",
			},
			"seed 文件, 如: database/seeds/init_user1_data.sql",
			true,
		},
	}
}

func (s *Seed) Execute(args []string) {
	values := s.ParseFlags(s.Name(), args, s.Help())
	color.Cyan("开始执行数据库 Seed...")
	manager := database.NewMigrationManager(global.DB)

	// 执行seed
	if err := manager.Seed(values["file"]); err != nil {
		color.Red("❌ %v", err)
		return
	}
}

func init() {
	cli.Register(&Seed{})
}
