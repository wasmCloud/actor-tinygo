package actor

import (
	"reflect"
	"unsafe"
)

//go:export __guest_call
func guestCall(operationSize uint32, payloadSize uint32) bool {
	operation := make([]byte, operationSize) // alloc
	payload := make([]byte, payloadSize)     // alloc
	guestRequest(bytesToPointer(operation), bytesToPointer(payload))

	if f, ok := allHandlers[string(operation)]; ok {
		response, err := f(payload)
		if err != nil {
			message := err.Error()
			guestError(stringToPointer(message), uint32(len(message)))

			return false
		}

		guestResponse(bytesToPointer(response), uint32(len(response)))

		return true
	}

	message := `No handler declared for operation "` + string(operation) + `"`
	guestError(stringToPointer(message), uint32(len(message)))

	return false
}

//go:inline
func bytesToPointer(s []byte) uintptr {
	return (*(*reflect.SliceHeader)(unsafe.Pointer(&s))).Data
}

//go:inline
func stringToPointer(s string) uintptr {
	return (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data
}
