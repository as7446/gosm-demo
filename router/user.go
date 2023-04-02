package router

import (
	"github.com/gin-gonic/gin"
	"gosm/controller"
)

func UserRouter(r *gin.Engine) {
	r.POST("/register", controller.UserRegister)
	r.POST("/login", controller.UserLogin)
}
