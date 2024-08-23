package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/docs"
	"github.com/isaiaspereira307/gopvp/handlers"
	"github.com/isaiaspereira307/gopvp/internal/db"
	"github.com/isaiaspereira307/gopvp/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine, queries *db.Queries) {
	handlers.InitializeHandlers(queries)
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	InitializeAuthRoutes(router)
	protected := router.Group(basePath)
	protected.Use(middleware.AuthMiddleware())
	{
		InitializeUserRoutes(protected)
		InitializeCategoryRoutes(protected)
		InitializeObjectiveRoutes(protected)
		InitializeSubcategoryRoutes(protected)
	}
}
