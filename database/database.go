package database

import (
	"main/pkg/util"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {

	conf := DBConfig{
		DBUser: util.GetEnv("DB_USER", "root"),
		DBHost: util.GetEnv("DB_HOST", "127.0.0.1"),
		DBPort: util.GetEnv("DB_PORT", "3306"),
		DBName: util.GetEnv("DB_NAME", "alterramvc"),
	}

	mysql := mysqlConfig{
		DBConfig: conf,
	}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
