package routes

import (
	"github.com/gin-gonic/gin"
	users_handlers "github.com/isaiaspereira307/gopvp/handlers/users_handler"
)

func InitializeUserRoutes(router *gin.RouterGroup) {
	router.GET("/users", users_handlers.GetUsers)
	router.GET("/users/:id", users_handlers.GetUser)
	router.POST("/users", users_handlers.CreateUser)
	router.PUT("/users/:id", users_handlers.UpdateUser)
	router.DELETE("/users/:id", users_handlers.DeleteUser)
}
