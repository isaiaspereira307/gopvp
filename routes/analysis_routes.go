package routes

import (
	"github.com/gin-gonic/gin"
	analysis_handler "github.com/isaiaspereira307/gopvp/handlers/analysis_handler"
)

func InitializeAnalysisRoutes(router *gin.RouterGroup) {
	router.GET("/analysis/:id", analysis_handler.GetAnalysisByUserId)
}
