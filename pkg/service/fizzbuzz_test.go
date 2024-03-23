package service

import (
	"testing"

	"github.com/TalalM/le-bon-fizz-buzz/pkg/model"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzSimple(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz", value)
}

func TestFizzBuzzLimitIsLowerThanInts(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  21,
		Int2:  29,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20", value)
}

func TestFizzBuzzBothIntsAreTheSame(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  3,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,fizzbuzz,4,5,fizzbuzz,7,8,fizzbuzz,10,11,fizzbuzz,13,14,fizzbuzz,16,17,fizzbuzz,19,20", value)
}

func TestFizzBuzzInt1Is0(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  0,
		Int2:  3,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, err := svc.FizzBuzz(query)

	assert.NotNil(t, err)
	assert.IsType(t, utils.ErrorDividerLessThanOrEqualToZero{}, err)
	assert.Equal(t, "", value)
}

func TestFizzBuzzInt2Is0(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  0,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, err := svc.FizzBuzz(query)

	assert.NotNil(t, err)
	assert.IsType(t, utils.ErrorDividerLessThanOrEqualToZero{}, err)
	assert.Equal(t, "", value)
}

func TestFizzBuzzInt1IsLessThan0(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  -2,
		Int2:  3,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, err := svc.FizzBuzz(query)

	assert.NotNil(t, err)
	assert.IsType(t, utils.ErrorDividerLessThanOrEqualToZero{}, err)
	assert.Equal(t, "", value)
}

func TestFizzBuzzInt2IsLessThan0(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  -4,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, err := svc.FizzBuzz(query)

	assert.NotNil(t, err)
	assert.IsType(t, utils.ErrorDividerLessThanOrEqualToZero{}, err)
	assert.Equal(t, "", value)
}

func TestFizzBuzzLimitIs0(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 0,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, err := svc.FizzBuzz(query)

	assert.NotNil(t, err)
	assert.IsType(t, utils.ErrorLimitLessThanOrEqualToZero{}, err)
	assert.Equal(t, "", value)
}

func TestFizzBuzzLimitIsNegative(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  7,
		Limit: -5,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, err := svc.FizzBuzz(query)

	assert.NotNil(t, err)
	assert.IsType(t, utils.ErrorLimitLessThanOrEqualToZero{}, err)
	assert.Equal(t, "", value)
}

func TestFizzBuzzInt1Is1(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  1,
		Int2:  3,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz,fizzbuzz,fizz,fizz", value)
}

func TestFizzBuzzInt2Is1(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  5,
		Int2:  1,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "buzz,buzz,buzz,buzz,fizzbuzz,buzz,buzz,buzz,buzz,fizzbuzz,buzz,buzz,buzz,buzz,fizzbuzz,buzz,buzz,buzz,buzz,fizzbuzz", value)
}

func TestFizzBuzzBothIntsAre1(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  1,
		Int2:  1,
		Limit: 20,
		Str1:  "fizz",
		Str2:  "buzz",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz,fizzbuzz", value)
}

func TestFizzBuzzStringsAreNotFizzAndBuzz(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "hello",
		Str2:  "world",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,hello,4,world,hello,7,8,hello,world,11,hello,13,14,helloworld,16,17,hello,19,world", value)
}

func TestFizzBuzzBothStringsAreTheSame(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "hello",
		Str2:  "hello",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,hello,4,hello,hello,7,8,hello,hello,11,hello,13,14,hellohello,16,17,hello,19,hello", value)
}

func TestFizzBuzzOneStringIsEmpty(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "hey",
		Str2:  "",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,hey,4,,hey,7,8,hey,,11,hey,13,14,hey,16,17,hey,19,", value)
}

func TestFizzBuzzBothStringsAreEmpty(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  3,
		Int2:  5,
		Limit: 20,
		Str1:  "",
		Str2:  "",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, "1,2,,4,,,7,8,,,11,,13,14,,16,17,,19,", value)
}

func TestFizzBuzzBothStringsAreEmptyAndBothIntsAre1(t *testing.T) {
	svc := NewFizzBuzzService(nil)

	query := model.Query{
		Int1:  1,
		Int2:  1,
		Limit: 20,
		Str1:  "",
		Str2:  "",
	}
	value, _ := svc.FizzBuzz(query)

	assert.Equal(t, ",,,,,,,,,,,,,,,,,,,", value)
}
