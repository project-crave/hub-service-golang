package main

import (
	hubConfig "crave/hub/configuration"
	"crave/hub/lib"
	"crave/shared/configuration"
	"crave/shared/database"

	"github.com/gin-gonic/gin"
)

func main() {
	variable := configuration.NewVariable()
	database.ConnectDatabase(&variable.Database)
	router := gin.Default()
	container := hubConfig.NewContainer(&variable.HubVariable, &database.DB, router)
	go startApiLib(container)
	router.Run(":3000")
}

func startApiLib(container *hubConfig.Container) {
	go lib.Start(container)
}
