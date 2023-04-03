package router

import (
	"github.com/gin-gonic/gin"
	"gosm/controller"
)

func StorageRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/storage")
	v1.GET("/osss")
	v1.POST("/oss", controller.Oss.CreateOssHandler)
	v1.GET("/oss")
	v1.PUT("/oss")
	v1.DELETE("/oss")
}
