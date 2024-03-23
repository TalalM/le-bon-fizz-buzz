# Le bon fizz buzz
This is a fizz buzz exercise example for Le Bon Coin.

What this project does is implement the FizzBuzz game: https://en.wikipedia.org/wiki/Fizz_buzz


# Installation

After installing [go](https://go.dev/doc/install), just run:
```
make server
```

To make sure that everything works as expected, call:

```
http://localhost/v1/fizz-buzz/3/5/100/fizz/buzz
```

## Testing

To run the tests:
```
make test
```

# Code structure

The code is structed like this:
```
- cmd
    - api
        - main.go
- pkg
    - factory
    - interface
    - middleware
    - model
    - repository
    - service
    - transport
    - utils
```

- `cmd/api/main.go` is the main command.
- The `pkg/factory` repository is the project's factory where all the services, middlewares and repositories are created.
- All the services and repositories interfaces should be defined in `pkg/interface`
- The `pkg/middleware`, `pkg/model`, `pkg/repository` and `pkg/service` are were the middleware/model/repository/service code lives.
- The routes are defined in `pkg/transport/router.go`
- The custom errors are defined in `pkg/utils/errors.go`

## Repository

For convenience, the only middleware statistics that has been implemented is an in-memory one.

To use a different one (database, cache, ...), create one that implements the `QueryCounterRepository` interface, and use it in the factory instead of the `MemoryRepo`.

# API documentation

The API documentation is generated with [swag](https://github.com/swaggo/swag), using annotations (see file `pkg/service/fizzbuzz.go`).

From the generated files, only the `swagger.yaml` file is kept. It can be imported to any Swagger UI.

To generate the doc, use:
```
make doc
```

There is a version of the documentation that is hosted [here](https://app.swaggerhub.com/apis/MAZROUITALAL/le-bon_fizz_buzz_api/1.0).

# Live example

There is a live example of this API hosted on `https://le-bon-fizz-buzz-7dca32904045.herokuapp.com/`.

For example, you can call: [https://le-bon-fizz-buzz-7dca32904045.herokuapp.com/v1/fizz-buzz/3/5/100/fizz/buzz](https://le-bon-fizz-buzz-7dca32904045.herokuapp.com/v1/fizz-buzz/3/5/100/fizz/buzz).
