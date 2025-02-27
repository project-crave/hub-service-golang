package configuration

import (
	"crave/hub/cmd/api/domain/service"
	"crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/api/presentation/handler"
	"crave/hub/cmd/target/cmd/api/infrastructure/externalApi"
	target "crave/hub/cmd/target/cmd/configuration"
	work "crave/hub/cmd/work/cmd/configuration"
	"crave/shared/database"
	"fmt"
	"net"
	"net/http"

	hubPb "crave/shared/proto/hub"
	minerPb "crave/shared/proto/miner"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HubWorkContainer struct {
	Variable        *Variable
	Router          *gin.Engine
	MysqlWrapper    *database.MysqlWrapper
	WorkContainer   *work.HubWorkContainer
	TargetContainer *target.HubWorkTargetContainer
	HubHandler      handler.IHandler
	HubController   controller.IController
	HubService      service.IService
	MinerClient     externalApi.IMinerClient
}

func (ctnr *HubWorkContainer) SetRouter(router any) {
	ctnr.Router = router.(*gin.Engine)
}

func (ctnr *HubWorkContainer) InitVariable() error {
	ctnr.Variable = NewVariable()
	return nil
}

func (ctnr *HubWorkContainer) DefineGrpc() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ctnr.Variable.GrpcApi.Ip, ctnr.Variable.GrpcApi.Port))
	if err != nil {
		return fmt.Errorf("failed to listen : %d, %w", ctnr.Variable.GrpcApi.Port, err)
	}
	server := grpc.NewServer()
	hubPb.RegisterHubServer(server, ctnr.HubHandler)
	if servErr := server.Serve(lis); servErr != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}
	return nil
}

func (ctnr *HubWorkContainer) DefineDatabase() error {
	return nil
}

func (ctnr *HubWorkContainer) DefineRoute() error {
	userGroup := ctnr.Router.Group("/hub")
	{
		userGroup.POST("/works", ctnr.HubHandler.CreateWork)
		userGroup.PATCH("/works/:workId/start", ctnr.HubHandler.BeginWork)
		userGroup.PATCH("/works/:workId/pause", ctnr.HubHandler.PauseWork)
		userGroup.PATCH("/works/:workId/continue", ctnr.HubHandler.ContinueWork)
	}
	ctnr.Router.Run(fmt.Sprintf(":%d", ctnr.Variable.Api.Port))
	return nil
}

func (ctnr *HubWorkContainer) GetHttpHandler() http.Handler {
	return ctnr.Router
}

func (ctnr *HubWorkContainer) InitDependency(database any) error {
	ctnr.DefineDatabase()
	ctnr.WorkContainer = work.NewHubWorkContainer()
	ctnr.TargetContainer = target.NewHubWorkTargetContainer()
	ctnr.MinerClient = externalApi.NewMinerClient(fmt.Sprintf("http://%s:%d", ctnr.Variable.Dependency.MinerGrpc.Ip, ctnr.Variable.Dependency.MinerGrpc.Port), ctnr.newGrpcMinerClient())
	ctnr.HubService = service.NewService(ctnr.TargetContainer.TargetRepository)
	ctnr.HubController = controller.NewController(ctnr.HubService, ctnr.WorkContainer.WorkService, ctnr.TargetContainer.TargetService)
	ctnr.HubHandler = handler.NewHandlerWork(ctnr.HubController)
	return nil
}

func (ctnr *HubWorkContainer) newGrpcMinerClient() minerPb.MinerClient {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d",
		ctnr.Variable.Dependency.MinerGrpc.Ip,
		ctnr.Variable.Dependency.MinerGrpc.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}

	return minerPb.NewMinerClient(conn)
}

func NewHubWorkContainer(router *gin.Engine) *HubWorkContainer {
	ctnr := &HubWorkContainer{}
	ctnr.InitVariable()
	ctnr.InitDependency(nil)
	ctnr.SetRouter(router)
	return ctnr
}
