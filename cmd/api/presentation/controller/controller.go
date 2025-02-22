package controller

import (
	"crave/hub/cmd/api/domain/service"
	hub "crave/hub/cmd/api/domain/service"
	"crave/hub/cmd/model"
	work "crave/hub/cmd/work/cmd/api/domain/service"
)

type Controller struct {
	svc     hub.IService
	workSvc work.IService
}

func NewController(svc service.IService, workSvc work.IService) *Controller {
	return &Controller{svc: svc, workSvc: workSvc}
}

func (c *Controller) SaveWork(work *model.Work) {
	c.workSvc.SaveWork(work)
}
