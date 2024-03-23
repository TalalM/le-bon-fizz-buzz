package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TalalM/le-bon-fizz-buzz/pkg/factory"
	_interface "github.com/TalalM/le-bon-fizz-buzz/pkg/interface"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/transport"
	"github.com/TalalM/le-bon-fizz-buzz/pkg/utils"
	"github.com/go-kit/log"
)

var (
	f      = factory.New()
	svc    _interface.FizzBuzzService
	logger = log.NewJSONLogger(os.Stdout)
)

func init() {
	svc = f.FizzBuzzServiceWithStatsMiddleware()
}

//	@title			Le Bon Fizz Buzz API
//	@version		1.0
//	@description	This is a fizz buzz server for Le Bon Coin.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	contact@lebonfizzbuzz.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		le-bon-fizz-buzz-7dca32904045.herokuapp.com
//	@BasePath	/v1

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	flag.Parse()

	options := transport.RouterOptions{
		Svc: svc,
	}

	router := transport.NewRouter(options, logger).FizzBuzzRouter("v1")
	_ = router.Run(fmt.Sprintf(":%s", utils.GetEnv("PORT", "80")))
}
