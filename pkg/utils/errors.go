package utils

import "errors"

type ErrorLimitLessThanOrEqualToZero struct {
	error
}

func NewErrorLimitLessThanOrEqualToZero() ErrorLimitLessThanOrEqualToZero {
	return ErrorLimitLessThanOrEqualToZero{
		errors.New("limit should be > 0"),
	}
}

type ErrorDividerLessThanOrEqualToZero struct {
	error
}

func NewErrorDividerLessThanOrEqualToZero() ErrorDividerLessThanOrEqualToZero {
	return ErrorDividerLessThanOrEqualToZero{
		errors.New("dividers should be > 0"),
	}
}
