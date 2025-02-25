package controller

import (
	"crave/hub/cmd/api/domain/service"
	hub "crave/hub/cmd/api/domain/service"
	"crave/hub/cmd/model"
	target "crave/hub/cmd/target/cmd/api/domain/service"
	work "crave/hub/cmd/work/cmd/api/domain/service"
	craveModel "crave/shared/model"
	"math"
)

type Controller struct {
	svc       hub.IService
	workSvc   work.IService
	targetSvc target.IService
}

func NewController(svc service.IService, workSvc work.IService, targetSvc target.IService) *Controller {
	return &Controller{svc: svc, workSvc: workSvc, targetSvc: targetSvc}
}

func (c *Controller) CreateWork(work *model.Work) error {
	savedWork, err := c.workSvc.SaveWork(work)
	if err != nil {
		return err
	}
	org, dest, err := c.createOriginAndDestination(savedWork)
	if err != nil {
		return err
	}
	return c.targetSvc.Init(org, dest)
}

func (c *Controller) createOriginAndDestination(work *model.Work) (*model.Target, *model.Target, error) {

	prio := c.getPriority(&work.Algorithm)

	origin := &model.Target{
		WorkId:   work.Id,
		Previous: 0,
		Name:     work.Origin,
		Id:       math.MinInt64,
		Priority: prio,
	}
	destination := &model.Target{
		WorkId:   work.Id,
		Previous: 0,
		Name:     work.Destination,
		Id:       math.MaxInt64,
		Priority: prio,
	}

	return origin, destination, nil
}

func (c *Controller) getPriority(algo *craveModel.Algorithm) int {

	return 0
}

func (c *Controller) BeginWork(workId uint16) error {
	work, err := c.workSvc.GetWork(workId)
	if err != nil {
		return err
	}
	go c.targetSvc.MineFirstTargets(work.Id, work.Algorithm, work.Page, work.Step, work.Filter)
	return nil
}
