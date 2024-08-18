package db

import (
	"fmt"

	"github.com/fuatanshori/gin-gorm/configs/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var errConnection error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
	DB, errConnection := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errConnection != nil {
		panic(errConnection)
	}
	return DB
}
