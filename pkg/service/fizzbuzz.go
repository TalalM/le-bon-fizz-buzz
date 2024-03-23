package service

import (
	"strconv"
	"strings"

	_interface "github.com/TalalM/le-bon-fizz-buzz/pkg/interface"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/model"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/utils"
)

type FizzBuzzService struct {
	counterRepo _interface.QueryCounterRepository
}

func NewFizzBuzzService(counterRepo _interface.QueryCounterRepository) *FizzBuzzService {
	return &FizzBuzzService{
		counterRepo: counterRepo,
	}
}

// FizzBuzz godoc
//
//	@Summary		Fizz buzzes numbers
//	@Description	Replaces numbers with fizz and buzz (or any other custom string)
//	@Accept			json
//	@Produce		json
//	@Param			int1	path		int		true	"First divider"													default(3)
//	@Param			int2	path		int		true	"Second divider"												default(5)
//	@Param			limit	path		int		true	"Limit to go to"												default(100)
//	@Param			str1	path		string	true	"String that replaces number dividable by the first divider"	default("fizz")
//	@Param			str2	path		string	true	"String that replaces number dividable by the second divider"	default("buzz")
//	@Success		200		{object}	model.Query
//	@Failure		400		{object}	transport.APIResponse
//	@Failure		400		{object}	transport.APIResponse
//	@Router			/fizz-buzz/{int1}/{int2}/{limit}/{str1}/{str2} [get]
func (f *FizzBuzzService) FizzBuzz(query model.Query) (string, error) {

	if query.Limit <= 0 {
		return "", utils.NewErrorLimitLessThanOrEqualToZero()
	}

	if query.Int1 <= 0 || query.Int2 <= 0 {
		return "", utils.NewErrorDividerLessThanOrEqualToZero()
	}

	elements := make([]string, 0)
	for i := 1; i <= query.Limit; i++ {
		fizzed := false
		// We start with an empty string
		element := ""

		if i%query.Int1 == 0 {
			// If the number is dividable by int1, then we add str1
			element = query.Str1
			fizzed = true
		}

		if i%query.Int2 == 0 {
			// If the number is dividable by int2, then we add str2
			element = element + query.Str2
			fizzed = true
		}

		if !fizzed {
			// If the number isn't dividable by int1 or int2, then we just use the number
			element = strconv.Itoa(i)
		}

		elements = append(elements, element)
	}
	return strings.Join(elements, ","), nil
}

// GetMostRequestedQuery godoc
//
//	@Summary		Get most requested query
//	@Description	Returns the most requested query of all time
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Query
//	@Router			/most-requested-query [get]
func (f *FizzBuzzService) GetMostRequestedQuery() *model.Query {
	return f.counterRepo.GetMostFrequentQuery()
}
