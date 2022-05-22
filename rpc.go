package actor

/// Message contains an rpc operation name and a binary (serialized) payload
type Message struct {
	Method string
	Arg    []byte
}

/// Context object passed through rpc methods
type Context struct{}

/// Transport interface used to send rpc messages over some connection
type Transport struct {
	binding   string
	namespace string
}

// ToProvider constructs a Transport for actor-to-provider calls
func ToProvider(contractId string, linkName string) Transport {
	return Transport{
		binding:   linkName,
		namespace: contractId,
	}
}

// ToActor constructs a Transport for actor-to-actor calls
func ToActor(actor_id string) Transport {
	return Transport{
		binding:   "",
		namespace: actor_id,
	}
}

// Send sends the rpc Message using a Transport
func (t *Transport) Send(ctx *Context, msg Message) ([]byte, error) {
	r, ok := HostCall(t.binding, t.namespace, msg.Method, msg.Arg)
	return r, ok
}

/// RpcError is an error type emitted by the rpc infrastructure
type RpcError struct {
	kind    string
	message string
}

/// NewRpcError constructs and RpcError
func NewRpcError(kind string, message string) *RpcError {
	return &RpcError{kind: kind, message: message}
}

/// Error is RpcError's implementation of the error interface
func (e *RpcError) Error() string {
	return e.message
}

//
// additional data types used in Smithy models
//

/// Timestamp identifies an instant in time
type Timestamp struct {
	// the number of non-leap seconds since unix epoch in UTC
	Sec int64
	// the number of nanoseconds since the beginning of the last whole non-leap second
	Nsec uint32
}

/// Document is an 'any' type interface and is not currently implemented for TinyGo actors
type Document struct{}

// ConsoleLog sends log message to host console.
// For internal use - Actors should use the logging capability provider for logging
// Deprecated - support for this function may go away
func ConsoleLog(msg string) {
	consoleLog(stringToPointer(msg), uint32(len(msg)))
}
