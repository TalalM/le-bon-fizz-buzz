package _interface

import "github.com/TalalM/le-bon-fizz-buzz/pkg/model"

type QueryCounterRepository interface {
	IncrementQuery(query model.Query) bool
	GetMostFrequentQuery() *model.Query
}
