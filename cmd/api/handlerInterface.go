package hub

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Default(c *gin.Context)
}
