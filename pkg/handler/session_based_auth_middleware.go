package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/events-api/pkg/logger"
)

func (h *Handler) checkUserSession(ctx *gin.Context) {
	// Retain session from session store
	session, err := h.sessionStore.Get(ctx.Request, COOKIE_NAME)
	if err != nil || session.IsNew {
		logger.LogHandlerIssue("check-user-session", errors.New("Session is expired or invalid"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "Session is expired or invalid")
		return
	}

	// Retain User context (User Id) from Cookie
	userId, ok := session.Values[SESSION_ID_COOKIE_PARAM]
	if !ok {
		logger.LogHandlerIssue("check-user-session", errors.New("Session is expired or invalid"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "Session is expired or invalid")
		return
	}

	ctx.Set(USER_CTX, userId)
}
