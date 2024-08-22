package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Objective
// @Description Show an Objective
// @Tags objective
// @Accept json
// @Produce json
// @Param id query string true "Show Objective Request"
// @Success 200 {object} ShowObjectiveResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /objective [get]
func GetObjective(ctx *gin.Context, queries *db.Queries) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt64 > math.MaxInt32 || idInt64 < math.MinInt32 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	idInt32 := int32(idInt64)
	Objective, err := queries.GetObjectiveById(ctx, idInt32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Objective"})
		return
	}

	ctx.JSON(http.StatusOK, Objective)
}
