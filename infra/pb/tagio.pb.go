// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: infra/pb/tagio.proto

package pb

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_pb_tagio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_infra_pb_tagio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_infra_pb_tagio_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TestOneof:
	//	*PushRequest_Request
	//	*PushRequest_Chunk
	TestOneof isPushRequest_TestOneof `protobuf_oneof:"test_oneof"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_pb_tagio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_pb_tagio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_infra_pb_tagio_proto_rawDescGZIP(), []int{1}
}

func (m *PushRequest) GetTestOneof() isPushRequest_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *PushRequest) GetRequest() *Request {
	if x, ok := x.GetTestOneof().(*PushRequest_Request); ok {
		return x.Request
	}
	return nil
}

func (x *PushRequest) GetChunk() *Chunk {
	if x, ok := x.GetTestOneof().(*PushRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isPushRequest_TestOneof interface {
	isPushRequest_TestOneof()
}

type PushRequest_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type PushRequest_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*PushRequest_Request) isPushRequest_TestOneof() {}

func (*PushRequest_Chunk) isPushRequest_TestOneof() {}

type PushResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushResult) Reset() {
	*x = PushResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_pb_tagio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResult) ProtoMessage() {}

func (x *PushResult) ProtoReflect() protoreflect.Message {
	mi := &file_infra_pb_tagio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResult.ProtoReflect.Descriptor instead.
func (*PushResult) Descriptor() ([]byte, []int) {
	return file_infra_pb_tagio_proto_rawDescGZIP(), []int{2}
}

type PullResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TestOneof:
	//	*PullResult_Progress
	//	*PullResult_Chunk
	TestOneof isPullResult_TestOneof `protobuf_oneof:"test_oneof"`
}

func (x *PullResult) Reset() {
	*x = PullResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_pb_tagio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullResult) ProtoMessage() {}

func (x *PullResult) ProtoReflect() protoreflect.Message {
	mi := &file_infra_pb_tagio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullResult.ProtoReflect.Descriptor instead.
func (*PullResult) Descriptor() ([]byte, []int) {
	return file_infra_pb_tagio_proto_rawDescGZIP(), []int{3}
}

func (m *PullResult) GetTestOneof() isPullResult_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *PullResult) GetProgress() *Progress {
	if x, ok := x.GetTestOneof().(*PullResult_Progress); ok {
		return x.Progress
	}
	return nil
}

func (x *PullResult) GetChunk() *Chunk {
	if x, ok := x.GetTestOneof().(*PullResult_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isPullResult_TestOneof interface {
	isPullResult_TestOneof()
}

type PullResult_Progress struct {
	Progress *Progress `protobuf:"bytes,1,opt,name=progress,proto3,oneof"`
}

type PullResult_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*PullResult_Progress) isPullResult_TestOneof() {}

func (*PullResult_Chunk) isPullResult_TestOneof() {}

type Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Offset      uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size        int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *Progress) Reset() {
	*x = Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_pb_tagio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_infra_pb_tagio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_infra_pb_tagio_proto_rawDescGZIP(), []int{4}
}

func (x *Progress) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Progress) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Progress) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_pb_tagio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_infra_pb_tagio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_infra_pb_tagio_proto_rawDescGZIP(), []int{5}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_infra_pb_tagio_proto protoreflect.FileDescriptor

var file_infra_pb_tagio_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x67, 0x69, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x51, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67, 0x0a,
	0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x0c, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x69, 0x0a, 0x0a, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x42, 0x0c, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22,
	0x58, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x60, 0x0a, 0x0c,
	0x54, 0x61, 0x67, 0x49, 0x4f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x50, 0x75, 0x6c, 0x6c, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x30, 0x01, 0x12, 0x29, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x28, 0x01, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x63,
	0x61, 0x72, 0x64, 0x6f, 0x6d, 0x61, 0x72, 0x61, 0x73, 0x63, 0x68, 0x69, 0x6e, 0x69, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x74, 0x61, 0x67, 0x73, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_pb_tagio_proto_rawDescOnce sync.Once
	file_infra_pb_tagio_proto_rawDescData = file_infra_pb_tagio_proto_rawDesc
)

func file_infra_pb_tagio_proto_rawDescGZIP() []byte {
	file_infra_pb_tagio_proto_rawDescOnce.Do(func() {
		file_infra_pb_tagio_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_pb_tagio_proto_rawDescData)
	})
	return file_infra_pb_tagio_proto_rawDescData
}

var file_infra_pb_tagio_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_infra_pb_tagio_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: pb.Request
	(*PushRequest)(nil), // 1: pb.PushRequest
	(*PushResult)(nil),  // 2: pb.PushResult
	(*PullResult)(nil),  // 3: pb.PullResult
	(*Progress)(nil),    // 4: pb.Progress
	(*Chunk)(nil),       // 5: pb.Chunk
}
var file_infra_pb_tagio_proto_depIdxs = []int32{
	0, // 0: pb.PushRequest.request:type_name -> pb.Request
	5, // 1: pb.PushRequest.chunk:type_name -> pb.Chunk
	4, // 2: pb.PullResult.progress:type_name -> pb.Progress
	5, // 3: pb.PullResult.chunk:type_name -> pb.Chunk
	0, // 4: pb.TagIOService.Pull:input_type -> pb.Request
	1, // 5: pb.TagIOService.Push:input_type -> pb.PushRequest
	3, // 6: pb.TagIOService.Pull:output_type -> pb.PullResult
	2, // 7: pb.TagIOService.Push:output_type -> pb.PushResult
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infra_pb_tagio_proto_init() }
func file_infra_pb_tagio_proto_init() {
	if File_infra_pb_tagio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_pb_tagio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_infra_pb_tagio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_infra_pb_tagio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResult); i {
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
		file_infra_pb_tagio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullResult); i {
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
		file_infra_pb_tagio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress); i {
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
		file_infra_pb_tagio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
	file_infra_pb_tagio_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PushRequest_Request)(nil),
		(*PushRequest_Chunk)(nil),
	}
	file_infra_pb_tagio_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*PullResult_Progress)(nil),
		(*PullResult_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_pb_tagio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_pb_tagio_proto_goTypes,
		DependencyIndexes: file_infra_pb_tagio_proto_depIdxs,
		MessageInfos:      file_infra_pb_tagio_proto_msgTypes,
	}.Build()
	File_infra_pb_tagio_proto = out.File
	file_infra_pb_tagio_proto_rawDesc = nil
	file_infra_pb_tagio_proto_goTypes = nil
	file_infra_pb_tagio_proto_depIdxs = nil
}
