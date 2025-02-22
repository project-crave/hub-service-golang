package configuration

import (
	"crave/hub/cmd/api/domain/service"
	"crave/hub/cmd/api/infrastructure/repository"
	"crave/hub/cmd/api/presentation/controller"
	"crave/hub/cmd/api/presentation/handler"
	"crave/hub/cmd/model"
	target "crave/hub/cmd/target/cmd/configuration"
	work "crave/hub/cmd/work/cmd/configuration"
	"crave/shared/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	return nil
}

func (ctnr *HubWorkContainer) DefineDatabase() error {
	ctnr.MysqlWrapper = database.ConnectMysqlDatabase(ctnr.Variable.Database)

	if err := ctnr.MysqlWrapper.Driver.AutoMigrate(&model.Target{}); err != nil {
		return err
	}
	return nil
}

func (ctnr *HubWorkContainer) DefineRoute() error {
	userGroup := ctnr.Router.Group("/hub")
	{
		userGroup.POST("/works", ctnr.HubHandler.CreateWork)
	}
	ctnr.Router.Run(fmt.Sprintf(":%d", ctnr.Variable.Api.Port))
	return nil
}

func (ctnr *HubWorkContainer) GetHttpHandler() http.Handler {
	return ctnr.Router
}

func (ctnr *HubWorkContainer) InitDependency(database any) error {
	ctnr.WorkContainer = work.NewHubWorkContainer()
	ctnr.TargetContainer = target.NewHubWorkTargetContainer()
	ctnr.MinerClient = externalApi.NewMinerClient(fmt.Sprintf("http://%s:%d", ctnr.Variable.Dependency.MinerGrpc.Ip, ctnr.Variable.Dependency.MinerGrpc.Port), ctnr.newGrpcMinerClient())
	ctnr.HubService = service.NewService(ctnr.MinerClient, ctnr.TargetContainer.TargetRepository)
	ctnr.HubController = controller.NewController(ctnr.HubService, ctnr.WorkContainer.WorkService, ctnr.TargetContainer.TargetService)
	ctnr.HubHandler = handler.NewHandlerWork(ctnr.HubController)
}

func NewHubWorkContainer(router *gin.Engine) *HubWorkContainer {
	ctnr := &HubWorkContainer{}
	ctnr.InitVariable()
	ctnr.DefineDatabase()
	ctnr.InitDependency(nil)
	ctnr.SetRouter(router)
	return ctnr
}
