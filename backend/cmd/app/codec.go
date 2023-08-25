package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

type JSONCodec struct {
	runtime.JSONPb
}

func (codec JSONCodec) Name() string {
	return "json"
}

func NewJSONCodec() *JSONCodec {
	return &JSONCodec{
		runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
		},
	}
}
