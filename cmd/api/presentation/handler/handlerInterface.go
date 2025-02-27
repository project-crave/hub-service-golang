package handler

import (
	api "crave/shared/common/api"
	pb "crave/shared/proto/hub"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	api.IHandler
	pb.HubServer
	CreateWork(c *gin.Context)
	BeginWork(c *gin.Context)
	PauseWork(c *gin.Context)
	ContinueWork(c *gin.Context)
}
