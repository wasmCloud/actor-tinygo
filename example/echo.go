package main

import (
	msgpack "github.com/wapc/tinygo-msgpack"
	actor "github.com/wasmcloud/actor-tinygo"
)

// Once we have code generation support, we'll be able to
// write code that looks like this:
//
// func main() {
//   httpserver.RegisterHandleRequest(handleRequest)
// }
//
// func handleRequest(request httpserver.HttpRequest) (httpserver.HttpResponse, error) {
// ...
// }
//
// where `httpserver` is an auto-generated Go package containing the types and wrappers
// that you might find at https://github.com/wasmCloud/interfaces/httpserver/go

func main() {
	actor.RegisterHandlers(actor.Handlers{
		"HttpServer.HandleRequest": echo,
		"Actor.HealthRequest":      healthRequest,
	})
}

func echo(payload []byte) ([]byte, error) {
	r := Response{
		StatusCode: 200,
		Header:     map[string]string{},
		Body:       []byte("hello"),
	}
	bytes := r.ToBuffer()
	return bytes, nil
}

func healthRequest(payload []byte) ([]byte, error) {
	r := HealthCheckResponse{
		Healthy: true,
	}
	bytes := r.ToBuffer()
	return bytes, nil
}

type HealthCheckResponse struct {
	Healthy bool
	Message string
}

type Response struct {
	StatusCode uint32
	Status     string
	Header     map[string]string
	Body       []byte
}

func (o *Response) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Encode(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

func DecodeResponseNullable(decoder *msgpack.Decoder) (*Response, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeResponse(decoder)
	return &decoded, err
}

func DecodeResponse(decoder *msgpack.Decoder) (Response, error) {
	var o Response
	err := o.Decode(decoder)
	return o, err
}

func (o *Response) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "statusCode":
			o.StatusCode, err = decoder.ReadUint32()
		case "status":
			o.Status, err = decoder.ReadString()
		case "header":
			mapSize, err := decoder.ReadMapSize()
			if err != nil {
				return err
			}
			o.Header = make(map[string]string, mapSize)
			for mapSize > 0 {
				mapSize--
				key, err := decoder.ReadString()
				if err != nil {
					return err
				}
				value, err := decoder.ReadString()
				if err != nil {
					return err
				}
				o.Header[key] = value
			}
		case "body":
			o.Body, err = decoder.ReadByteArray()
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Response) Encode(encoder msgpack.Writer) error {
	if o == nil {
		encoder.WriteNil()
		return nil
	}
	encoder.WriteMapSize(4)
	encoder.WriteString("statusCode")
	encoder.WriteUint32(o.StatusCode)
	encoder.WriteString("status")
	encoder.WriteString(o.Status)
	encoder.WriteString("header")
	encoder.WriteMapSize(uint32(len(o.Header)))
	if o.Header != nil { // TinyGo bug: ranging over nil maps panics.
		for k, v := range o.Header {
			encoder.WriteString(k)
			encoder.WriteString(v)
		}
	}
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return nil
}

func DecodeHealthCheckResponseNullable(decoder *msgpack.Decoder) (*HealthCheckResponse, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeHealthCheckResponse(decoder)
	return &decoded, err
}

func DecodeHealthCheckResponse(decoder *msgpack.Decoder) (HealthCheckResponse, error) {
	var o HealthCheckResponse
	err := o.Decode(decoder)
	return o, err
}

func (o *HealthCheckResponse) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "healthy":
			o.Healthy, err = decoder.ReadBool()
		case "message":
			o.Message, err = decoder.ReadString()
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *HealthCheckResponse) Encode(encoder msgpack.Writer) error {
	if o == nil {
		encoder.WriteNil()
		return nil
	}
	encoder.WriteMapSize(2)
	encoder.WriteString("healthy")
	encoder.WriteBool(o.Healthy)
	encoder.WriteString("message")
	encoder.WriteString(o.Message)

	return nil
}

func (o *HealthCheckResponse) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Encode(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}
