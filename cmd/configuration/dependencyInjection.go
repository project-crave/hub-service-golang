package configuration

import (
	"crave/hub/cmd/api"

	"github.com/gin-gonic/gin"
)

type Container struct {
	Variable      *Variable
	Router        *gin.Engine
	HubHandler    api.IHandler
	HubController api.IController
	HubService    api.IService
	HubRepository api.IRepository
}

func NewContainer(router *gin.Engine) *Container {
	if router == nil {
		router = gin.Default()
	}
	variable := NewVariable()

	repo := api.NewRepository()
	service := api.NewService(repo)
	controller := api.NewController(service)
	handler := api.NewHandler(controller)
	return &Container{
		Variable:      variable,
		Router:        router,
		HubRepository: repo,
		HubService:    service,
		HubController: controller,
		HubHandler:    handler,
	}
}
