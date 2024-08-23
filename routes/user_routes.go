package routes

import (
	"github.com/gin-gonic/gin"
	users_handlers "github.com/isaiaspereira307/gopvp/handlers/users_handler"
)

func InitializeUserRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", users_handlers.GetUsers)
		v1.GET("/users/:id", users_handlers.GetUser)
		v1.POST("/users", users_handlers.CreateUser)
		v1.PUT("/users/:id", users_handlers.UpdateUser)
		v1.DELETE("/users/:id", users_handlers.DeleteUser)
	}
}
