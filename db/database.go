package db

import (
	"CanRich/config"
	"CanRich/logger"
	"CanRich/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func InitDB() {
	username := config.GlobalConfig.GetString("datasource.username")
	password := config.GlobalConfig.GetString("datasource.password")
	host := config.GlobalConfig.GetString("datasource.host")
	port := config.GlobalConfig.GetString("datasource.port")
	dbName := config.GlobalConfig.GetString("datasource.database")
	driverName := config.GlobalConfig.GetString("datasource.driverName")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName)

	var err error
	DB, err = gorm.Open(driverName, dbUrl)
	if err != nil {
		logger.SugarLogger.Error("Failed to connect database. Error: ", err.Error())
	}

	DB.SingularTable(true)
	//数据库迁移
	DB.AutoMigrate(&model.User{})
	DB.DB().SetMaxOpenConns(50)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetConnMaxLifetime(10 * time.Second)
	//DB.LogMode(true)
}
