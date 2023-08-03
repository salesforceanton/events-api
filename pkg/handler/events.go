package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/events-api/domain"
)

type EventsResponse struct {
	Data []domain.Event
}

func (h *Handler) GetAll(ctx *gin.Context) {
	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := h.services.Events.GetAll(userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(
		http.StatusOK, EventsResponse{result},
	)
}
func (h *Handler) GetById(ctx *gin.Context) {
	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	eventId, err := h.getUrlParam(ctx, "id")
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Invalid param in url: [id]")
		return
	}

	result, err := h.services.Events.GetById(userId, eventId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)

}
func (h *Handler) Create(ctx *gin.Context) {
	var request domain.Event

	if err := ctx.BindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid type")
	}

	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := h.services.Events.Create(userId, request)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"Status": fmt.Sprintf("Event record [id]:%d has been saved successfully", result),
	})
}

func (h *Handler) Update(ctx *gin.Context) {
	var request domain.Event

	if err := ctx.BindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid type")
		return
	}

	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	eventId, err := h.getUrlParam(ctx, "id")
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Invalid param in url: [id]")
		return
	}

	result, err := h.services.Events.Update(userId, eventId, request)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
func (h *Handler) Delete(ctx *gin.Context) {
	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	eventId, err := h.getUrlParam(ctx, "id")
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Invalid param in url: [id]")
		return
	}

	err = h.services.Events.Delete(userId, eventId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Status": fmt.Sprintf("Event record [id]:%d has been deleted successfully", eventId),
	})
}
