// wasmcloud platform core data structures
package actor

import (
	"github.com/wapc/tinygo-msgpack" //nolint
)

// List of linked actors for a provider
type ActorLinks []LinkDefinition

func (o *ActorLinks) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.Encode(encoder)
	}

	return nil
}
func DecodeActorLinks(d msgpack.Decoder) (ActorLinks, error) {

	isNil, err := d.IsNextNil()
	if err == nil && isNil {
		d.Skip()
		return make([]LinkDefinition, 0), nil
	}
	size, err := d.ReadArraySize()
	if err != nil {
		size = 0
	}
	val := make([]LinkDefinition, size)
	for i := uint32(0); i < size; i++ {
		item, err := DecodeLinkDefinition(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type ClusterIssuerKey string

func (o *ClusterIssuerKey) Encode(encoder msgpack.Writer) error {
	encoder.WriteString(string(*o))
	return nil
}
func DecodeClusterIssuerKey(d msgpack.Decoder) (ClusterIssuerKey, error) {

	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ClusterIssuerKey(val), nil

}

type ClusterIssuers []ClusterIssuerKey

func (o *ClusterIssuers) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.Encode(encoder)
	}

	return nil
}
func DecodeClusterIssuers(d msgpack.Decoder) (ClusterIssuers, error) {

	isNil, err := d.IsNextNil()
	if err == nil && isNil {
		d.Skip()
		return make([]ClusterIssuerKey, 0), nil
	}
	size, err := d.ReadArraySize()
	if err != nil {
		size = 0
	}
	val := make([]ClusterIssuerKey, size)
	for i := uint32(0); i < size; i++ {
		item, err := DecodeClusterIssuerKey(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil

}

// health check request parameter
type HealthCheckRequest struct {
}

func (o *HealthCheckRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(0)

	return nil
}
func DecodeHealthCheckRequest(d msgpack.Decoder) (HealthCheckRequest, error) {

	var val HealthCheckRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Return value from actors and providers for health check status
type HealthCheckResponse struct {
	// A flag that indicates the the actor is healthy
	Healthy bool `json:"Healthy"`
	// A message containing additional information about the actors health
	Message string `json:"Message"`
}

func (o *HealthCheckResponse) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Healthy")
	encoder.WriteBool(o.Healthy)
	encoder.WriteString("Message")
	encoder.WriteString(o.Message)

	return nil
}
func DecodeHealthCheckResponse(d msgpack.Decoder) (HealthCheckResponse, error) {

	var val HealthCheckResponse
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		case "Healthy":
			val.Healthy, err = d.ReadBool()
		case "Message":
			val.Message, err = d.ReadString()
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// initialization data for a capability provider
type HostData struct {
	HostId             string        `json:"host_id"`
	LatticeRpcPrefix   string        `json:"lattice_rpc_prefix"`
	LinkName           string        `json:"link_name"`
	LatticeRpcUserJwt  string        `json:"lattice_rpc_user_jwt"`
	LatticeRpcUserSeed string        `json:"lattice_rpc_user_seed"`
	LatticeRpcUrl      string        `json:"lattice_rpc_url"`
	ProviderKey        string        `json:"provider_key"`
	InvocationSeed     string        `json:"invocation_seed"`
	EnvValues          HostEnvValues `json:"env_values"`
	InstanceId         string        `json:"instance_id"`
	// initial list of links for provider
	LinkDefinitions ActorLinks `json:"link_definitions"`
	// list of cluster issuers
	ClusterIssuers ClusterIssuers `json:"cluster_issuers"`
	// Optional configuration JSON sent to a given link name of a provider
	// without an actor context
	ConfigJson string `json:"config_json"`
	// Host-wide default RPC timeout for rpc messages, in milliseconds.  Defaults to 2000.
	DefaultRpcTimeoutMs uint64 `json:"default_rpc_timeout_ms"`
}

func (o *HostData) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(14)
	encoder.WriteString("host_id")
	encoder.WriteString(o.HostId)
	encoder.WriteString("lattice_rpc_prefix")
	encoder.WriteString(o.LatticeRpcPrefix)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("lattice_rpc_user_jwt")
	encoder.WriteString(o.LatticeRpcUserJwt)
	encoder.WriteString("lattice_rpc_user_seed")
	encoder.WriteString(o.LatticeRpcUserSeed)
	encoder.WriteString("lattice_rpc_url")
	encoder.WriteString(o.LatticeRpcUrl)
	encoder.WriteString("provider_key")
	encoder.WriteString(o.ProviderKey)
	encoder.WriteString("invocation_seed")
	encoder.WriteString(o.InvocationSeed)
	encoder.WriteString("env_values")
	o.EnvValues.Encode(encoder)
	encoder.WriteString("instance_id")
	encoder.WriteString(o.InstanceId)
	encoder.WriteString("link_definitions")
	o.LinkDefinitions.Encode(encoder)
	encoder.WriteString("cluster_issuers")
	o.ClusterIssuers.Encode(encoder)
	encoder.WriteString("config_json")
	encoder.WriteString(o.ConfigJson)
	encoder.WriteString("default_rpc_timeout_ms")
	encoder.WriteUint64(o.DefaultRpcTimeoutMs)

	return nil
}
func DecodeHostData(d msgpack.Decoder) (HostData, error) {

	var val HostData
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		case "host_id":
			val.HostId, err = d.ReadString()
		case "lattice_rpc_prefix":
			val.LatticeRpcPrefix, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "lattice_rpc_user_jwt":
			val.LatticeRpcUserJwt, err = d.ReadString()
		case "lattice_rpc_user_seed":
			val.LatticeRpcUserSeed, err = d.ReadString()
		case "lattice_rpc_url":
			val.LatticeRpcUrl, err = d.ReadString()
		case "provider_key":
			val.ProviderKey, err = d.ReadString()
		case "invocation_seed":
			val.InvocationSeed, err = d.ReadString()
		case "env_values":
			val.EnvValues, err = DecodeHostEnvValues(d)
		case "instance_id":
			val.InstanceId, err = d.ReadString()
		case "link_definitions":
			val.LinkDefinitions, err = DecodeActorLinks(d)
		case "cluster_issuers":
			val.ClusterIssuers, err = DecodeClusterIssuers(d)
		case "config_json":
			val.ConfigJson, err = d.ReadString()
		case "default_rpc_timeout_ms":
			val.DefaultRpcTimeoutMs, err = d.ReadUint64()
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Environment settings for initializing a capability provider
type HostEnvValues map[string]string

func (o *HostEnvValues) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return nil
}
func DecodeHostEnvValues(d msgpack.Decoder) (HostEnvValues, error) {

	isNil, err := d.IsNextNil()
	if err != nil && isNil {
		d.Skip()
		return make(map[string]string, 0), nil
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil

}

// RPC message to capability provider
type Invocation struct {
	Origin        WasmCloudEntity `json:"Origin"`
	Target        WasmCloudEntity `json:"Target"`
	Operation     string          `json:"Operation"`
	Msg           []byte          `json:"Msg"`
	Id            string          `json:"Id"`
	EncodedClaims string          `json:"encoded_claims"`
	HostId        string          `json:"host_id"`
	// total message size (optional)
	ContentLength uint64 `json:"content_length"`
	// Open Telemetry tracing support
	TraceContext TraceContext `json:"TraceContext"`
}

func (o *Invocation) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(9)
	encoder.WriteString("Origin")
	o.Origin.Encode(encoder)
	encoder.WriteString("Target")
	o.Target.Encode(encoder)
	encoder.WriteString("Operation")
	encoder.WriteString(o.Operation)
	encoder.WriteString("Msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("Id")
	encoder.WriteString(o.Id)
	encoder.WriteString("encoded_claims")
	encoder.WriteString(o.EncodedClaims)
	encoder.WriteString("host_id")
	encoder.WriteString(o.HostId)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("TraceContext")
	if o.TraceContext == nil {
		encoder.WriteNil()
	} else {
		o.TraceContext.Encode(encoder)
	}

	return nil
}
func DecodeInvocation(d msgpack.Decoder) (Invocation, error) {

	var val Invocation
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		case "Origin":
			val.Origin, err = DecodeWasmCloudEntity(d)
		case "Target":
			val.Target, err = DecodeWasmCloudEntity(d)
		case "Operation":
			val.Operation, err = d.ReadString()
		case "Msg":
			val.Msg, err = d.ReadByteArray()
		case "Id":
			val.Id, err = d.ReadString()
		case "encoded_claims":
			val.EncodedClaims, err = d.ReadString()
		case "host_id":
			val.HostId, err = d.ReadString()
		case "content_length":
			val.ContentLength, err = d.ReadUint64()
		case "TraceContext":
			val.TraceContext, err = DecodeTraceContext(d)
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Response to an invocation
type InvocationResponse struct {
	// serialize response message
	Msg []byte `json:"Msg"`
	// id connecting this response to the invocation
	InvocationId string `json:"invocation_id"`
	// optional error message
	Error string `json:"Error"`
	// total message size (optional)
	ContentLength uint64 `json:"content_length"`
}

func (o *InvocationResponse) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("Msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("invocation_id")
	encoder.WriteString(o.InvocationId)
	encoder.WriteString("Error")
	encoder.WriteString(o.Error)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)

	return nil
}
func DecodeInvocationResponse(d msgpack.Decoder) (InvocationResponse, error) {

	var val InvocationResponse
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		case "Msg":
			val.Msg, err = d.ReadByteArray()
		case "invocation_id":
			val.InvocationId, err = d.ReadString()
		case "Error":
			val.Error, err = d.ReadString()
		case "content_length":
			val.ContentLength, err = d.ReadUint64()
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Link definition for binding actor to provider
type LinkDefinition struct {
	// actor public key
	ActorId string `json:"actor_id"`
	// provider public key
	ProviderId string `json:"provider_id"`
	// link name
	LinkName string `json:"link_name"`
	// contract id
	ContractId string       `json:"contract_id"`
	Values     LinkSettings `json:"Values"`
}

func (o *LinkDefinition) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("actor_id")
	encoder.WriteString(o.ActorId)
	encoder.WriteString("provider_id")
	encoder.WriteString(o.ProviderId)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	encoder.WriteString(o.ContractId)
	encoder.WriteString("Values")
	o.Values.Encode(encoder)

	return nil
}
func DecodeLinkDefinition(d msgpack.Decoder) (LinkDefinition, error) {

	var val LinkDefinition
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		case "actor_id":
			val.ActorId, err = d.ReadString()
		case "provider_id":
			val.ProviderId, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = d.ReadString()
		case "Values":
			val.Values, err = DecodeLinkSettings(d)
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Settings associated with an actor-provider link
type LinkSettings map[string]string

func (o *LinkSettings) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return nil
}
func DecodeLinkSettings(d msgpack.Decoder) (LinkSettings, error) {

	isNil, err := d.IsNextNil()
	if err != nil && isNil {
		d.Skip()
		return make(map[string]string, 0), nil
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil

}

// Environment settings for initializing a capability provider
type TraceContext map[string]string

func (o *TraceContext) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return nil
}
func DecodeTraceContext(d msgpack.Decoder) (TraceContext, error) {

	isNil, err := d.IsNextNil()
	if err != nil && isNil {
		d.Skip()
		return make(map[string]string, 0), nil
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil

}

type WasmCloudEntity struct {
	PublicKey  string               `json:"public_key"`
	LinkName   string               `json:"link_name"`
	ContractId CapabilityContractId `json:"contract_id"`
}

func (o *WasmCloudEntity) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("public_key")
	encoder.WriteString(o.PublicKey)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	o.ContractId.Encode(encoder)

	return nil
}
func DecodeWasmCloudEntity(d msgpack.Decoder) (WasmCloudEntity, error) {

	var val WasmCloudEntity
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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

		case "public_key":
			val.PublicKey, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = DecodeCapabilityContractId(d)
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Actor service
type Actor interface {
	// Perform health check. Called at regular intervals by host
	HealthRequest(ctx *Context, arg HealthCheckRequest) (*HealthCheckResponse, error)
}

// ActorReceiver receives messages defined in the Actor service interface
// Actor service
type ActorReceiver struct{}

func (r *ActorReceiver) dispatch(ctx *Context, svc Actor, message *Message) (*Message, error) {
	switch message.Method {
	case "HealthRequest":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeHealthCheckRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.HealthRequest(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.Encode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.Encode(enc)

			return &Message{Method: "Actor.HealthRequest", Arg: buf}, nil
		}
	default:
		return nil, NewRpcError("MethodNotHandled", "Actor."+message.Method)
	}
}

// ActorSender sends messages to a Actor service
// Actor service
type ActorSender struct{ transport Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorActorSender(actor_id string) *ActorSender {
	transport := ToActor(actor_id)
	return &ActorSender{transport: transport}
}

// Perform health check. Called at regular intervals by host
func (s *ActorSender) HealthRequest(ctx *Context, arg HealthCheckRequest) (*HealthCheckResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, Message{Method: "Actor.HealthRequest", Arg: buf})

	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeHealthCheckResponse(d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
