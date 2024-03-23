package _interface

import "github.com/TalalM/le-bon-fizz-buzz/pkg/model"

type FizzBuzzService interface {
	FizzBuzz(query model.Query) (string, error)
	GetMostRequestedQuery() *model.Query
}
