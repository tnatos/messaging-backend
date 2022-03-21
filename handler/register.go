package handler

import (
	"log"
	"messaging-backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerData struct {
	username string `json:"username" binding:"required"`
	email    string `json:"email" binding:"required,email"`
	password string `json:"password" binding:"required,gte=6,lte=20"`
}

func (h *Handler) Register(ctx *gin.Context) {

	var body registerData

	// TODO: Separate into different method
	// Check if json body is valid
	if ctx.ContentType() != "application/json" {
		log.Print("Received unsupported media type")
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "only accepts application/json",
		})
		return
	}

	if err := ctx.ShouldBind(body); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json body",
		})
		return
	}

	user := &model.User{
		Username: body.username,
		Email:    body.email,
		Password: body.password,
	}

	err := h.UserService.Register(ctx.Request.Context(), *user)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
