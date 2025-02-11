package handler

import (
	"UltimateDeluxeSuperAirlineManagerGOTY/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	document := router.Group("/document")
	{
		document.GET("/:id", h.getDocument)
		document.DELETE("/:id", h.deleteDocument)
		document.PUT("/", h.updateDocument)
		document.POST("/", h.createDocument)
	}
	return router
}
