package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/configs"
	"github.com/isaiaspereira307/gopvp/internal/db"
)

var (
	logger *configs.Logger
)

func Initialize(queries *db.Queries) {

	corsConfig := configs.GetCORSConfig()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     corsConfig.AllowOrigins,
		AllowMethods:     corsConfig.AllowMethods,
		AllowHeaders:     corsConfig.AllowHeaders,
		ExposeHeaders:    corsConfig.ExposeHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		MaxAge:           time.Duration(corsConfig.MaxAge) * time.Second,
	}))

	InitializeRoutes(r, queries)
	port := fmt.Sprintf(":%s", configs.GetServerPort())

	if err := r.Run(port); err != nil {
		logger.Errorf("Failed to run server: %v", err)
	}
}
