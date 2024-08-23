package routes

import (
	"github.com/gin-gonic/gin"
	category_handler "github.com/isaiaspereira307/gopvp/handlers/category_handler"
)

func InitializeCategoryRoutes(router *gin.RouterGroup) {
	router.GET("/category/:id", category_handler.GetCategory)
	router.GET("/categories", category_handler.GetCategories)
	router.POST("/category", category_handler.CreateCategory)
	router.PUT("/category/:id", category_handler.UpdateCategory)
	router.DELETE("/category/:id", category_handler.DeleteCategory)
}
