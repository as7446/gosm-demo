package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRouter(r *gin.Engine) {
	r.GET("/api/v1/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "状态正常",
		})
	})
}
