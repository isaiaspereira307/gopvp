package routes

import (
	"github.com/gin-gonic/gin"
	objectives_handler "github.com/isaiaspereira307/gopvp/handlers/objectives_handler"
)

func InitializeObjectiveRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/objective/:id", objectives_handler.GetObjective)
		v1.GET("/objectives", objectives_handler.GetObjectives)
		v1.POST("/objective", objectives_handler.CreateObjective)
		v1.PUT("/objective/:id", objectives_handler.UpdateObjective)
		v1.DELETE("/objective/:id", objectives_handler.DeleteObjective)
	}
}
