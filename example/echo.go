package echo

import "github.com/wasmcloud/actor-tinygo"

func main() {
	me := Echo{}
	actor.RegisterHandler(me, "HttpServer", HttpServerReceiver{})
	actor.RegisterHandler(me, "Actor", actor.ActorReceiver{})
}

type Echo struct{}

func (e *Echo) HandleRequest(ctx *actor.Context, arg HttpRequest) (*HttpResponse, error) {
	r := HttpResponse{
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
