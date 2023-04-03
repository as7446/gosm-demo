package main

import (
	"gosm/dao/mysql"
	"gosm/router"
	"gosm/service"
)

func main() {
	mysql.InitMysql()
	service.InitOssClient()
	service.InitRamClient()
	r := router.InitRouter()
	r.Run(":8000")

}
