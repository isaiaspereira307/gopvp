package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Delete an Subcategory
// @Description Delete an Subcategory
// @Tags subcategory
// @Accept json
// @Produce json
// @Param id query string true "Delete Subcategory Param"
// @Success 200 {object} DeleteSubcategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /subcategory [delete]
func DeleteSubcategory(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)

	err = queries.DeleteSubcategory(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete Subcategory"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
