package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/docs"
	"github.com/isaiaspereira307/gopvp/handlers"
	"github.com/isaiaspereira307/gopvp/internal/db"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine, queries *db.Queries) {
	handlers.InitializeHandlers(queries)
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	InitializeUserRoutes(router)
	InitializeCategoryRoutes(router)
	InitializeObjectiveRoutes(router)
	InitializeSubcategoryRoutes(router)
	InitializeAuthRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
