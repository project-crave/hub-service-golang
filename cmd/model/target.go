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

// type TargetStack struct {
// 	Id       uint   `gorm:"primaryKey;autoIncrement"`
// 	Work_id  uint   `gorm:"primaryKey;uniqueIndex:idx_target_stack"`
// 	Name     string `gorm:"uniqueIndex:idx_target_stack;type:varchar(20);not null"`
// 	Previous uint   `gorm:"uniqueIndex:idx_target_stack;not null"`
// 	Done     bool   `gorm:"default:false"`
// }

// func (TargetStack) TableName() string {
// 	return "target_stack"
// }

// func (ts *TargetStack) BeforeCreate(tx *gorm.DB) error {
// 	newId, err := beforeCreate(tx, ts.TableName(), ts.Work_id)
// 	ts.Id = newId
// 	return err
// }

// type TargetQueue struct {
// 	Id       uint   `gorm:"primaryKey;autoIncrement"`
// 	Work_id  uint   `gorm:"primaryKey;uniqueIndex:idx_target_queue"`
// 	Name     string `gorm:"type:varchar(20);not null;uniqueIndex:idx_target_queue"`
// 	Previous uint   `gorm:"uniqueIndex:idx_target_queue;not null"`
// 	Done     bool   `gorm:"default:false"`
// }

// func (TargetQueue) TableName() string {
// 	return "target_queue"
// }

// func (tq *TargetQueue) BeforeCreate(tx *gorm.DB) error {
// 	newId, err := beforeCreate(tx, tq.TableName(), tq.Work_id)
// 	tq.Id = newId
// 	return err
// }
