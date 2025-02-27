package controller

import (
	"crave/hub/cmd/model"
	api "crave/shared/common/api"
)

type IController interface {
	api.IController
	CreateWork(work *model.Work) (uint16, error)
	BeginWork(workId uint16) error
	HandleParsedTargets(name string, targets []string) error
	PauseWork(workId uint16) error
	ContinueWork(workId uint16) error
}
