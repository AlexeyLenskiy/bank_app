package api

import (
	db "bank-app/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Pin       string `json:"pin"`
	Phone     string `json:"phone"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:        req.Email,
		PasswordHash: req.Password,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Pin:          req.Pin,
		Phone:        req.Phone,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user.ID)
}
