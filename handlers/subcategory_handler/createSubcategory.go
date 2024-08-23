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
// @Summary Create an Subcategory
// @Description Create an Subcategory
// @Tags subcategory
// @Accept json
// @Produce json
// @Param request body CreateSubcategoryRequest true "Create Subcategory Params"
// @Success 200 {object} CreateSubcategoryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subcategory [post]
func CreateSubcategory(ctx *gin.Context) {
	var req CreateSubcategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
		return
	}
	subcategoy := db.CreateSubcategoryParams{
		ID:   generateUniqueID(),
		Name: req.Name,
	}
	err := queries.CreateSubcategory(ctx, subcategoy)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
		return
	}

	sendSuccess(ctx, "createSubcategory", req)
}
