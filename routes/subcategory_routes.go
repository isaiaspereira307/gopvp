package routes

import (
	"github.com/gin-gonic/gin"
	subcategory_handler "github.com/isaiaspereira307/gopvp/handlers/subcategory_handler"
)

func InitializeSubcategoryRoutes(router *gin.RouterGroup) {
	router.GET("/subcategory/:id", subcategory_handler.GetSubcategory)
	router.GET("/subcategories", subcategory_handler.GetSubcategories)
	router.POST("/subcategory", subcategory_handler.CreateSubcategory)
	router.PUT("/subcategory/:id", subcategory_handler.UpdateSubcategory)
	router.DELETE("/subcategory/:id", subcategory_handler.DeleteSubcategory)
}
