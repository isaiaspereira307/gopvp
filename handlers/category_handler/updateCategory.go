package handlers

import (
	"net/http"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update an Category
// @Description Update an Category
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param request body UpdateCategoryRequest true "Update Category Request"
// @Success 200 {object} UpdateCategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /category/{id} [put]
func UpdateCategory(ctx *gin.Context) {
	var req UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request"})
		return
	}

	newCategory := db.UpdateCategoryParams{
		ID:   req.ID,
		Name: req.Name,
	}
	err := queries.UpdateCategory(ctx, newCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to update Category"})
		return
	}
	resp := UpdateCategoryResponse{
		Message: "Category updated successfully",
		Data:    db.Category(newCategory),
	}
	ctx.JSON(http.StatusOK, resp)
}
