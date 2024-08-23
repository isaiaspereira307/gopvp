package routes

import (
	"github.com/gin-gonic/gin"
	objectives_handler "github.com/isaiaspereira307/gopvp/handlers/objectives_handler"
)

func InitializeObjectiveRoutes(router *gin.RouterGroup) {
	router.GET("/objective/:id", objectives_handler.GetObjective)
	router.GET("/objectives", objectives_handler.GetObjectives)
	router.POST("/objective", objectives_handler.CreateObjective)
	router.PUT("/objective/:id", objectives_handler.UpdateObjective)
	router.DELETE("/objective/:id", objectives_handler.DeleteObjective)
}
