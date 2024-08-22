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
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		// v1.GET("/users/:id", func(ctx *gin.Context) {
		// 	handlers.GetUser(ctx, queries)
		// })
		// v1.GET("/users", func(ctx *gin.Context) {
		// 	handlers.GetUsers(ctx)
		// })
		// v1.POST("/users", func(ctx *gin.Context) {
		// 	handlers.CreateUser(ctx)
		// })
		// v1.PUT("/users/:id", func(ctx *gin.Context) {
		// 	handlers.UpdateUser(ctx, queries)
		// })
		// v1.DELETE("/users/:id", func(ctx *gin.Context) {
		// 	handlers.DeleteUser(ctx, queries)
		// })

		v1.POST("/login", func(ctx *gin.Context) {
			handlers.Login(ctx, queries)
		})
		// v1.GET("/home", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		// 	handlers.Home(ctx)
		// })
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
