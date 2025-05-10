package marshaler

import (
	"bytes"
	"encoding/gob"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MarshalerType string

const (
	ProtoMarshalerType MarshalerType = "proto"
	JsonMarshalerType  MarshalerType = "json"
	GobMarshalerType   MarshalerType = "gob"
)

type ProtoMarshaler interface {
	Marshal(m proto.Message) ([]byte, error)
	Unmarshal(b []byte, m proto.Message) error
}

func NewMarshaler(marshalType MarshalerType) ProtoMarshaler {
	switch marshalType {
	case ProtoMarshalerType:
		return &Proto{}
	case GobMarshalerType:
		return &Gob{}
	default:
		return &Json{}
	}
}

type Proto struct{}

// Marshal implements ProtoMarshaler.
func (*Proto) Marshal(m protoreflect.ProtoMessage) ([]byte, error) {
	return proto.Marshal(m)
}

// Unmarshal implements ProtoMarshaler.
func (*Proto) Unmarshal(b []byte, m protoreflect.ProtoMessage) error {
	return proto.Unmarshal(b, m)
}

type Json struct{}

// Marshal implements ProtoMarshaler.
func (*Json) Marshal(m protoreflect.ProtoMessage) ([]byte, error) {
	return protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}.Marshal(m)
}

// Unmarshal implements ProtoMarshaler.
func (*Json) Unmarshal(b []byte, m protoreflect.ProtoMessage) error {
	return protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}.Unmarshal(b, m)
}

type Gob struct{}

// Marshal implements ProtoMarshaler.
func (*Gob) Marshal(m protoreflect.ProtoMessage) ([]byte, error) {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(m)
	return buffer.Bytes(), err

}

// Unmarshal implements ProtoMarshaler.
func (*Gob) Unmarshal(b []byte, m protoreflect.ProtoMessage) error {
	buffer := bytes.NewBuffer(b)
	err := gob.NewDecoder(buffer).Decode(m)
	return err
}
