// wasmcloud core data models for messaging and code generation
package actor

import (
    "github.com/wapc/tinygo-msgpack" //nolint
)

// Capability contract id, e.g. 'wasmcloud:httpserver'
type CapabilityContractId string

func (o *CapabilityContractId) Encode(encoder msgpack.Writer) error {
    encoder.WriteString(string(*o))
    return nil
}

func DecodeCapabilityContractId(d msgpack.Decoder) (CapabilityContractId, error) {

    val, err := d.ReadString()
    if err != nil {
        return "", err
    }
    return CapabilityContractId(val), nil

}

// 32-bit float
type F32 float32

// 64-bit float aka double
type F64 float64

// signed 16-bit int
type I16 int16

// signed 32-bit int
type I32 int32

// signed 64-bit int
type I64 int64

// signed byte
type I8 int8

// list of identifiers
type IdentifierList []string

func (o *IdentifierList) Encode(encoder msgpack.Writer) error {

    encoder.WriteArraySize(uint32(len(*o)))
    for _, item_o := range *o {
        encoder.WriteString(item_o)
    }

    return nil
}
func DecodeIdentifierList(d msgpack.Decoder) (IdentifierList, error) {

    isNil, err := d.IsNextNil()
    if err == nil && isNil {
        d.Skip()
        return make([]string, 0), nil
    }
    size, err := d.ReadArraySize()
    if err != nil {
        size = 0
    }
    val := make([]string, size)
    for i := uint32(0); i < size; i++ {
        item, err := d.ReadString()
        if err != nil {
            return val, err
        }
        val = append(val, item)
    }
    return val, nil

}

// unsigned 16-bit int
type U16 int16

// unsigned 32-bit int
type U32 int32

// unsigned 64-bit int
type U64 int64

// unsigned byte
type U8 int8

// Unit type
type Unit struct {
}

func (o *Unit) Encode(encoder msgpack.Writer) error {
    encoder.WriteNil()
    return nil
}
func DecodeUnit(d msgpack.Decoder) (Unit, error) {
    _ = d.Skip()
    return Unit{}, nil
}

// Rust codegen traits
type CodegenRust struct {
    // if true, disables deriving 'Default' trait
    NoDeriveDefault bool `json:"NoDeriveDefault"`
    // if true, disables deriving 'Eq' trait
    NoDeriveEq bool `json:"NoDeriveEq"`
    // adds `[#non_exhaustive]` attribute to a struct declaration
    NonExhaustive bool `json:"NonExhaustive"`
    // if true, do not generate code for this item.
    // This trait can be used if an item needs to be hand-generated
    Skip bool `json:"Skip"`
}

// indicates that a trait or class extends one or more bases
type Extends struct {
    Base IdentifierList `json:"Base"`
}

// Field sequence number. A zero-based field number for each member of a structure,
// to enable deterministic cbor serialization and improve forward and backward compatibility.
// Although the values are not required to be sequential, gaps are filled with nulls
// during encoding and so will slightly increase the encoding size.
type N int16

// A non-empty string (minimum length 1)
type NonEmptyString string

// Rename item(s) in target language.
// Useful if the item name (operation, or field) conflicts with a keyword in the target language.
// example: @rename({lang:"python",name:"delete"})
type Rename []RenameItem

// list element of trait @rename. the item name in the target language
// see '@rename'
type RenameItem struct {
    // language
    Lang string `json:"Lang"`
    // the name of the structure/operation/field
    Name string `json:"Name"`
}

// Overrides for serializer & deserializer
type Serialization struct {
    // (optional setting) Override field name when serializing and deserializing
    // By default, (when `name` not specified) is the exact declared name without
    // casing transformations. This setting does not affect the field name
    // produced in code generation, which is always lanaguage-idiomatic
    Name string `json:"Name"`
}

// This trait doesn't have any functional impact on codegen. It is simply
// to document that the defined type is a synonym, and to silence
// the default validator that prints a notice for synonyms with no traits.
type Synonym struct {
}

// The unsignedInt trait indicates that one of the number types is unsigned
type UnsignedInt struct {
}

// a protocol defines the semantics
// of how a client and server communicate.
type Wasmbus struct {
    // indicates this service's operations are handled by an actor (default false)
    ActorReceive bool `json:"ActorReceive"`
    // capability id such as "wasmcloud:httpserver"
    // always required for providerReceive, but optional for actorReceive
    ContractId CapabilityContractId `json:"ContractId"`
    // Binary message protocol version. Defaults to "0" if unset.
    // Be aware that changing this value can break binary compatibility unless
    // all users of this interface recompile
    Protocol string `json:"Protocol"`
    // indicates this service's operations are handled by an provider (default false)
    ProviderReceive bool `json:"ProviderReceive"`
}

// data sent via wasmbus
// This trait is required for all messages sent via wasmbus
type WasmbusData struct {
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
