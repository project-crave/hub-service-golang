package model

const (
	BRIDGE_ID int64 = 0
)

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

func NewTargetFromPrevious(name string, previous *Target) Target {

	return Target{
		WorkId:   previous.WorkId,
		Previous: previous.Id,
		Name:     name,
		Id:       previous.Id,
		Priority: 0,
		Status:   IDLE,
	}
}

func NewBridgeFrom(name string, previous *Target) *Target {
	return &Target{
		WorkId:   previous.WorkId,
		Previous: previous.Id,
		Name:     name,
		Id:       0,
		Status:   IDLE,
		Priority: 1,
	}
}

func NewTargetsFrom(names *[]string, previous *Target) *[]Target {
	var targets []Target
	for _, name := range *names {
		if ValidateName(name) {
			targets = append(targets, NewTargetFromPrevious(name, previous))
		}
	}
	return &targets
}

func ValidateName(name string) bool {
	return len(name) < 20
}
