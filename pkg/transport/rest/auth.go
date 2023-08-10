package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/events-api/domain"
	"github.com/salesforceanton/events-api/pkg/logger"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary     Registration
// @Tags        Auth
// @Description Register a new User in the system
// @ID          sign-up
// @Accept      json
// @Produce     json
// @Param       input   body      domain.User  true "Account Info"
// @Success     200     {integer} integer 1
// @Failure     400,404 {object}  ErrorResponse
// @Failure     500     {object}  ErrorResponse
// @Router      /auth/sign-up [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var request domain.User

	if err := ctx.BindJSON(&request); err != nil {
		logger.LogHandlerIssue("sign-up", err)
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid type")
		return
	}

	id, err := h.services.Authorization.CreateUser(request)
	if err != nil {
		logger.LogHandlerIssue("sign-up", err)
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// @Summary     Login
// @Tags        Auth
// @Description Login via Username and Password credentials
// @ID          login
// @Accept      json
// @Produce     json
// @Param       input   body     SignInInput  true   "Credentials"
// @Success     200     {string} string       "token"
// @Failure     400,404 {object} ErrorResponse
// @Failure     500     {object} ErrorResponse
// @Router      /auth/sign-in [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	var request SignInInput

	if err := ctx.BindJSON(&request); err != nil {
		logger.LogHandlerIssue("sign-in", err)
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid type")
		return
	}

	token, err := h.services.Authorization.GenerateToken(request.Username, request.Password)
	if err != nil {
		logger.LogHandlerIssue("sign-up", err)
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})

}
