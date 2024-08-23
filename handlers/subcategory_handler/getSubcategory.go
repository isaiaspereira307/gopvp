package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Subcategory
// @Description Show an Subcategory
// @Tags subcategory
// @Accept json
// @Produce json
// @Param id path string true "Show Subcategory Request"
// @Success 200 {object} ShowSubcategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /subcategory/{id} [get]
func GetSubcategory(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)
	Subcategory, err := queries.GetSubcategoryById(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Subcategory"})
		return
	}

	ctx.JSON(http.StatusOK, Subcategory)
}
