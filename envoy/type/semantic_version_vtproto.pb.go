//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/type/semantic_version.proto

package _type

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *SemanticVersion) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SemanticVersion) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *SemanticVersion) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Patch != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Patch))
		i--
		dAtA[i] = 0x18
	}
	if m.MinorNumber != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MinorNumber))
		i--
		dAtA[i] = 0x10
	}
	if m.MajorNumber != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.MajorNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SemanticVersion) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MajorNumber != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MajorNumber))
	}
	if m.MinorNumber != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.MinorNumber))
	}
	if m.Patch != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Patch))
	}
	n += len(m.unknownFields)
	return n
}
