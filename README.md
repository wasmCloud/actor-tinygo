# wasmCloud TinyGo Actor SDK
This module is for those looking to develop [wasmCloud](https://wasmcloud.dev) actors using [TinyGo](https://tinygo.org). Typically you won't use this module directly, as it will be referenced by one or more code generate modules containing strongly-typed structs and wrappers around function handler registration.

A typical actor might look like this

```go
package main

import (
	"github.com/wasmcloud/actor-tinygo"
	"github.com/wasmcloud/interfaces/httpserver/tinygo"
)

func main() {
	// In the main function, construct an instance of your actor, and pass in the instance,
	// followed by a handler for each of the interfaces you implement.
	// This actor responsds to http requests from the http server capability provider,
	// so it includes the HttpServerHandler.
	// All actors must include the ActorHandler and implement `HealthRequest`.
	me := Echo{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me), actor.ActorHandler(&me))
}

type Echo struct{}

// HandleRequest implements the callback from the http server capability provider.
func (e *Echo) HandleRequest(ctx *actor.Context, arg httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte("hello"),
	}
	return &r, nil
}

// HealthRequest implements the health check from the wasmcloud host.
func (e *Echo) HealthRequest(ctx *actor.Context, arg actor.HealthCheckRequest) (*actor.HealthCheckResponse, error) {
	var r actor.HealthCheckResponse
	r.Healthy = true
	return &r, nil
}
```