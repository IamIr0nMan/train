// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: train.proto

package trainService

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To      string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User    *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Price   float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Section string  `protobuf:"bytes,5,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{1}
}

func (x *Ticket) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Ticket) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Ticket) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Ticket) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ticket) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

var File_train_proto protoreflect.FileDescriptor

var file_train_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x58, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xbf, 0x02, 0x0a,
	0x0c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x14, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x14,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x3c, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x0f,
	0x5a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_train_proto_rawDescOnce sync.Once
	file_train_proto_rawDescData = file_train_proto_rawDesc
)

func file_train_proto_rawDescGZIP() []byte {
	file_train_proto_rawDescOnce.Do(func() {
		file_train_proto_rawDescData = protoimpl.X.CompressGZIP(file_train_proto_rawDescData)
	})
	return file_train_proto_rawDescData
}

var file_train_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_train_proto_goTypes = []interface{}{
	(*User)(nil),   // 0: trainService.User
	(*Ticket)(nil), // 1: trainService.Ticket
}
var file_train_proto_depIdxs = []int32{
	0, // 0: trainService.Ticket.user:type_name -> trainService.User
	1, // 1: trainService.TrainService.PurchaseTicket:input_type -> trainService.Ticket
	0, // 2: trainService.TrainService.GetReceipt:input_type -> trainService.User
	1, // 3: trainService.TrainService.GetUsersBySection:input_type -> trainService.Ticket
	0, // 4: trainService.TrainService.CancelTicket:input_type -> trainService.User
	1, // 5: trainService.TrainService.ModifyUserSeat:input_type -> trainService.Ticket
	1, // 6: trainService.TrainService.PurchaseTicket:output_type -> trainService.Ticket
	1, // 7: trainService.TrainService.GetReceipt:output_type -> trainService.Ticket
	1, // 8: trainService.TrainService.GetUsersBySection:output_type -> trainService.Ticket
	1, // 9: trainService.TrainService.CancelTicket:output_type -> trainService.Ticket
	1, // 10: trainService.TrainService.ModifyUserSeat:output_type -> trainService.Ticket
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_train_proto_init() }
func file_train_proto_init() {
	if File_train_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_train_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_train_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
			RawDescriptor: file_train_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_train_proto_goTypes,
		DependencyIndexes: file_train_proto_depIdxs,
		MessageInfos:      file_train_proto_msgTypes,
	}.Build()
	File_train_proto = out.File
	file_train_proto_rawDesc = nil
	file_train_proto_goTypes = nil
	file_train_proto_depIdxs = nil
}