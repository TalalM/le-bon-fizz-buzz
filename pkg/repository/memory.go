package repository

import "github.com/TalalM/le-bon-fizz-buzz/pkg/model"

type MemoryRepository struct {
	currentMax      int
	currentMaxQuery *model.Query
	queryCounts     map[string]int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		currentMax:  0,
		queryCounts: map[string]int{},
	}
}

func (m *MemoryRepository) IncrementQuery(query model.Query) bool {
	key := query.String()
	count, exists := m.queryCounts[key]
	if !exists {
		m.queryCounts[key] = 0
	}
	newCount := count + 1
	m.queryCounts[key] = newCount

	if newCount > m.currentMax {
		m.currentMax = newCount
		query.Count = newCount
		m.currentMaxQuery = &query
	}

	return true
}

func (m *MemoryRepository) GetMostFrequentQuery() *model.Query {
	return m.currentMaxQuery
}
