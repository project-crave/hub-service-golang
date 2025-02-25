package model

type Target struct {
	WorkId   uint16 `gorm:"primaryKey;uniqueIndex:idx_target"`
	Previous int64  `gorm:"primaryKey"`
	Name     string `gorm:"primaryKey;type:varchar(20);uniqueIndex:idx_target"`
	Id       int64  `gorm:"not null"`
	Status   Status `gorm:"default:0"`
	Priority int    `gorm:"uniqueIndex:idx_target;not null"`
}

func (Target) TableName() string {
	return "target"
}
