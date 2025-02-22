package handler

import (
	hub "crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerWork struct {
	ctrl hub.IController
}

func NewHandlerWork(ctrl hub.IController) *HandlerWork {
	return &HandlerWork{ctrl: ctrl}
}

func (h *HandlerWork) CreateWork(c *gin.Context) {
	page := c.Query("page")
	origin := c.Query("origin")
	destionation := c.Query("destination")
	algorithm := c.Query("algorithm")
	step := c.Query("step")
	filter := c.Query("filter")
	h.ctrl.CreateWork(model.WorkFrom(page, origin, destionation, algorithm, step, filter))
	c.Status(http.StatusOK)
}
