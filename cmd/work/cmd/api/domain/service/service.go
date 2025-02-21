package service

import (
	"crave/hub/cmd/model"
	"crave/hub/cmd/work/cmd/api/infrastructure/cache"
	"fmt"
	"strings"
)

type Service struct {
	cache cache.ICache
}

func NewService(cache cache.ICache) *Service {
	return &Service{cache: cache}
}

func (s *Service) SaveWork(work *model.Work) (*model.Work, error) {

	s.cache.Set(fmt.Sprintf("%s-%s", work.Origin, work.Destination), work, 0)
	return work, nil
}

func (s *Service) GetWork(key string) (*model.Work, error) {
	workCache, exist := s.cache.Get(key)
	if !exist {
		return nil, fmt.Errorf("👻%s work does not exist", key)
	}
	part := strings.Split(key, "-")
	origin := part[0]
	destination := part[1]
	return &model.Work{
		Page:        workCache.Page,
		Origin:      origin,
		Destination: destination,
		Algorithm:   workCache.Algorithm,
		Step:        workCache.Step,
		Filter:      workCache.Filter,
	}, nil
}
