package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/handlers"
)

func InitializeAuthRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", handlers.Login)
	}
}
