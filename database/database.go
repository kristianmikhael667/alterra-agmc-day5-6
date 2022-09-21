package database

import (
	"main/pkg/util"
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	dbusername := os.Getenv("DB_USERNAME")
	dbpass := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	conf := DBConfig{
		DBUser: util.GetEnv("DB_USERNAME", dbusername),
		DBPass: util.GetEnv("DB_PASSWORD", dbpass),
		DBHost: util.GetEnv("DB_PASSWORD", dbport),
		DBPort: util.GetEnv("DB_PASSWORD", dbhost),
		DBName: util.GetEnv("DB_PASSWORD", dbname),
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
