//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package transcoderv2

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

func (m *GrpcJsonTranscoder_PrintOptions) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *GrpcJsonTranscoder_PrintOptions) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.PreserveProtoFieldNames {
		i--
		if m.PreserveProtoFieldNames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.AlwaysPrintEnumsAsInts {
		i--
		if m.AlwaysPrintEnumsAsInts {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.AlwaysPrintPrimitiveFields {
		i--
		if m.AlwaysPrintPrimitiveFields {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.AddWhitespace {
		i--
		if m.AddWhitespace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GrpcJsonTranscoder) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *GrpcJsonTranscoder) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *GrpcJsonTranscoder) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.ConvertGrpcStatus {
		i--
		if m.ConvertGrpcStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.IgnoreUnknownQueryParameters {
		i--
		if m.IgnoreUnknownQueryParameters {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.AutoMapping {
		i--
		if m.AutoMapping {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for iNdEx := len(m.IgnoredQueryParameters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IgnoredQueryParameters[iNdEx])
			copy(dAtA[i:], m.IgnoredQueryParameters[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.IgnoredQueryParameters[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MatchIncomingRequestRoute {
		i--
		if m.MatchIncomingRequestRoute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if msg, ok := m.DescriptorSet.(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.PrintOptions != nil {
		size, err := m.PrintOptions.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Services[iNdEx])
			copy(dAtA[i:], m.Services[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.Services[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if msg, ok := m.DescriptorSet.(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtoDescriptor)
	copy(dAtA[i:], m.ProtoDescriptor)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ProtoDescriptor)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtoDescriptorBin)
	copy(dAtA[i:], m.ProtoDescriptorBin)
	i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.ProtoDescriptorBin)))
	i--
	dAtA[i] = 0x22
	return len(dAtA) - i, nil
}
func (m *GrpcJsonTranscoder_PrintOptions) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddWhitespace {
		n += 2
	}
	if m.AlwaysPrintPrimitiveFields {
		n += 2
	}
	if m.AlwaysPrintEnumsAsInts {
		n += 2
	}
	if m.PreserveProtoFieldNames {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *GrpcJsonTranscoder) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.DescriptorSet.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.PrintOptions != nil {
		l = m.PrintOptions.SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.MatchIncomingRequestRoute {
		n += 2
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for _, s := range m.IgnoredQueryParameters {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.AutoMapping {
		n += 2
	}
	if m.IgnoreUnknownQueryParameters {
		n += 2
	}
	if m.ConvertGrpcStatus {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtoDescriptor)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtoDescriptorBin)
	n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	return n
}
