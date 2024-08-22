package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/isaiaspereira307/gopvp/configs"
)

var (
	logger *configs.Logger
)

func sendErr(ctx *gin.Context, err error, status int) {
	logger.Errorf("error: %s", err.Error())
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{"error": err.Error()})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"operation": operation, "data": data})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type CreateCategoryResponse struct {
	Message string      `json:"message"`
	Data    db.Category `json:"data"`
}

type DeleteCategoryResponse struct {
	Message string      `json:"message"`
	Data    db.Category `json:"data"`
}

type ShowCategoryResponse struct {
	Message string      `json:"message"`
	Data    db.Category `json:"data"`
}

type UpdateCategoryResponse struct {
	Message string      `json:"message"`
	Data    db.Category `json:"data"`
}

type ListCategoryResponse struct {
	Message string        `json:"message"`
	Data    []db.Category `json:"data"`
}
