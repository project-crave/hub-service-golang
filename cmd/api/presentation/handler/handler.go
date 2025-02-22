package handler

import (
	hub "crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctrl hub.IController
}

func NewHandler(ctrl hub.IController) *Handler {

	return &Handler{ctrl: ctrl}
}

func (h *Handler) CreateWork(c *gin.Context) {
	input := model.Work{}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.Status(http.StatusBadRequest)
	}
	h.ctrl.CreateWork(&input)
	c.Status(http.StatusOK)
}
