package actor

// import msgpack "github.com/wapc/tinygo-msgpack"

type Transport struct {
	binding   string
	namespace string
}

func (t *Transport) Send(ctx *Context, msg Message) ([]byte, error) {
	r, ok := HostCall(t.binding, t.namespace, msg.Method, msg.Arg)
	return r, ok
}

type Duration struct{}
type Context struct{}
type RpcError struct {
	kind    string
	message string
}

// TODO: move to core
type Timestamp struct {
	// the number of non-leap seconds since unix epoch in UTC
	Sec int64
	// the number of nanoseconds since the beginning of the last whole non-leap second
	Nsec uint32
}

// TODO: move to core
type Document struct{}

// TODO: move to core
type DocumentRef struct{}

// TODO: move to model
type CapabilityContractId string

func NewRpcError(kind string, message string) *RpcError {
	return &RpcError{kind: kind, message: message}
}

func (e *RpcError) Error() string {
	return e.message
}

type Message struct {
	Method string
	Arg    []byte
}

// actor-to-provider calls
func ToProvider(contractId string, linkName string) Transport {
	return Transport{
		binding:   linkName,
		namespace: contractId,
	}
}

// for actor-to-actor calls
func ToActor(actor_id string) Transport {
	return Transport{
		binding:   "",
		namespace: actor_id,
	}
}

/// msgpack deserialize
func Deserialize(b []byte) (interface{}, error) {
	return nil, nil
}

/// msgpack serialize
func Serialize(val interface{}) []byte {
	//var sizer msgpack.Sizer
	//val.Encode(&sizer)
	//buffer := make([]byte, sizer.Len())
	//encoder := msgpack.NewEncoder(buffer)
	//val.Encode(&encoder)

	buffer := make([]byte, 0)
	return buffer
}

//type CborEncoder struct{}
//type CborDecoder struct{}

//func (e *CborEncoder) into_inner() []byte { return nil }

//func CborDecode() (interface{}, error) {
//	return nil, nil
//}

//func NewCborEncoder(alloc bool) CborEncoder {
//	return CborEncoder{}
//}
