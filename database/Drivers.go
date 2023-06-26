package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Database string
	Password string
	Username string
	Port     string
}

func PostgresqlDriver(config *Config) (*gorm.DB, error) {
	dsn := "host=" + config.Host + " user=" + config.Username + " password=" + config.Password +
		" dbname=" + config.Database + " port=" + config.Port + " sslmode=disable TimeZone=Asia/Tehran"

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func MysqlDriver(config *Config) (*gorm.DB, error) {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port +
		")/gorm?charset=utf8&parseTime=True&loc=Local"

	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
}
