package factory

import (
	_interface "github.com/TalalM/le-bon-fizz-buzz/pkg/interface"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/middleware"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/repository"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/service"
)

type Factory struct {
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) FizzBuzzService() _interface.FizzBuzzService {
	memoryRepo := f.MemoryRepo()
	return service.NewFizzBuzzService(memoryRepo)
}

func (f *Factory) FizzBuzzServiceWithStatsMiddleware() _interface.FizzBuzzService {
	memoryRepo := f.MemoryRepo()
	svc := service.NewFizzBuzzService(memoryRepo)
	statsMiddleware := middleware.NewStatisticsMiddleware(svc, memoryRepo)
	return statsMiddleware
}

func (f *Factory) MemoryRepo() *repository.MemoryRepository {
	return repository.NewMemoryRepository()
}
