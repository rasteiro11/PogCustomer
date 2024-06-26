// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: crypto/crypto.proto

package crypto

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

type ProcessDepositEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount      string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	EndToEnd    string `protobuf:"bytes,3,opt,name=end_to_end,json=endToEnd,proto3" json:"end_to_end,omitempty"`
	PaymentType string `protobuf:"bytes,4,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
}

func (x *ProcessDepositEventRequest) Reset() {
	*x = ProcessDepositEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDepositEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDepositEventRequest) ProtoMessage() {}

func (x *ProcessDepositEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDepositEventRequest.ProtoReflect.Descriptor instead.
func (*ProcessDepositEventRequest) Descriptor() ([]byte, []int) {
	return file_crypto_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessDepositEventRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ProcessDepositEventRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ProcessDepositEventRequest) GetEndToEnd() string {
	if x != nil {
		return x.EndToEnd
	}
	return ""
}

func (x *ProcessDepositEventRequest) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

type ProcessDepositEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProcessDepositEventResponse) Reset() {
	*x = ProcessDepositEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessDepositEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessDepositEventResponse) ProtoMessage() {}

func (x *ProcessDepositEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessDepositEventResponse.ProtoReflect.Descriptor instead.
func (*ProcessDepositEventResponse) Descriptor() ([]byte, []int) {
	return file_crypto_crypto_proto_rawDescGZIP(), []int{1}
}

type RegisterUserWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Wallet  string `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *RegisterUserWalletRequest) Reset() {
	*x = RegisterUserWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserWalletRequest) ProtoMessage() {}

func (x *RegisterUserWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserWalletRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserWalletRequest) Descriptor() ([]byte, []int) {
	return file_crypto_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterUserWalletRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterUserWalletRequest) GetWallet() string {
	if x != nil {
		return x.Wallet
	}
	return ""
}

func (x *RegisterUserWalletRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type RegisterUserWalletResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterUserWalletResponse) Reset() {
	*x = RegisterUserWalletResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_crypto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserWalletResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserWalletResponse) ProtoMessage() {}

func (x *RegisterUserWalletResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_crypto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserWalletResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserWalletResponse) Descriptor() ([]byte, []int) {
	return file_crypto_crypto_proto_rawDescGZIP(), []int{3}
}

var File_crypto_crypto_proto protoreflect.FileDescriptor

var file_crypto_crypto_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x67, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x45, 0x6e, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x66, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x1c, 0x0a, 0x1a, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x26, 0x2e, 0x70, 0x6f, 0x67, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6f, 0x67, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x6f, 0x67,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6f, 0x67, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9d, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x67, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x42, 0x0b,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x73, 0x74, 0x65, 0x69,
	0x72, 0x6f, 0x31, 0x31, 0x2f, 0x50, 0x6f, 0x67, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x50, 0x43, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x6f, 0x67,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0xca, 0x02, 0x0a, 0x50, 0x6f, 0x67, 0x5c, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0xe2, 0x02, 0x16, 0x50, 0x6f, 0x67, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b,
	0x50, 0x6f, 0x67, 0x3a, 0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_crypto_crypto_proto_rawDescOnce sync.Once
	file_crypto_crypto_proto_rawDescData = file_crypto_crypto_proto_rawDesc
)

func file_crypto_crypto_proto_rawDescGZIP() []byte {
	file_crypto_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_crypto_proto_rawDescData)
	})
	return file_crypto_crypto_proto_rawDescData
}

var file_crypto_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crypto_crypto_proto_goTypes = []interface{}{
	(*ProcessDepositEventRequest)(nil),  // 0: pog.crypto.ProcessDepositEventRequest
	(*ProcessDepositEventResponse)(nil), // 1: pog.crypto.ProcessDepositEventResponse
	(*RegisterUserWalletRequest)(nil),   // 2: pog.crypto.RegisterUserWalletRequest
	(*RegisterUserWalletResponse)(nil),  // 3: pog.crypto.RegisterUserWalletResponse
}
var file_crypto_crypto_proto_depIdxs = []int32{
	0, // 0: pog.crypto.CryptoService.ProcessDepositEvent:input_type -> pog.crypto.ProcessDepositEventRequest
	2, // 1: pog.crypto.CryptoService.RegisterUserWallet:input_type -> pog.crypto.RegisterUserWalletRequest
	1, // 2: pog.crypto.CryptoService.ProcessDepositEvent:output_type -> pog.crypto.ProcessDepositEventResponse
	3, // 3: pog.crypto.CryptoService.RegisterUserWallet:output_type -> pog.crypto.RegisterUserWalletResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypto_crypto_proto_init() }
func file_crypto_crypto_proto_init() {
	if File_crypto_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDepositEventRequest); i {
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
		file_crypto_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessDepositEventResponse); i {
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
		file_crypto_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserWalletRequest); i {
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
		file_crypto_crypto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserWalletResponse); i {
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
			RawDescriptor: file_crypto_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_crypto_proto_depIdxs,
		MessageInfos:      file_crypto_crypto_proto_msgTypes,
	}.Build()
	File_crypto_crypto_proto = out.File
	file_crypto_crypto_proto_rawDesc = nil
	file_crypto_crypto_proto_goTypes = nil
	file_crypto_crypto_proto_depIdxs = nil
}
