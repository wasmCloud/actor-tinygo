package actor

import (
	"reflect"
	"strings"
	"unsafe"
)

/// ServiceDispatch defines the interface that all Receivers implement
type ServiceDispatch interface {
	// dispatch calls the actor's registered handler for the operation
	Dispatch(ctx *Context, actor interface{}, message *Message) (*Message, error)
}

// Handler used to invoke callbacks when actor receives a message
type Handler struct {
	service  string
	dispatch ServiceDispatch
	actor    interface{}
}

// NewHandler constructor - called by service (generated interface) during actor initialization
// Not used directly by actors. Actors should use RegisterHandlers
func NewHandler(service string, dispatch ServiceDispatch, actor interface{}) Handler {
	return Handler{service: service, dispatch: dispatch, actor: actor}
}

var allHandlers []Handler

// RegisterHandlers is called by actors during main()
// Example:
// ```
// me = MyActor{}
// actor.RegisterHandlers(me, actor.Handler(), httpserver.Handler())
// ```
func RegisterHandlers(handlers ...Handler) {
	for _, h := range handlers {
		allHandlers = append(allHandlers, h)
	}
}

func fail(errorMessage string) bool {
	guestError(stringToPointer(errorMessage), uint32(len(errorMessage)))
	return false
}

//go:export __guest_call
func guestCall(operationSize uint32, payloadSize uint32) bool { //nolint
	operation := make([]byte, operationSize) // alloc
	payload := make([]byte, payloadSize)     // alloc
	guestRequest(bytesToPointer(operation), bytesToPointer(payload))

	op := string(operation)

	splits := strings.SplitN(op, ".", 2)
	if len(splits) < 2 {
		return fail("invalid operation: " + op)
	}
	ctx := Context{}
	service := splits[0]
	method := splits[1]
	message := Message{Method: method, Arg: payload}
	for _, handler := range allHandlers {
		if handler.service == service {
			msg, err := handler.dispatch.Dispatch(&ctx, handler.actor, &message)
			if err != nil {
				return fail(op + ": " + err.Error())
			}
			guestResponse(bytesToPointer(msg.Arg), uint32(len(msg.Arg)))
			return true
		}
	}
	return fail(op + ": No handler registered")
}

//go:inline
func bytesToPointer(s []byte) uintptr {
	return (*(*reflect.SliceHeader)(unsafe.Pointer(&s))).Data
}

//go:inline
func stringToPointer(s string) uintptr {
	return (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data
}
