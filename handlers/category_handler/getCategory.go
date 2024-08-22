package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Category
// @Description Show a Category
// @Tags category
// @Accept json
// @Produce json
// @Param id query string true "Show Category Request"
// @Success 200 {object} ShowCategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /category [get]
func GetCategory(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)
	Category, err := queries.GetCategoryById(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Category"})
		return
	}

	ctx.JSON(http.StatusOK, Category)
}
