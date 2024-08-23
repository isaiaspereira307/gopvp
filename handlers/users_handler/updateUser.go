package handlers

import (
	"net/http"
	"strconv"

	"github.com/isaiaspereira307/gopvp/internal/db"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Update an user
// @Description Update an user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body UpdateUserRequest true "Update User Request"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt32, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid id"})
		return
	}

	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request"})
		return
	}

	newUser := db.UpdateUserParams{
		ID:        int32(idInt32),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}
	err = queries.UpdateUser(ctx, newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to update user"})
		return
	}
	user := db.User{
		ID:        int32(idInt32),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	resp := UpdateUserResponse{
		Message: "user updated successfully",
		Data:    user,
	}
	ctx.JSON(http.StatusOK, resp)
}
