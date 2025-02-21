package configuration

// import (
// 	"crave/hub/cmd/api/domain/service"
// 	"crave/hub/cmd/api/presentation/controller"
// 	"crave/hub/cmd/api/presentation/handler"
// 	"crave/hub/cmd/work/cmd/api/infrastructure/repository"
// 	"crave/miner/cmd/api"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type HubContainer struct {
// 	Variable      *Variable
// 	Router        *gin.Engine
// 	HubHandler    handler.IHandler
// 	HubController controller.IController
// 	HubService    service.IService
// 	HubRepository repository.IRepository
// }

// func NewHubContainer(router *gin.Engine) *HubContainer {
// 	ctnr := &HubContainer{}
// 	ctnr.InitVariable()
// 	ctnr.SetRouter(router)
// 	ctnr.InitDependency(nil)
// 	return ctnr
// }

// func (ctnr *HubContainer) SetRouter(router any) {
// 	ctnr.Router = router.(*gin.Engine)
// }

// func (ctnr *HubContainer) GetIp() string {
// 	return ctnr.Variable.Api.Ip
// }

// func (ctnr *HubContainer) GetPort() uint16 {
// 	return ctnr.Variable.Api.Port
// }

// func (ctnr *HubContainer) GetHttpHandler() http.Handler {
// 	return ctnr.Router
// }

// func (ctnr *HubContainer) InitDependency(param any) {
// 	ctnr.HubRepository = repository.NewRepository()
// 	ctnr.HubService = api.NewService(ctnr.HubRepository)
// 	ctnr.HubController = api.NewController(ctnr.HubService)
// 	ctnr.HubHandler = api.NewHandler(ctnr.HubController)
// }

// func (ctnr *HubContainer) InitVariable() {
// 	ctnr.Variable = NewVariable()
// }
