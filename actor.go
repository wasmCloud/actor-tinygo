package actor

type (
	Handler struct {
		actor    interface{}
		service  string
		dispatch interface{}
	}
	HostError struct {
		message string
	}
)

var allHandlers []Handler

func RegisterHandler(actor interface{}, service string, dispatch interface{}) {
	allHandlers = append(allHandlers, Handler{
		actor:    actor,
		service:  service,
		dispatch: dispatch,
	})
}
