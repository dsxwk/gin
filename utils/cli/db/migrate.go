package db

import (
	"gin/common/base"
	"gin/common/global"
	"gin/database"
	"gin/utils/cli"
	"github.com/fatih/color"
)

type Migrate struct{}

func (s *Migrate) Name() string {
	return "db:migrate"
}

func (s *Migrate) Description() string {
	return "数据库迁移(自动建表/更新结构)"
}

func (s *Migrate) Help() []base.CommandOption {
	return []base.CommandOption{}
}

func (s *Migrate) Execute(args []string) {
	color.Cyan("🚀 开始执行数据库迁移...")

	manager := database.NewMigrationManager(global.DB)
	err := manager.Migrate("database/migrations")
	if err != nil {
		color.Red("❌ Migration error: %v", err)
		return
	}

	color.Green("✅  数据库迁移成功!")
}

func init() {
	cli.Register(&Migrate{})
}
