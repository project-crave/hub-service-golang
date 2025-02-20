package main

import (
	"crave/hub/cmd/configuration"
	"crave/hub/cmd/lib"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//router.Use()
	router := gin.Default()
	go func() {
		defer wg.Done()
		startLib(router)
	}()
	wg.Wait()
}

func startLib(router *gin.Engine) {
	container := configuration.NewContainer(router)
	go lib.Start(container)
	router.Run(fmt.Sprintf(":%d", container.Variable.HubApiPort))
}
