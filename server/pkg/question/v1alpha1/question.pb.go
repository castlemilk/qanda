// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: question/v1alpha1/question.proto

package questionsv1alpha1

import (
	v1alpha1 "github.com/castlemilk/qanda/server/pkg/api/v1alpha1"
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

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1alpha1.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Author   *v1alpha1.Author   `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Spec     *QuestionSpec      `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_question_proto_rawDescGZIP(), []int{0}
}

func (x *Question) GetMetadata() *v1alpha1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Question) GetAuthor() *v1alpha1.Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Question) GetSpec() *QuestionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type QuestionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string          `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Tags    []*v1alpha1.Tag `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *QuestionSpec) Reset() {
	*x = QuestionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionSpec) ProtoMessage() {}

func (x *QuestionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionSpec.ProtoReflect.Descriptor instead.
func (*QuestionSpec) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_question_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionSpec) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *QuestionSpec) GetTags() []*v1alpha1.Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_question_v1alpha1_question_proto protoreflect.FileDescriptor

var file_question_v1alpha1_question_proto_rawDesc = []byte{
	0x0a, 0x20, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x4f, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x6c, 0x65, 0x6d,
	0x69, 0x6c, 0x6b, 0x2f, 0x71, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_question_v1alpha1_question_proto_rawDescOnce sync.Once
	file_question_v1alpha1_question_proto_rawDescData = file_question_v1alpha1_question_proto_rawDesc
)

func file_question_v1alpha1_question_proto_rawDescGZIP() []byte {
	file_question_v1alpha1_question_proto_rawDescOnce.Do(func() {
		file_question_v1alpha1_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_v1alpha1_question_proto_rawDescData)
	})
	return file_question_v1alpha1_question_proto_rawDescData
}

var file_question_v1alpha1_question_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_question_v1alpha1_question_proto_goTypes = []interface{}{
	(*Question)(nil),          // 0: questions.v1alpha1.Question
	(*QuestionSpec)(nil),      // 1: questions.v1alpha1.QuestionSpec
	(*v1alpha1.Metadata)(nil), // 2: api.v1alpha1.Metadata
	(*v1alpha1.Author)(nil),   // 3: api.v1alpha1.Author
	(*v1alpha1.Tag)(nil),      // 4: api.v1alpha1.Tag
}
var file_question_v1alpha1_question_proto_depIdxs = []int32{
	2, // 0: questions.v1alpha1.Question.metadata:type_name -> api.v1alpha1.Metadata
	3, // 1: questions.v1alpha1.Question.author:type_name -> api.v1alpha1.Author
	1, // 2: questions.v1alpha1.Question.spec:type_name -> questions.v1alpha1.QuestionSpec
	4, // 3: questions.v1alpha1.QuestionSpec.tags:type_name -> api.v1alpha1.Tag
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_question_v1alpha1_question_proto_init() }
func file_question_v1alpha1_question_proto_init() {
	if File_question_v1alpha1_question_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_question_v1alpha1_question_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
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
		file_question_v1alpha1_question_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionSpec); i {
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
			RawDescriptor: file_question_v1alpha1_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_question_v1alpha1_question_proto_goTypes,
		DependencyIndexes: file_question_v1alpha1_question_proto_depIdxs,
		MessageInfos:      file_question_v1alpha1_question_proto_msgTypes,
	}.Build()
	File_question_v1alpha1_question_proto = out.File
	file_question_v1alpha1_question_proto_rawDesc = nil
	file_question_v1alpha1_question_proto_goTypes = nil
	file_question_v1alpha1_question_proto_depIdxs = nil
}
