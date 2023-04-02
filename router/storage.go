package router

import "github.com/gin-gonic/gin"

func StorageRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("/storage/:type")
}
