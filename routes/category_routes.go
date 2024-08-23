package routes

import (
	"github.com/gin-gonic/gin"
	category_handler "github.com/isaiaspereira307/gopvp/handlers/category_handler"
)

func InitializeCategoryRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/category/:id", category_handler.GetCategory)
		v1.GET("/categories", category_handler.GetCategories)
		v1.POST("/category", category_handler.CreateCategory)
		v1.PUT("/category/:id", category_handler.UpdateCategory)
		v1.DELETE("/category/:id", category_handler.DeleteCategory)
	}
}
