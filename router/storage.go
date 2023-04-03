package router

import (
	"github.com/gin-gonic/gin"
	"gosm/controller"
)

func StorageRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/storage")
	v1.GET("/buckets")
	v1.POST("/bucket", controller.Oss.CreateOssHandler)
	v1.GET("/bucket")
	v1.PUT("/bucket")
	v1.DELETE("/bucket")
}
