package handlers

import (
	"net/http"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update an Subcategory
// @Description Update an Subcategory
// @Tags subcategory
// @Accept json
// @Produce json
// @Param id path string true "Subcategory ID"
// @Param request body UpdateSubcategoryRequest true "Update Subcategory Request"
// @Success 200 {object} UpdateSubcategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subcategory/{id} [put]
func UpdateSubcategory(ctx *gin.Context) {
	var req UpdateSubcategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request"})
		return
	}
	newSubcategory := db.UpdateSubcategoryParams{
		ID:   req.ID,
		Name: req.Name,
	}
	err := queries.UpdateSubcategory(ctx, newSubcategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to update Subcategory"})
		return
	}
	resp := UpdateSubcategoryResponse{
		Message: "Subcategory updated successfully",
		Data:    db.Subcategory(newSubcategory),
	}
	ctx.JSON(http.StatusOK, resp)
}
