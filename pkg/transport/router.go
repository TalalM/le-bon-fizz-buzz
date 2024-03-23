package transport

import (
	"fmt"

	_interface "github.com/TalalM/le-bon-fizz-buzz/pkg/interface"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

type Router struct {
	svc _interface.FizzBuzzService

	logger log.Logger
}

type RouterOptions struct {
	Svc _interface.FizzBuzzService
}

func NewRouter(options RouterOptions, logger log.Logger) *Router {
	return &Router{
		svc: options.Svc,

		logger: logger,
	}
}

func (r Router) FizzBuzzRouter(version string) *gin.Engine {
	engine := gin.Default()

	path := fmt.Sprintf("%s/fizz-buzz/:int1/:int2/:limit/:str1/:str2", version)
	engine.GET(path, func(c *gin.Context) {
		EncodeResponse(MakeFizzBuzzEndpoint, r.svc, c)
	})

	path = fmt.Sprintf("%s/most-requested-query", version)
	engine.GET(path, func(c *gin.Context) {
		EncodeResponse(MakeGetMostRequestedQuery, r.svc, c)
	})

	return engine
}
