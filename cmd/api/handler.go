package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctrl IController
}

func NewHandler(ctrl IController) *Handler {

	return &Handler{ctrl: ctrl}
}

func (h *Handler) Create(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
