package main

import (
	"github.com/wasmcloud/actor-tinygo"
	"github.com/wasmcloud/interfaces/httpserver/tinygo"
)

func main() {
	me := Echo{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me), actor.ActorHandler(&me))
}

type Echo struct{}

func (e *Echo) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte("hello"),
	}
	return &r, nil
}

func (e *Echo) HealthRequest(ctx *actor.Context, arg actor.HealthCheckRequest) (*actor.HealthCheckResponse, error) {
	var r actor.HealthCheckResponse
	r.Healthy = true
	return &r, nil
}
