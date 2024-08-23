package handlers

import (
	"net/http"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update an Objective
// @Description Update an Objective
// @Tags objective
// @Accept json
// @Produce json
// @Param id path string true "Objective ID"
// @Param request body UpdateObjectiveRequest true "Update Objective Request"
// @Success 200 {object} UpdateObjectiveResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /objective/{id} [put]
func UpdateObjective(ctx *gin.Context) {
	var req UpdateObjectiveRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request"})
		return
	}

	newObjective := db.UpdateObjectiveParams{
		ID:             req.ID,
		UserID:         req.UserID,
		CategoryID:     req.CategoryID,
		SubcategoryID:  req.SubcategoryID,
		HowAmI:         req.HowAmI,
		HowDoIWantToBe: req.HowDoIWantToBe,
		WhatWillIDo:    req.WhatWillIDo,
		WhenWillIDoIt:  req.WhenWillIDoIt,
	}
	err := queries.UpdateObjective(ctx, newObjective)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to update Objective"})
		return
	}
	resp := UpdateObjectiveResponse{
		Message: "Objective updated successfully",
		Data:    db.Objective(newObjective),
	}
	ctx.JSON(http.StatusOK, resp)
}
