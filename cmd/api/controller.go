package hub

type Controller struct {
	svc IService
}

func NewController(svc IService) *Controller {
	return &Controller{svc: svc}
}
