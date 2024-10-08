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

type CreateUserResponse struct {
	Message string  `json:"message"`
	Data    db.User `json:"data"`
}

type DeleteUserResponse struct {
	Message string  `json:"message"`
	Data    db.User `json:"data"`
}

type ShowUserResponse struct {
	Message string  `json:"message"`
	Data    db.User `json:"data"`
}

type UpdateUserResponse struct {
	Message string  `json:"message"`
	Data    db.User `json:"data"`
}

type ListUserResponse struct {
	Message string    `json:"message"`
	Data    []db.User `json:"data"`
}
