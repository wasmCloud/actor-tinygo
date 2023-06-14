//go:build tinygo.wasm || wasm

package actor

import "unsafe"

//go:wasmimport wasmbus __guest_request
func guestRequest(operationPtr unsafe.Pointer, payloadPtr unsafe.Pointer)

//go:wasmimport wasmbus __guest_response
func guestResponse(ptr unsafe.Pointer, len uint32) //nolint

//go:wasmimport wasmbus __guest_error
func guestError(ptr unsafe.Pointer, len uint32)

//go:wasmimport wasmbus __host_call
func hostCall(
	bindingPtr unsafe.Pointer, bindingLen uint32,
	namespacePtr unsafe.Pointer, namespaceLen uint32,
	operationPtr unsafe.Pointer, operationLen uint32,
	payloadPtr unsafe.Pointer, payloadLen uint32) bool

//go:wasmimport wasmbus __host_response_len
func hostResponseLen() uint32

//go:wasmimport wasmbus __host_response
func hostResponse(ptr unsafe.Pointer)

//go:wasmimport wasmbus __host_error_len
func hostErrorLen() uint32

//go:wasmimport wasmbus __host_error
func hostError(ptr unsafe.Pointer)

//go:wasmimport wasmbus __console_log
func consoleLog(str unsafe.Pointer, strLen uint32)
