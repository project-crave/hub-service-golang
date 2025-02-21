package lib

import (
	"crave/shared/configuration"
	"fmt"
)

func Start(cntr configuration.IContainer) error {

	errChan := make(chan error, 2)
	go func() {
		if err := startApiServer(cntr); err != nil {
			errChan <- fmt.Errorf("API Server error: %w", err)
		}
	}()
	go func() {
		if err := startGrpcServer(cntr); err != nil {
			errChan <- fmt.Errorf("gRPC Server error: %w", err)
		}
	}()

	return <-errChan
}
func startApiServer(ctnr configuration.IContainer) error {
	return ctnr.DefineRoute()
}

func startGrpcServer(ctnr configuration.IContainer) error {
	return ctnr.DefineGrpc()
	// lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", container.Variable.HubApiIp, container.Variable.HubApiPort))
	// if err != nil {
	// 	return fmt.Errorf("failed to listen : %d, %w", container.Variable.HubApiPort, err)
	// }
	// s := grpc.NewServer()
	// if servErr := s.Serve(lis); servErr != nil {
	// 	return fmt.Errorf("failed to create server: %w", err)
	// }
}
