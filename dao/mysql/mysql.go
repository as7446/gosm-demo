package mysql

import (
	"fmt"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gosm/model"
	"log"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "fuxuhao:fuxuhao123@tcp(10.10.12.68:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connection err: %s", err.Error())
	}
	log.Println("mysql init success...")
	DB = db
	DB.AutoMigrate(model.OssUserInfo{})
	DB.AutoMigrate(model.OssLabel{}, model.OssStorageInfo{}, model.OssUserInfo{})
	DB.AutoMigrate(model.AzureLabel{}, model.AzureStorageInfo{})
	fmt.Println("初始化完成.")
}
