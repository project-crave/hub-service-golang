package service

import (
	"crave/hub/cmd/model"
	api "crave/shared/common/api"
)

type IService interface {
	api.IService
	SaveWork(work *model.Work) (*model.Work, error)
}
