package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/events-api/domain"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var request domain.User

	if err := ctx.BindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid type")
	}

	id, err := h.services.Authorization.CreateUser(request)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) SignIn(ctx *gin.Context) {
	var request SignInInput

	if err := ctx.BindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid type")
	}

	token, err := h.services.Authorization.GenerateToken(request.Username, request.Password)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})

}
