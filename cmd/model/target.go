package model

type Target struct {
	WorkId   uint16 `gorm:"primaryKey;uniqueIndex:idx_target"`
	Previous uint   `gorm:"primaryKey"`
	Name     string `gorm:"primaryKey;type:varchar(20);uniqueIndex:idx_target"`
	Id       uint64 `gorm:"not null"`
	Done     bool   `gorm:"default:false"`
	Priority int    `gorm:"uniqueIndex:idx_target;not null"`
}

func (Target) TableName() string {
	return "target"
}
