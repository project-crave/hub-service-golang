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
	Filter      craveModel.Filter
}
