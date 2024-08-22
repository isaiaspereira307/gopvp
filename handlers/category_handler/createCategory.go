package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/internal/db"
)

// @BasePath /api/v1
// @Summary Create a Category
// @Description Create a Category
// @Tags category
// @Accept json
// @Produce json
// @Param request body CreateCategoryRequest true "Create Category Params"
// @Success 200 {object} CreateCategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /category [post]
func CreateCategory(ctx *gin.Context) {
	var req CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}
	category := db.CreateCategoryParams{
		Name: req.Name,
	}

	err := queries.CreateCategory(ctx, category)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createCategory", req)
}
