package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message":"Hello World!",
	})
}

func (h *Handler) signIn(c *gin.Context) {
	
}