package handler

import (
	hub "crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/model"
	"fmt"
	"net/http"
	"strconv"

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
	createdWorkId, err := h.ctrl.CreateWork(&input)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d", createdWorkId))
}

func (h *Handler) BeginWork(c *gin.Context) {
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
