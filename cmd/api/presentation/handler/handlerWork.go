package handler

import (
	hub "crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/model"
	pb "crave/shared/proto/hub"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerWork struct {
	pb.UnimplementedHubServer
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
	createdWorkId, err := h.ctrl.CreateWork(model.WorkFrom(page, origin, destionation, algorithm, step, filter))
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d", createdWorkId))
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
func (h *HandlerWork) PauseWork(c *gin.Context) {
	workId := c.Param("workId")
	workIdUint64, err := strconv.ParseUint(workId, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse work id: %w", err))
	}
	if err := h.ctrl.PauseWork(uint16(workIdUint64)); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to pause work with id: %w", err))
	}
	c.Status(http.StatusOK)
}

func (h *HandlerWork) ContinueWork(c *gin.Context) {
	workId := c.Param("workId")
	workIdUint64, err := strconv.ParseUint(workId, 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse work id: %w", err))
	}
	if err := h.ctrl.ContinueWork(uint16(workIdUint64)); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to continue work with id: %w", err))
	}
	c.Status(http.StatusOK)
}
