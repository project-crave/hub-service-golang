package main

import (
	"crave/hub/cmd/configuration"
	"crave/hub/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {

	//router.Use()
	startLib()
}

func startLib() {
	router := gin.Default()
	ctnr := configuration.NewHubWorkContainer(router)
	lib.Start(ctnr)
}
