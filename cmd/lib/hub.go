package lib

import (
	"context"
	"crave/hub/cmd/configuration"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) error {
	container := configuration.NewContainer(router)

	errChan := make(chan error, 2)
	go func() {
		if err := startApiServer(container); err != nil {
			errChan <- fmt.Errorf("API Server error: %w", err)
		}
	}()
	go func() {
		if err := startGrpcServer(container); err != nil {
			errChan <- fmt.Errorf("gRPC Server error: %w", err)
		}
	}()

	return <-errChan
}
func startApiServer(container *configuration.Container) error {
	hubGroup := container.Router.Group("/hub")
	{
		hubGroup.POST("/work", container.HubHandler.Create)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", container.Variable.HubApiPort),
		Handler: container.Router,
	}

	fmt.Println("✅ API Server started on port", container.Variable.HubApiPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed : %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\n⚠️  Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	return nil
}

func startGrpcServer(container *configuration.Container) error {
	// lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", container.Variable.HubApiIp, container.Variable.HubApiPort))
	// if err != nil {
	// 	return fmt.Errorf("failed to listen : %d, %w", container.Variable.HubApiPort, err)
	// }
	// s := grpc.NewServer()
	// if servErr := s.Serve(lis); servErr != nil {
	// 	return fmt.Errorf("failed to create server: %w", err)
	// }
	return nil
}
