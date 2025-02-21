package configuration

import (
	"crave/shared/configuration"

	"github.com/gin-gonic/gin"
)

func NewContainer(router *gin.Engine) configuration.IContainer {
	ctnr := NewHubWorkContainer(router)
	return ctnr

}