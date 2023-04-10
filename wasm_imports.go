//go:build tinygo.wasm

package actor

//go:wasmimport wasmbus __guest_request
func guestRequest(operationPtr uintptr, payloadPtr uintptr)

//go:wasmimport wasmbus __guest_response
func guestResponse(ptr uintptr, len uint32) //nolint

//go:wasmimport wasmbus __guest_error
func guestError(ptr uintptr, len uint32)

//go:wasmimport wasmbus __host_call
func hostCall(
	bindingPtr uintptr, bindingLen uint32,
	namespacePtr uintptr, namespaceLen uint32,
	operationPtr uintptr, operationLen uint32,
	payloadPtr uintptr, payloadLen uint32) bool

//go:wasmimport wasmbus __host_response_len
func hostResponseLen() uint32

//go:wasmimport wasmbus __host_response
func hostResponse(ptr uintptr)

//go:wasmimport wasmbus __host_error_len
func hostErrorLen() uint32

//go:wasmimport wasmbus __host_error
func hostError(ptr uintptr)

//go:wasmimport wasmbus __console_log
func consoleLog(str uintptr, strLen uint32)
