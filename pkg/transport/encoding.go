package transport

import (
	"strconv"

	_interface "github.com/TalalM/le-bon-fizz-buzz/pkg/interface"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/model"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/utils"
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}

type Endpoint func(svc _interface.FizzBuzzService, c *gin.Context) (interface{}, error)

func MakeFizzBuzzEndpoint(svc _interface.FizzBuzzService, c *gin.Context) (interface{}, error) {
	query := model.Query{
		Int1:  intVal(c.Param("int1")),
		Int2:  intVal(c.Param("int2")),
		Limit: intVal(c.Param("limit")),
		Str1:  c.Param("str1"),
		Str2:  c.Param("str2"),
	}
	return svc.FizzBuzz(query)
}

func MakeGetMostRequestedQuery(svc _interface.FizzBuzzService, c *gin.Context) (interface{}, error) {
	return svc.GetMostRequestedQuery(), nil
}

func EncodeResponse(e Endpoint, svc _interface.FizzBuzzService, c *gin.Context) {
	data, err := e(svc, c)
	response := APIResponse{
		Data:    data,
		Success: err == nil,
	}
	code := 200
	if err != nil {
		response.Error = err.Error()
		switch err.(type) {
		case utils.ErrorDividerLessThanOrEqualToZero, utils.ErrorLimitLessThanOrEqualToZero:
			code = 400
		default:
			code = 500
		}
	}
	c.JSON(code, response)
}

func intVal(value string) int {
	valueInt, _ := strconv.Atoi(value)
	return valueInt
}
