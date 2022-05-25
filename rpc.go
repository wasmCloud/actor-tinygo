package actor

import (
	cbor "github.com/wasmcloud/tinygo-cbor"
	msgpack "github.com/wasmcloud/tinygo-msgpack"
)

// Message contains an rpc operation name and a binary (serialized) payload
type Message struct {
	Method string
	Arg    []byte
}

// Context object passed through rpc methods
type Context struct{}

// Transport interface used to send rpc messages over some connection
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

// RpcError is an error type emitted by the rpc infrastructure
type RpcError struct {
	kind    string
	message string
}

// NewRpcError constructs an RpcError
func NewRpcError(kind string, message string) *RpcError {
	return &RpcError{kind: kind, message: message}
}

// Error is RpcError's implementation of the error interface
func (e *RpcError) Error() string {
	return e.message
}

//
// additional data types used in Smithy models
//

// Document is an 'any' type interface and is not currently implemented for TinyGo actors
type Document struct{}

func (o *Document) MEncode(encoder msgpack.Writer) error {
	return nil
}
func (o *Document) CEncode(encoder cbor.Writer) error {
	return nil
}
func MDecodeDocument(d *msgpack.Decoder) Document {
	return Document{}
}
func CDecodeDocument(d *cbor.Decoder) Document {
	return Document{}
}

// ConsoleLog sends log message to host console.
// For internal use - Actors should use the logging capability provider for logging
// Deprecated - support for this function may go away
func ConsoleLog(msg string) {
	consoleLog(stringToPointer(msg), uint32(len(msg)))
}

// Timestamp identifies an instant in time
type Timestamp struct {
	// the number of non-leap seconds since unix epoch in UTC
	Sec int64
	// the number of nanoseconds since the beginning of the last whole non-leap second
	Nsec uint32
}

// MEncode serializes a Timestamp using msgpack
func (o *Timestamp) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("sec")
	encoder.WriteInt64(o.Sec)
	encoder.WriteString("nsec")
	encoder.WriteUint32(o.Nsec)
	return encoder.CheckError()
}

// MDecodeTimestamp deserializes a Timestamp using msgpack
func MDecodeTimestamp(d *msgpack.Decoder) (Timestamp, error) {
	var val Timestamp
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "sec":
			val.Sec, err = d.ReadInt64()
		case "nsec":
			val.Nsec, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// CEncode serializes a Timestamp using cbor
func (o *Timestamp) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("sec")
	encoder.WriteInt64(o.Sec)
	encoder.WriteString("nsec")
	encoder.WriteUint32(o.Nsec)
	return encoder.CheckError()
}

// CDecodeTimestamp deserializes a Timestamp using cbor
func CDecodeTimestamp(d *cbor.Decoder) (Timestamp, error) {
	var val Timestamp
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "sec":
			val.Sec, err = d.ReadInt64()
		case "nsec":
			val.Nsec, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}
