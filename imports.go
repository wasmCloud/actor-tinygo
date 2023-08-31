//go:build !purego && !appengine && !wasm && !tinygo.wasm

package actor

func guestRequest(operationPtr uint32, payloadPtr uint32) {} //nolint

func guestResponse(ptr uint32, len uint32) {}

func guestError(ptr uint32, len uint32) {} //nolint

func hostCall(
	bindingPtr uint32, bindingLen uint32,
	namespacePtr uint32, namespaceLen uint32,
	operationPtr uint32, operationLen uint32,
	payloadPtr uint32, payloadLen uint32) int32 {
	return 1
}

func hostResponseLen() uint32 { return 0 }

func hostResponse(ptr uint32) {}

func hostErrorLen() uint32 { return 0 }

func hostError(ptr uint32) {}

func consoleLog(ptr uint32, sz uint32) {}
