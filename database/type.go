package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	DBConfig struct {
		DBHost string
		DBUser string
		DBPass string
		DBPort string
		DBName string
	}

	mysqlConfig struct {
		DBConfig
	}
)

func (conf mysqlConfig) Connect() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBPass,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
