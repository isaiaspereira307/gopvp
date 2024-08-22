package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Categories
// @Description Show all Categories
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} ListCategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /categories [get]
func GetCategories(ctx *gin.Context) {
	category, err := queries.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Category"})
		return
	}
	if len(category) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no Category found"})
		return
	}

	ctx.JSON(http.StatusOK, category)
}
