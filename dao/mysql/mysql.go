package mysql

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/cost_account?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connection err: %s", err.Error())
	}
	log.Println("mysql init success...")
	DB = db
}
