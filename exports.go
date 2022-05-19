package actor

import (
	"reflect"
	"strings"
	"unsafe"
)

func fail(errorMessage string) bool {
	guestError(stringToPointer(errorMessage), uint32(len(errorMessage)))
	return false
}

//go:export __guest_call
func guestCall(operationSize uint32, payloadSize uint32) bool {
	operation := make([]byte, operationSize) // alloc
	payload := make([]byte, payloadSize)     // alloc
	guestRequest(bytesToPointer(operation), bytesToPointer(payload))

	op := string(operation)
	splits := strings.SplitN(op, ",", 2)
	if len(splits) < 2 {
		return fail("invalid operation: " + op)
	}
	ctx := Context{}
	service := splits[0]
	method := splits[1]
	message := Message{Method: method, Arg: payload}
	for _, handler := range allHandlers {
		if handler.service == service {
			disp, _ := handler.dispatch.(ServiceDispatch)
			rc, err := disp.dispatch(&ctx, handler.actor, message)
			if err != nil {
				return fail(op + ": " + err.Error())
			}
			guestResponse(bytesToPointer(rc.Arg), uint32(len(rc.Arg)))
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
