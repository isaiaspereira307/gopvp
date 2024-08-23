package routes

import (
	"github.com/gin-gonic/gin"
	subcategory_handler "github.com/isaiaspereira307/gopvp/handlers/subcategory_handler"
)

func InitializeSubcategoryRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/subcategory/:id", subcategory_handler.GetSubcategory)
		v1.GET("/subcategories", subcategory_handler.GetSubcategories)
		v1.POST("/subcategory", subcategory_handler.CreateSubcategory)
		v1.PUT("/subcategory/:id", subcategory_handler.UpdateSubcategory)
		v1.DELETE("/subcategory/:id", subcategory_handler.DeleteSubcategory)
	}
}
