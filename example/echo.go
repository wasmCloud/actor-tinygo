package echo

import (
	"github.com/wasmcloud/actor-tinygo"
)

func main() {
	actor.RegisterHandlers(Echo{},
		httpserver.HttpServerHandler(), actor.ActorHandler())
}

type Echo struct{}

func (e *Echo) HandleRequest(ctx *actor.Context, arg httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(HeaderMap, 0),
		Body:       []byte("hello"),
	}
	return &r, nil
}

func (e *Echo) HealthRequest(ctx *actor.Context, arg actor.HealthCheckRequest) (*actor.HealthCheckResponse, error) {
	var r actor.HealthCheckResponse
	r.Healthy = true
	return &r, nil
}
