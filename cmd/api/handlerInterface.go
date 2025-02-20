package api

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Create(c *gin.Context)
}
