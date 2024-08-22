package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/docs"
	"github.com/isaiaspereira307/gopvp/handlers"
	category_handler "github.com/isaiaspereira307/gopvp/handlers/category_handler"
	obecjtives_handler "github.com/isaiaspereira307/gopvp/handlers/objectives_handler"
	subcategory_handler "github.com/isaiaspereira307/gopvp/handlers/subcategory_handler"
	users_handlers "github.com/isaiaspereira307/gopvp/handlers/users_handler"
	"github.com/isaiaspereira307/gopvp/internal/db"
	"github.com/isaiaspereira307/gopvp/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine, queries *db.Queries) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/user/:id", middleware.AuthMiddleware(), func(ctx *gin.Context) {
			users_handlers.GetUser(ctx, queries)
		})
		v1.GET("/users", func(ctx *gin.Context) {
			users_handlers.GetUsers(ctx)
		})
		v1.POST("/user", func(ctx *gin.Context) {
			users_handlers.CreateUser(ctx)
		})
		v1.PUT("/user/:id", func(ctx *gin.Context) {
			users_handlers.UpdateUser(ctx, queries)
		})
		v1.DELETE("/user/:id", func(ctx *gin.Context) {
			users_handlers.DeleteUser(ctx, queries)
		})

		v1.GET("/category/:id", func(ctx *gin.Context) {
			category_handler.GetCategory(ctx, queries)
		})
		v1.GET("/categories", func(ctx *gin.Context) {
			category_handler.GetCategories(ctx)
		})
		v1.POST("/category", func(ctx *gin.Context) {
			category_handler.CreateCategory(ctx)
		})
		v1.PUT("/category/:id", func(ctx *gin.Context) {
			category_handler.UpdateCategory(ctx, queries)
		})
		v1.DELETE("/category/:id", func(ctx *gin.Context) {
			category_handler.DeleteCategory(ctx, queries)
		})

		v1.GET("/objective/:id", func(ctx *gin.Context) {
			obecjtives_handler.GetObjective(ctx, queries)
		})
		v1.GET("/objectives", func(ctx *gin.Context) {
			obecjtives_handler.GetObjectives(ctx)
		})
		v1.POST("/objective", func(ctx *gin.Context) {
			obecjtives_handler.CreateObjective(ctx)
		})
		v1.PUT("/objective/:id", func(ctx *gin.Context) {
			obecjtives_handler.UpdateObjective(ctx, queries)
		})
		v1.DELETE("/objective/:id", func(ctx *gin.Context) {
			obecjtives_handler.DeleteObjective(ctx, queries)
		})

		v1.GET("/subcategory/:id", func(ctx *gin.Context) {
			subcategory_handler.GetSubcategory(ctx, queries)
		})
		v1.GET("/subcategories", func(ctx *gin.Context) {
			subcategory_handler.GetSubcategories(ctx)
		})
		v1.POST("/subcategory", func(ctx *gin.Context) {
			subcategory_handler.CreateSubcategory(ctx)
		})
		v1.PUT("/subcategory/:id", func(ctx *gin.Context) {
			subcategory_handler.UpdateSubcategory(ctx, queries)
		})
		v1.DELETE("/subcategory/:id", func(ctx *gin.Context) {
			subcategory_handler.DeleteSubcategory(ctx, queries)
		})

		v1.POST("/login", func(ctx *gin.Context) {
			handlers.Login(ctx, queries)
		})
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
