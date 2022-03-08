package handler

import "github.com/gin-gonic/gin"

// 	Handler:
//	- Holds required services for handler to function
type Handler struct{}

//	Config:
//	- Holds necessary service to be injected into handler layer on handler initialization
type Config struct {
	R *gin.Engine
}

//	NewHandler:
//	- Initialze the handler
//	- Initialze Http Routes
func NewHandler(config *Config) {

	// Create the instance of handler
	h := &Handler{}

	// Gin.Engine
	r := config.R

	// Http Routes
	r.POST("/ping", h.ping)
}

//	Test path to confirm if server is connected
func (h *Handler) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
