package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AUTH_HEADER = "Authorization"
	USER_CTX    = "user_id"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	authHeader := ctx.GetHeader(AUTH_HEADER)
	if authHeader == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization Header is empty")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization Header is invalid")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Error with a parsing access token")
		return
	}

	ctx.Set(USER_CTX, userId)
}

func (h *Handler) getUserContext(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get(USER_CTX)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "User id is not found")
		return 0, errors.New("User id is not found")
	}

	return userId.(int), nil
}
