package cache

import (
	"crave/hub/cmd/model"
	"time"
)

type ICache interface {
	Set(key string, value *model.Work, ttl time.Duration)
	Get(key string) (model.WorkCache, bool)
	Delete(key string) error
}
