package controller

import (
	"crave/hub/cmd/model"
	api "crave/shared/common/api"
)

type IController interface {
	api.IController
	CreateWork(work *model.Work) error
	BeginWork(workId uint16) error
}
