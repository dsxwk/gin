package config

import (
	"gin/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openPgsql() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(getPgsqlDsn()), &gorm.Config{})
}

func getPgsqlDsn() string {
	return utils.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Conf.Pgsql.Host,
		Conf.Pgsql.Username,
		Conf.Pgsql.Password,
		Conf.Pgsql.Database,
		Conf.Pgsql.Port,
	)
}
