package handler

import (
	api "crave/shared/common/api"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	api.IHandler
	CreateWork(c *gin.Context)
}
