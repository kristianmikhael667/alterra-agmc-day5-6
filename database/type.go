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

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBUser,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
