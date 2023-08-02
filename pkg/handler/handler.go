package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/events-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("api", h.userIdentity)
	{
		events := api.Group("events")
		{
			events.GET("/", h.GetAll)
			events.POST("/", h.Upsert)
			events.GET("/:id", h.GetById)
			events.DELETE("/:id", h.Delete)
		}
	}

	return router
}
