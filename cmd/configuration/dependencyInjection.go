package configuration

import (
	hub "crave/hub/api"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Container struct {
	Variable      *Variable
	Router        *gin.Engine
	HubHandler    hub.IHandler
	HubController hub.IController
	HubService    hub.IService
	HubRepository hub.IRepository
}

func NewContainer(variable *Variable, DB *neo4j.Driver, router *gin.Engine) *Container {
	if variable != nil {
		variable = NewVariable()
	}

	repo := hub.NewRepository(DB)
	service := hub.NewService(repo)
	controller := hub.NewController(service)
	handler := hub.NewHandler(controller)
	return &Container{
		Variable:      variable,
		Router:        router,
		HubRepository: repo,
		HubService:    service,
		HubController: service,
		HubHandler:    handler,
	}
}
