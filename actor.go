package actor

type (
	Handler   func(payload []byte) ([]byte, error)
	Handlers  map[string]Handler
	HostError struct {
		message string
	}
)

var (
	allHandlers = Handlers{}
)

func RegisterHandlers(handlers Handlers) {
	for name, fn := range handlers {
		allHandlers[name] = fn
	}
}

func RegisterHandler(name string, fn Handler) {
	allHandlers[name] = fn
}
