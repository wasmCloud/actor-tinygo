package actor

type Transport struct {
	binding   string
	namespace string
}

func (t *Transport) Send(ctx *Context, msg Message) ([]byte, error) {
	r, ok := HostCall(t.binding, t.namespace, msg.Method, msg.Arg)
	return r, ok
}

type Context struct{}

/// All services implement this interfaca
type ServiceDispatch interface {
	dispatch(ctx *Context, actor interface{}, message Message) (*Message, error)
}

type Timestamp struct {
	// the number of non-leap seconds since unix epoch in UTC
	Sec int64
	// the number of nanoseconds since the beginning of the last whole non-leap second
	Nsec uint32
}

type Document struct{}

type RpcError struct {
	kind    string
	message string
}

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
