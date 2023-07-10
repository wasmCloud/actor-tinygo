//go:build tinygo.wasm

package actor

//go:wasmimport wasmbus __guest_request
func guestRequest(operationPtr uint32, payloadPtr uint32)

//go:wasmimport wasmbus __guest_response
func guestResponse(ptr uint32, len uint32) //nolint

//go:wasmimport wasmbus __guest_error
func guestError(ptr uint32, len uint32)

//go:wasmimport wasmbus __host_call
func hostCall(
	bindingPtr uint32, bindingLen uint32,
	namespacePtr uint32, namespaceLen uint32,
	operationPtr uint32, operationLen uint32,
	payloadPtr uint32, payloadLen uint32) int32

//go:wasmimport wasmbus __host_response_len
func hostResponseLen() uint32

//go:wasmimport wasmbus __host_response
func hostResponse(ptr uint32)

//go:wasmimport wasmbus __host_error_len
func hostErrorLen() uint32

//go:wasmimport wasmbus __host_error
func hostError(ptr uint32)

//go:wasmimport wasmbus __console_log
func consoleLog(str uint32, strLen uint32)
