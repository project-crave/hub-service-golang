package model

import (
	craveModel "crave/shared/model"
)

type Work struct {
	Id          uint16 `"gorm:"primaryKey":not null`
	Page        craveModel.Page
	Origin      string
	Destination string
	Algorithm   craveModel.Algorithm
	Step        craveModel.Step
	Filter      craveModel.Filter
	Status      Status
}

func (Work) TableName() string {
	return "work"
}

func WorkFrom(page, org, dest, algo, step, filter string) *Work {
	return &Work{
		Page:        craveModel.PageFromString(page),
		Origin:      org,
		Destination: dest,
		Algorithm:   craveModel.AlgorithmFromString(algo),
		Step:        craveModel.StepFromString(step),
		Filter:      craveModel.FilterFromString(filter),
		Status:      1,
	}
}

type WorkCache struct {
	Processing  bool
	Page        craveModel.Page
	Origin      string
	Destination string
	Algorithm   craveModel.Algorithm
	Step        craveModel.Step
	Filter      craveModel.Filter
	Status      Status
}

func (work *Work) ToCache() WorkCache {
	return WorkCache{
		Processing:  false,
		Page:        work.Page,
		Origin:      work.Origin,
		Destination: work.Destination,
		Algorithm:   work.Algorithm,
		Step:        work.Step,
		Filter:      work.Filter,
		Status:      work.Status,
	}
}

func (wc *WorkCache) ToWork(id uint16) *Work {
	return &Work{
		Id:          id,
		Page:        wc.Page,
		Origin:      wc.Origin,
		Destination: wc.Destination,
		Algorithm:   wc.Algorithm,
		Step:        wc.Step,
		Filter:      wc.Filter,
		Status:      wc.Status,
	}
}
