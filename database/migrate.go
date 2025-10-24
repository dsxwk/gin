package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Migration struct {
	ID        int       `json:"id"`
	Version   string    `json:"version"`
	Name      string    `json:"name"`
	SQL       string    `json:"sql"`
	CreatedAt time.Time `json:"created_at"`
}

type MigrationManager struct {
	db *gorm.DB
}

func NewMigrationManager(db *gorm.DB) *MigrationManager {
	return &MigrationManager{db: db}
}

// InitMigrationTable 初始化migrations表
func (m *MigrationManager) InitMigrationTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS migrations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		version VARCHAR(50) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		sql_content TEXT NOT NULL,
		executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_version (version),
		INDEX idx_executed_at (executed_at)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`
	return m.db.Exec(query).Error
}

// GetExecutedMigrations 获取已执行的迁移
func (m *MigrationManager) GetExecutedMigrations() (map[string]Migration, error) {
	var migrations []Migration
	err := m.db.Raw(`SELECT id, version, name, sql_content, executed_at FROM migrations ORDER BY version`).Scan(&migrations).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]Migration)
	for _, model := range migrations {
		result[model.Version] = model
	}
	return result, nil
}

// LoadMigrationFiles 加载迁移文件
func (m *MigrationManager) LoadMigrationFiles(dir string) ([]Migration, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		return nil, err
	}

	var (
		migrations []Migration
		content    []byte
	)
	for _, file := range files {
		filename := filepath.Base(file)
		parts := strings.Split(strings.TrimSuffix(filename, ".sql"), "_")
		if len(parts) < 2 {
			log.Printf("⚠️  跳过无效文件名: %s", filename)
			continue
		}
		version := parts[0]
		name := strings.Join(parts[1:], "_")

		content, err = os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		migrations = append(migrations, Migration{
			Version: version,
			Name:    name,
			SQL:     string(content),
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})
	return migrations, nil
}

// ExecuteMigration 执行单个迁移
func (m *MigrationManager) ExecuteMigration(mg Migration) error {
	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Exec(mg.SQL).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("执行迁移失败 %s: %w", mg.Version, err)
	}

	if err := tx.Exec(`INSERT INTO migrations (version, name, sql_content) VALUES (?, ?, ?)`,
		mg.Version, mg.Name, mg.SQL).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录迁移失败 %s: %w", mg.Version, err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("提交迁移失败 %s: %w", mg.Version, err)
	}

	log.Printf("✅  执行迁移成功: %s - %s", mg.Version, mg.Name)
	return nil
}

// Migrate 执行全部迁移
func (m *MigrationManager) Migrate(dir string) error {
	if err := m.InitMigrationTable(); err != nil {
		return err
	}

	executed, err := m.GetExecutedMigrations()
	if err != nil {
		return err
	}

	migrations, err := m.LoadMigrationFiles(dir)
	if err != nil {
		return err
	}

	count := 0
	for _, mg := range migrations {
		if _, exists := executed[mg.Version]; !exists {
			if err = m.ExecuteMigration(mg); err != nil {
				return err
			}
			count++
		}
	}

	if count == 0 {
		log.Println("✅  没有待执行的迁移文件")
	} else {
		log.Printf("✅  已执行 %d 个迁移", count)
	}

	return nil
}

// Rollback 回滚最后一次迁移
func (m *MigrationManager) Rollback() error {
	var migration Migration
	err := m.db.Raw(`SELECT id, version, name FROM migrations ORDER BY id DESC LIMIT 1`).Scan(&migration).Error
	if err != nil {
		return fmt.Errorf("查询最后迁移失败: %w", err)
	}
	if migration.ID == 0 {
		return fmt.Errorf("没有迁移可回滚")
	}

	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err = tx.Exec(`DELETE FROM migrations WHERE id = ?`, migration.ID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("删除迁移记录失败: %w", err)
	}

	if err = tx.Commit().Error; err != nil {
		return fmt.Errorf("提交回滚失败: %w", err)
	}

	log.Printf("✅  已回滚迁移: %s - %s", migration.Version, migration.Name)
	return nil
}

// Status 显示迁移状态
func (m *MigrationManager) Status(dir string) error {
	executed, err := m.GetExecutedMigrations()
	if err != nil {
		return err
	}

	migrations, err := m.LoadMigrationFiles(dir)
	if err != nil {
		return err
	}

	fmt.Println("Migration Status:")
	fmt.Println("==================")

	var pending int
	for _, mg := range migrations {
		if _, ok := executed[mg.Version]; ok {
			fmt.Printf("✅  %s - %s (executed)\n", mg.Version, mg.Name)
		} else {
			fmt.Printf("⏳  %s - %s (pending)\n", mg.Version, mg.Name)
			pending++
		}
	}
	fmt.Printf("\nTotal: %d, Pending: %d\n", len(migrations), pending)
	return nil
}

// Seed 执行seed
func (m *MigrationManager) Seed(file string) error {
	if file == "" {
		file = "database/seeds/initial_data.sql"
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return fmt.Errorf("seed 文件不存在: %s", file)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if err = m.db.Exec(string(content)).Error; err != nil {
		return fmt.Errorf("执行 seed 失败: %w", err)
	}
	fmt.Printf("✅  Seed 执行完成: %s", file)
	return nil
}

// Reset 重置数据库
func (m *MigrationManager) Reset(dir string) error {
	log.Println("⚠️  此操作将删除所有表并重新建表，是否继续？(y/N)")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return err
	}
	if strings.ToLower(input) != "y" && strings.ToLower(input) != "yes" {
		log.Println("操作已取消")
		return nil
	}

	query := `
	SET FOREIGN_KEY_CHECKS = 0;
	DROP TABLE IF EXISTS migrations, user1;
	SET FOREIGN_KEY_CHECKS = 1;
	`
	if err = m.db.Exec(query).Error; err != nil {
		return err
	}

	log.Println("✅  所有表已清空")
	return m.Migrate(dir)
}
