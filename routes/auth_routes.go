package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/handlers"
)

func InitializeAuthRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
}
