// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: market_data_provider/market_data_provider.proto

package market_data_provider

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *GetTickRequest) Reset() {
	*x = GetTickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_data_provider_market_data_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTickRequest) ProtoMessage() {}

func (x *GetTickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_market_data_provider_market_data_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTickRequest.ProtoReflect.Descriptor instead.
func (*GetTickRequest) Descriptor() ([]byte, []int) {
	return file_market_data_provider_market_data_provider_proto_rawDescGZIP(), []int{0}
}

func (x *GetTickRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type GetTickResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol         string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	InstrumentType string  `protobuf:"bytes,2,opt,name=instrumentType,proto3" json:"instrumentType,omitempty"`
	Currency       string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Timestamp      int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Open           float64 `protobuf:"fixed64,5,opt,name=open,proto3" json:"open,omitempty"`
	High           float64 `protobuf:"fixed64,6,opt,name=high,proto3" json:"high,omitempty"`
	Low            float64 `protobuf:"fixed64,7,opt,name=low,proto3" json:"low,omitempty"`
	Close          float64 `protobuf:"fixed64,8,opt,name=close,proto3" json:"close,omitempty"`
	Volume         int64   `protobuf:"varint,9,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *GetTickResponse) Reset() {
	*x = GetTickResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_market_data_provider_market_data_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTickResponse) ProtoMessage() {}

func (x *GetTickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_market_data_provider_market_data_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTickResponse.ProtoReflect.Descriptor instead.
func (*GetTickResponse) Descriptor() ([]byte, []int) {
	return file_market_data_provider_market_data_provider_proto_rawDescGZIP(), []int{1}
}

func (x *GetTickResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetTickResponse) GetInstrumentType() string {
	if x != nil {
		return x.InstrumentType
	}
	return ""
}

func (x *GetTickResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetTickResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetTickResponse) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *GetTickResponse) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *GetTickResponse) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *GetTickResponse) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *GetTickResponse) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

var File_market_data_provider_market_data_provider_proto protoreflect.FileDescriptor

var file_market_data_provider_market_data_provider_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x70, 0x6f, 0x67, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0xf3, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x75, 0x0a, 0x11,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x60, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x28, 0x2e, 0x70,
	0x6f, 0x67, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x6f, 0x67, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0xf5, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x67, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x42, 0x17, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x73, 0x74,
	0x65, 0x69, 0x72, 0x6f, 0x31, 0x31, 0x2f, 0x50, 0x6f, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x50, 0x4d, 0x58, 0xaa, 0x02, 0x16, 0x50, 0x6f, 0x67,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0xca, 0x02, 0x16, 0x50, 0x6f, 0x67, 0x5c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0xe2, 0x02, 0x22, 0x50,
	0x6f, 0x67, 0x5c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x17, 0x50, 0x6f, 0x67, 0x3a, 0x3a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_market_data_provider_market_data_provider_proto_rawDescOnce sync.Once
	file_market_data_provider_market_data_provider_proto_rawDescData = file_market_data_provider_market_data_provider_proto_rawDesc
)

func file_market_data_provider_market_data_provider_proto_rawDescGZIP() []byte {
	file_market_data_provider_market_data_provider_proto_rawDescOnce.Do(func() {
		file_market_data_provider_market_data_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_market_data_provider_market_data_provider_proto_rawDescData)
	})
	return file_market_data_provider_market_data_provider_proto_rawDescData
}

var file_market_data_provider_market_data_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_market_data_provider_market_data_provider_proto_goTypes = []interface{}{
	(*GetTickRequest)(nil),  // 0: pog.market_data_provider.GetTickRequest
	(*GetTickResponse)(nil), // 1: pog.market_data_provider.GetTickResponse
}
var file_market_data_provider_market_data_provider_proto_depIdxs = []int32{
	0, // 0: pog.market_data_provider.MarketDataService.GetTick:input_type -> pog.market_data_provider.GetTickRequest
	1, // 1: pog.market_data_provider.MarketDataService.GetTick:output_type -> pog.market_data_provider.GetTickResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_market_data_provider_market_data_provider_proto_init() }
func file_market_data_provider_market_data_provider_proto_init() {
	if File_market_data_provider_market_data_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_market_data_provider_market_data_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTickRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_market_data_provider_market_data_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTickResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_market_data_provider_market_data_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_market_data_provider_market_data_provider_proto_goTypes,
		DependencyIndexes: file_market_data_provider_market_data_provider_proto_depIdxs,
		MessageInfos:      file_market_data_provider_market_data_provider_proto_msgTypes,
	}.Build()
	File_market_data_provider_market_data_provider_proto = out.File
	file_market_data_provider_market_data_provider_proto_rawDesc = nil
	file_market_data_provider_market_data_provider_proto_goTypes = nil
	file_market_data_provider_market_data_provider_proto_depIdxs = nil
}
