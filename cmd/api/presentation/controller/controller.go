package controller

import (
	"crave/hub/cmd/api/domain/service"
	hub "crave/hub/cmd/api/domain/service"
	"crave/hub/cmd/model"
	target "crave/hub/cmd/target/cmd/api/domain/service"
	work "crave/hub/cmd/work/cmd/api/domain/service"
	craveModel "crave/shared/model"
	"fmt"
	"math"
	"sync"
)

type Controller struct {
	svc       hub.IService
	workSvc   work.IService
	targetSvc target.IService
}

func NewController(svc service.IService, workSvc work.IService, targetSvc target.IService) *Controller {
	return &Controller{svc: svc, workSvc: workSvc, targetSvc: targetSvc}
}

func (c *Controller) CreateWork(work *model.Work) (uint16, error) {
	savedWork, err := c.workSvc.SaveWork(work)
	if err != nil {
		return 0, err
	}
	org, dest, err := c.createOriginAndDestination(savedWork)
	if err != nil {
		return 0, err
	}
	if err := c.targetSvc.Init(org, dest); err != nil {
		return 0, err
	}
	return savedWork.Id, nil
}

func (c *Controller) createOriginAndDestination(work *model.Work) (*model.Target, *model.Target, error) {

	prio := c.getPriority(&work.Algorithm)

	origin := &model.Target{
		WorkId:   work.Id,
		Previous: 0,
		Name:     work.Origin,
		Id:       math.MinInt64,
		Priority: prio,
		Status:   1,
	}
	destination := &model.Target{
		WorkId:   work.Id,
		Previous: 0,
		Name:     work.Destination,
		Id:       math.MaxInt64,
		Priority: prio,
		Status:   1,
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
	go c.workSvc.UpdateStatus(work, model.PROCESSING)
	if work.Step != craveModel.Dual {
		go c.beginSingleStepsWork(work)
		return nil
	}
	go c.beginDualStepsWork(work)
	return nil
}

func (c *Controller) beginSingleStepsWork(work *model.Work) {
	minedTarget, err := c.targetSvc.MineFirstTarget(work)
	if err != nil {
		return
	}
	for {
		if work.Status != model.PROCESSING {
			break
		}
		minedTarget, err = c.mineNext(work, minedTarget)
		if err != nil {
			break
		}
	}
}

func (c *Controller) processFirstStep(work *model.Work, step craveModel.Step, wg *sync.WaitGroup, resultChan chan<- *model.Target, errChan chan<- error) {
	defer wg.Done()
	workCopy := *work
	workCopy.Step = step
	minedTarget, err := c.mineFirstTarget(&workCopy)
	if err != nil {
		errChan <- fmt.Errorf("error mining first target for %v: %w", step, err)
		return
	}
	resultChan <- minedTarget
}

func (c *Controller) mineFirstTarget(work *model.Work) (*model.Target, error) {
	return c.targetSvc.MineFirstTarget(work)
}

func (c *Controller) processLoopStep(work *model.Work, step craveModel.Step, previous *model.Target, errChan chan<- error) {

	for work.Status == model.PROCESSING {
		workCopy := *work
		workCopy.Step = step
		minedTarget, err := c.mineNext(&workCopy, previous)
		if err != nil {
			errChan <- fmt.Errorf("error mining next target for %v: %w", step, err)
			break
		}
		previous = minedTarget
	}
}

func (c *Controller) mineNext(work *model.Work, previous *model.Target) (*model.Target, error) {
	target, err := c.targetSvc.GetNextTarget(work, previous)
	if err != nil {
		return nil, err
	}

	targetNames, err := c.targetSvc.SendRefineRequest(target, work.Page, work.Step, work.Filter)
	if err != nil {
		return nil, err
	}

	c.targetSvc.SaveTargets(targetNames, target, work.Step)

	return target, nil
}

func (c *Controller) waitProcess(wg *sync.WaitGroup, errChan chan<- error) {
	wg.Wait()
	close(errChan)
}

func (c *Controller) beginDualStepsWork(work *model.Work) {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)
	frontChan := make(chan *model.Target, 1)
	backChan := make(chan *model.Target, 1)
	{
		wg.Add(2)
		go c.processFirstStep(work, craveModel.Back, &wg, backChan, errChan)
		go c.processFirstStep(work, craveModel.Front, &wg, frontChan, errChan)
		go c.waitProcess(&wg, errChan)

		for err := range errChan {
			fmt.Errorf("🛑Error encountered: %v", err)
			return
		}

	}

	{
		errChan := make(chan error, 2)
		wg.Add(2)
		go c.processLoopStep(work, craveModel.Back, <-backChan, errChan)
		go c.processLoopStep(work, craveModel.Front, <-frontChan, errChan)
		for err := range errChan {
			fmt.Errorf("🛑Error encountered: %v", err)
			return
		}
	}
	return
}

func (c *Controller) PauseWork(workId uint16) error {
	return c.workSvc.PauseWork(workId)
}

func (c *Controller) ContinueWork(workId uint16) error {
	work, err := c.workSvc.GetWork(workId)
	if err != nil {
		return err
	}
	go c.workSvc.UpdateStatus(work, model.PROCESSING)
	if work.Step != craveModel.Dual {
		go c.beginSingleStepsWork(work)
		return nil
	}
	go c.beginDualStepsWork(work)
	return nil
}
