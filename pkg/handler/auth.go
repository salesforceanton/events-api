package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/events-api/domain"
)

// @Summary SignUp
// @Tags auth
// @Description Register a new User in the system
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body domain.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description Login via Username and Password credentials
// @ID login
// @Accept  json
// @Produce  json
// @Param input body SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /auth/sign-in [post]
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
