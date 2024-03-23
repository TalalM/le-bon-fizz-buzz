package middleware

import (
	_interface "github.com/TalalM/le-bon-fizz-buzz/pkg/interface"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/model"
)

type StatisticsMiddleware struct {
	_interface.FizzBuzzService
	_interface.QueryCounterRepository
}

func NewStatisticsMiddleware(svc _interface.FizzBuzzService, counterRepo _interface.QueryCounterRepository) *StatisticsMiddleware {
	return &StatisticsMiddleware{
		svc,
		counterRepo,
	}
}

func (s *StatisticsMiddleware) FizzBuzz(query model.Query) (string, error) {

	s.QueryCounterRepository.IncrementQuery(query)

	return s.FizzBuzzService.FizzBuzz(query)
}
