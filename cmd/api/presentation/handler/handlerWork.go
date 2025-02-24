package handler

import (
	hub "crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/model"
	"fmt"
	"net/http"
	"strconv"

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

func (h *HandlerWork) BeginWork(c *gin.Context) {
	workId := c.Param("workId")
	workIdUint64, err := strconv.ParseUint(workId, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse work id: %w", err))
	}
	if err := h.ctrl.BeginWork(uint16(workIdUint64)); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to begin work with id: %w", err))
	}
	c.Status(http.StatusOK)
}
