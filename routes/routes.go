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
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine, queries *db.Queries) {
	users_handlers.InitializeUserHandlers(queries)
	category_handler.InitializeCategoryHandlers(queries)
	obecjtives_handler.InitializeObjectiveHandlers(queries)
	subcategory_handler.InitializeSubcategoryHandlers(queries)
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/users", users_handlers.GetUsers)
		v1.GET("/users/:id", users_handlers.GetUser)
		v1.POST("/users", users_handlers.CreateUser)
		v1.PUT("/users/:id", users_handlers.UpdateUser)
		v1.DELETE("/users/:id", users_handlers.DeleteUser)

		v1.GET("/category/:id", category_handler.GetCategory)
		v1.GET("/categories", category_handler.GetCategories)
		v1.POST("/category", category_handler.CreateCategory)
		v1.PUT("/category/:id", category_handler.UpdateCategory)
		v1.DELETE("/category/:id", category_handler.DeleteCategory)

		v1.GET("/objective/:id", obecjtives_handler.GetObjective)
		v1.GET("/objectives", obecjtives_handler.GetObjectives)
		v1.POST("/objective", obecjtives_handler.CreateObjective)
		v1.PUT("/objective/:id", obecjtives_handler.UpdateObjective)
		v1.DELETE("/objective/:id", obecjtives_handler.DeleteObjective)

		v1.GET("/subcategory/:id", subcategory_handler.GetSubcategory)
		v1.GET("/subcategories", subcategory_handler.GetSubcategories)
		v1.POST("/subcategory", subcategory_handler.CreateSubcategory)
		v1.PUT("/subcategory/:id", subcategory_handler.UpdateSubcategory)
		v1.DELETE("/subcategory/:id", subcategory_handler.DeleteSubcategory)

		v1.POST("/login", func(ctx *gin.Context) {
			handlers.Login(ctx, queries)
		})
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
