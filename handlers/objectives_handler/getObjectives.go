package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Show Objectives
// @Description Show all Objectives
// @Tags objective
// @Accept json
// @Produce json
// @Success 200 {object} ListObjectiveResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /objectives [get]
func GetObjectives(ctx *gin.Context) {
	Objectives, err := queries.GetObjectives(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get Objective"})
		return
	}
	if len(Objectives) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no Objectives found"})
		return
	}

	ctx.JSON(http.StatusOK, Objectives)
}
