package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRouter(r *gin.Engine) {
	r.Group("/api/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "状态正常",
		})
	})
}
