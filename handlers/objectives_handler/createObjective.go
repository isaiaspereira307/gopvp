package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/internal/db"
)

var (
	lastID int32
	mu     sync.Mutex
)

func generateUniqueID() int32 {
	mu.Lock()
	defer mu.Unlock()
	lastID++
	return lastID
}

// @BasePath /api/v1
// @Summary Create an Objective
// @Description Create an Objective
// @Tags objective
// @Accept json
// @Produce json
// @Param request body CreateObjectiveRequest true "Create Objective Params"
// @Success 200 {object} CreateObjectiveResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /objective [post]
func CreateObjective(ctx *gin.Context) {
	var req CreateObjectiveRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	err := req.Validate()
	if err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}

	objective := db.CreateObjectiveParams{
		ID:             generateUniqueID(),
		UserID:         req.UserID,
		CategoryID:     req.CategoryID,
		SubcategoryID:  req.SubcategoryID,
		HowAmI:         req.HowAmI,
		HowDoIWantToBe: req.HowDoIWantToBe,
		WhatWillIDo:    req.WhatWillIDo,
		WhenWillIDoIt:  req.WhenWillIDoIt,
	}

	err = queries.CreateObjective(ctx, objective)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createObjective", objective)
}
