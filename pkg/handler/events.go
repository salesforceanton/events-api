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

// @Summary     Get all
// @Tags        Events
// @Description Get all events available for current user
// @ID          get-all
// @Accept      json
// @Produce     json
// @Success     200     {array}  domain.Event
// @Failure     400,404 {object} ErrorResponse
// @Failure     500     {object} ErrorResponse
// @Router      /api/events/ [get]
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

// @Summary     Get by Id
// @Tags        Events
// @Description Get Event data by defined Id if current User has access to this Event record
// @ID          get-by-id
// @Accept      json
// @Produce     json
// @Param       id      path     int           true  "Event Id"
// @Success     200     {object} domain.Event
// @Failure     400,404 {object} ErrorResponse
// @Failure     500     {object} ErrorResponse
// @Router      /api/events/{id} [get]
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

// @Summary     Create
// @Tags        Events
// @Description Create Event record with current User as Organizer
// @ID          create
// @Accept      json
// @Produce     json
// @Param       input   body     domain.SaveEventRequest true "Request"
// @Success     201
// @Failure     400,404 {object} ErrorResponse
// @Failure     500     {object} ErrorResponse
// @Router      /api/events/ [post]
func (h *Handler) Create(ctx *gin.Context) {
	var request domain.SaveEventRequest

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

// @Summary     Update
// @Tags        Events
// @Description Update defined Event data if current User has access to this Event record
// @ID          update
// @Accept      json
// @Produce     json
// @Param       id      path     int                     true "Event Id"
// @Param       input   body     domain.SaveEventRequest true "Request"
// @Success     201     {object} domain.Event
// @Failure     400,404 {object} ErrorResponse
// @Failure     500     {object} ErrorResponse
// @Router      /api/events/{id} [post]
func (h *Handler) Update(ctx *gin.Context) {
	var request domain.SaveEventRequest

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

// @Summary     Delete
// @Tags        Events
// @Description Delete Event with defined Id if current User has access to this Event record
// @ID          delete
// @Accept      json
// @Produce     json
// @Param       id      path     int          true "Event Id"
// @Success     200
// @Failure     400,404 {object} ErrorResponse
// @Failure     500     {object} ErrorResponse
// @Router      /api/events/{id} [delete]
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
