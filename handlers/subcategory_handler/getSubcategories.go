package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Subcategories
// @Description Show all Subcategories
// @Tags subcategory
// @Accept json
// @Produce json
// @Success 200 {object} ListUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /subcategories [get]
func GetSubcategories(ctx *gin.Context) {
	subcategories, err := queries.GetSubcategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}
	if len(subcategories) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no Subcategories found"})
		return
	}

	ctx.JSON(http.StatusOK, subcategories)
}
