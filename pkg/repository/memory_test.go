package repository

import (
	"testing"

	"github.com/TalalM/le-bon-fizz-buzz/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNoQuerySent(t *testing.T) {
	repo := NewMemoryRepository()

	query := repo.GetMostFrequentQuery()

	assert.Nil(t, query)
}

func TestOneQuerySent(t *testing.T) {
	repo := NewMemoryRepository()

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	repo.IncrementQuery(query)
	mostFrequentQuery := repo.GetMostFrequentQuery()

	assert.NotNil(t, mostFrequentQuery)
	assert.Equal(t, query.String(), mostFrequentQuery.String())
	assert.Equal(t, 1, mostFrequentQuery.Count)
}

func TestTwoQueriesSent(t *testing.T) {
	repo := NewMemoryRepository()

	query1 := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	// Send the query twice, it becomes the most frequent query
	repo.IncrementQuery(query1)
	repo.IncrementQuery(query1)
	mostFrequentQuery := repo.GetMostFrequentQuery()

	assert.NotNil(t, mostFrequentQuery)
	assert.Equal(t, query1.String(), mostFrequentQuery.String())
	assert.Equal(t, 2, mostFrequentQuery.Count)

	query2 := model.Query{
		Int1:  5,
		Int2:  7,
		Limit: 20,
		Str1:  "abc",
		Str2:  "def",
	}
	// Send the query twice: it shouldn't change the most frequent query
	repo.IncrementQuery(query2)
	repo.IncrementQuery(query2)
	mostFrequentQuery = repo.GetMostFrequentQuery()

	assert.NotNil(t, mostFrequentQuery)
	assert.Equal(t, query1.String(), mostFrequentQuery.String())
	assert.Equal(t, 2, mostFrequentQuery.Count)

	// Send the query a third time: it should become the most frequent one
	repo.IncrementQuery(query2)
	mostFrequentQuery = repo.GetMostFrequentQuery()

	assert.NotNil(t, mostFrequentQuery)
	assert.Equal(t, query2.String(), mostFrequentQuery.String())
	assert.Equal(t, 3, mostFrequentQuery.Count)
}
