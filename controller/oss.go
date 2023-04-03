package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosm/dao/mysql"
	"gosm/model"
	"gosm/service"
	"log"
	"net/http"
)

var Oss oss

type oss struct{}

func (o *oss) CreateOssHandler(ctx *gin.Context) {
	parmas := new(struct {
		Username      string            `form:"username"`
		BucketName    string            `form:"bucket_name"`
		Region        string            `form:"region"`
		ProjectNumber string            `form:"project_number"`
		ProjectUsage  string            `form:"project_usage"`
		Labels        map[string]string `form:"labels"`
	})
	if err := ctx.Bind(parmas); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":  "解析请求参数错误",
			"data": nil,
		})
		return
	}
	user := model.OssUserInfo{Username: parmas.Username}
	tx := mysql.DB.Where(&user).First(&user)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fmt.Sprintf("error: %s ，未知的user", tx.Error),
			"data": nil,
		})
		return
	}
	err := service.Oss.CreateStorage(user, parmas.BucketName, parmas.Region, parmas.ProjectNumber, parmas.ProjectUsage, parmas.Labels)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "创建Oss Storage失败",
			"data": nil,
		})
		return
	}
	log.Println(tx.RowsAffected, tx.Row(), tx.Error, user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "请求成功",
		"data": parmas,
	})
}
