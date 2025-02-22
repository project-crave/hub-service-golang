package model

import (
	craveModel "crave/shared/model"
)

type Work struct {
	Id          uint `"gorm:"primaryKey":not null`
	Page        craveModel.Page
	Origin      string
	Destination string
	Algorithm   craveModel.Algorithm
	Step        craveModel.Step
	Filter      []craveModel.Filter
}

func WorkFrom(page, org, dest, algo, step, filter string) *Work {
	return &Work{
		Page:        craveModel.PageFromString(page),
		Origin:      org,
		Destination: dest,
		Algorithm:   craveModel.AlgorithmFromString(algo),
		Step:        craveModel.StepFromString(step),
		Filter:      craveModel.FilterFromString(filter),
	}
}

type WorkCache struct {
	Page      craveModel.Page
	Algorithm craveModel.Algorithm
	Step      craveModel.Step
	Filter    []craveModel.Filter
}

func (work *Work) ToCache() *WorkCache {
	return &WorkCache{
		Page:      work.Page,
		Algorithm: work.Algorithm,
		Step:      work.Step,
		Filter:    append([]craveModel.Filter(nil), work.Filter...),
	}
}

func (wc *WorkCache) Copy() *WorkCache {
	copiedValue := &WorkCache{
		Page:      wc.Page,
		Algorithm: wc.Algorithm,
		Step:      wc.Step,
		Filter:    make([]craveModel.Filter, len(wc.Filter)),
	}
	copy(copiedValue.Filter, wc.Filter)
	return copiedValue
}
