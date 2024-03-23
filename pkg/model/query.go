package model

import (
	"fmt"
)

type Query struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Count int    `json:"count"`
}

func (q Query) String() string {
	return fmt.Sprintf("%d-%d-%d-%s-%s", q.Int1, q.Int2, q.Limit, q.Str1, q.Str2)
}
