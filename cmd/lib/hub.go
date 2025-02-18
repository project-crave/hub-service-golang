package lib

import (
	"crave/hub/configuration"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func Start(container *configuration.Container) error {
	if err := startApiServer(container); err != nil {
		return err
	}
	if err := startGrpcServer(container); err != nil {
		return err
	}
	return nil
}
func startApiServer(container *configuration.Container) error {

	hubGroup := container.Router.Group("/hub")
	{
		hubGroup.GET("", container.HubHandler.Default)
	}
	return nil
}

func startGrpcServer(container *configuration.Container) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", container.Variable.HubApiIp, container.Variable.HubApiPort))
	if err != nil {
		return fmt.Errorf("failed to listen : %d, %w", container.Variable.HubApiPort, err)
	}
	s := grpc.NewServer()

	if servErr := s.Serve(lis); servErr != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	return nil
}
