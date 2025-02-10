package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
