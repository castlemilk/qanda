// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: question/v1alpha1/service.proto

package questionsv1alpha1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ------------ LIST ------------
// ListQuestionsRequest
type ListQuestionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQuestionsRequest) Reset() {
	*x = ListQuestionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestionsRequest) ProtoMessage() {}

func (x *ListQuestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestionsRequest.ProtoReflect.Descriptor instead.
func (*ListQuestionsRequest) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{0}
}

type ListQuestionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Questions []*Question `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *ListQuestionsResponse) Reset() {
	*x = ListQuestionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestionsResponse) ProtoMessage() {}

func (x *ListQuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestionsResponse.ProtoReflect.Descriptor instead.
func (*ListQuestionsResponse) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListQuestionsResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

// ------------ GET ------------
type GetQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetQuestionRequest) Reset() {
	*x = GetQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionRequest) ProtoMessage() {}

func (x *GetQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionRequest) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetQuestionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *GetQuestionResponse) Reset() {
	*x = GetQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionResponse) ProtoMessage() {}

func (x *GetQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionResponse) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetQuestionResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type CreateQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *CreateQuestionRequest) Reset() {
	*x = CreateQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionRequest) ProtoMessage() {}

func (x *CreateQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionRequest.ProtoReflect.Descriptor instead.
func (*CreateQuestionRequest) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateQuestionRequest) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type CreateQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *CreateQuestionResponse) Reset() {
	*x = CreateQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestionResponse) ProtoMessage() {}

func (x *CreateQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestionResponse.ProtoReflect.Descriptor instead.
func (*CreateQuestionResponse) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateQuestionResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type UpdateQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Question *Question `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *UpdateQuestionRequest) Reset() {
	*x = UpdateQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionRequest) ProtoMessage() {}

func (x *UpdateQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuestionRequest) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateQuestionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateQuestionRequest) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type UpdateQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *UpdateQuestionResponse) Reset() {
	*x = UpdateQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestionResponse) ProtoMessage() {}

func (x *UpdateQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestionResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuestionResponse) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateQuestionResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type DeleteQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQuestionRequest) Reset() {
	*x = DeleteQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestionRequest) ProtoMessage() {}

func (x *DeleteQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestionRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuestionRequest) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteQuestionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question *Question `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *DeleteQuestionResponse) Reset() {
	*x = DeleteQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_question_v1alpha1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestionResponse) ProtoMessage() {}

func (x *DeleteQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_question_v1alpha1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestionResponse.ProtoReflect.Descriptor instead.
func (*DeleteQuestionResponse) Descriptor() ([]byte, []int) {
	return file_question_v1alpha1_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteQuestionResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

var File_question_v1alpha1_service_proto protoreflect.FileDescriptor

var file_question_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x61, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x52, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0xac, 0x05, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x86,
	0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x89, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x42, 0xb3, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x6c, 0x65, 0x6d, 0x69, 0x6c, 0x6b, 0x2f, 0x71, 0x61,
	0x6e, 0x64, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x92, 0x41, 0x63, 0x12, 0x0a, 0x32, 0x08, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2a, 0x01, 0x02, 0x72, 0x52, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x20, 0x41, 0x50, 0x49, 0x12, 0x41, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x6c,
	0x65, 0x6d, 0x69, 0x6c, 0x6b, 0x2f, 0x71, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x71, 0x61, 0x6e, 0x64,
	0x61, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_question_v1alpha1_service_proto_rawDescOnce sync.Once
	file_question_v1alpha1_service_proto_rawDescData = file_question_v1alpha1_service_proto_rawDesc
)

func file_question_v1alpha1_service_proto_rawDescGZIP() []byte {
	file_question_v1alpha1_service_proto_rawDescOnce.Do(func() {
		file_question_v1alpha1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_question_v1alpha1_service_proto_rawDescData)
	})
	return file_question_v1alpha1_service_proto_rawDescData
}

var file_question_v1alpha1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_question_v1alpha1_service_proto_goTypes = []interface{}{
	(*ListQuestionsRequest)(nil),   // 0: questions.v1alpha1.ListQuestionsRequest
	(*ListQuestionsResponse)(nil),  // 1: questions.v1alpha1.ListQuestionsResponse
	(*GetQuestionRequest)(nil),     // 2: questions.v1alpha1.GetQuestionRequest
	(*GetQuestionResponse)(nil),    // 3: questions.v1alpha1.GetQuestionResponse
	(*CreateQuestionRequest)(nil),  // 4: questions.v1alpha1.CreateQuestionRequest
	(*CreateQuestionResponse)(nil), // 5: questions.v1alpha1.CreateQuestionResponse
	(*UpdateQuestionRequest)(nil),  // 6: questions.v1alpha1.UpdateQuestionRequest
	(*UpdateQuestionResponse)(nil), // 7: questions.v1alpha1.UpdateQuestionResponse
	(*DeleteQuestionRequest)(nil),  // 8: questions.v1alpha1.DeleteQuestionRequest
	(*DeleteQuestionResponse)(nil), // 9: questions.v1alpha1.DeleteQuestionResponse
	(*Question)(nil),               // 10: questions.v1alpha1.Question
}
var file_question_v1alpha1_service_proto_depIdxs = []int32{
	10, // 0: questions.v1alpha1.ListQuestionsResponse.questions:type_name -> questions.v1alpha1.Question
	10, // 1: questions.v1alpha1.GetQuestionResponse.question:type_name -> questions.v1alpha1.Question
	10, // 2: questions.v1alpha1.CreateQuestionRequest.question:type_name -> questions.v1alpha1.Question
	10, // 3: questions.v1alpha1.CreateQuestionResponse.question:type_name -> questions.v1alpha1.Question
	10, // 4: questions.v1alpha1.UpdateQuestionRequest.question:type_name -> questions.v1alpha1.Question
	10, // 5: questions.v1alpha1.UpdateQuestionResponse.question:type_name -> questions.v1alpha1.Question
	10, // 6: questions.v1alpha1.DeleteQuestionResponse.question:type_name -> questions.v1alpha1.Question
	0,  // 7: questions.v1alpha1.QuestionService.ListQuestions:input_type -> questions.v1alpha1.ListQuestionsRequest
	2,  // 8: questions.v1alpha1.QuestionService.GetQuestion:input_type -> questions.v1alpha1.GetQuestionRequest
	4,  // 9: questions.v1alpha1.QuestionService.CreateQuestion:input_type -> questions.v1alpha1.CreateQuestionRequest
	8,  // 10: questions.v1alpha1.QuestionService.DeleteQuestion:input_type -> questions.v1alpha1.DeleteQuestionRequest
	6,  // 11: questions.v1alpha1.QuestionService.UpdateQuestion:input_type -> questions.v1alpha1.UpdateQuestionRequest
	1,  // 12: questions.v1alpha1.QuestionService.ListQuestions:output_type -> questions.v1alpha1.ListQuestionsResponse
	3,  // 13: questions.v1alpha1.QuestionService.GetQuestion:output_type -> questions.v1alpha1.GetQuestionResponse
	5,  // 14: questions.v1alpha1.QuestionService.CreateQuestion:output_type -> questions.v1alpha1.CreateQuestionResponse
	9,  // 15: questions.v1alpha1.QuestionService.DeleteQuestion:output_type -> questions.v1alpha1.DeleteQuestionResponse
	7,  // 16: questions.v1alpha1.QuestionService.UpdateQuestion:output_type -> questions.v1alpha1.UpdateQuestionResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_question_v1alpha1_service_proto_init() }
func file_question_v1alpha1_service_proto_init() {
	if File_question_v1alpha1_service_proto != nil {
		return
	}
	file_question_v1alpha1_question_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_question_v1alpha1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuestionsRequest); i {
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
		file_question_v1alpha1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuestionsResponse); i {
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
		file_question_v1alpha1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuestionRequest); i {
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
		file_question_v1alpha1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuestionResponse); i {
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
		file_question_v1alpha1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionRequest); i {
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
		file_question_v1alpha1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestionResponse); i {
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
		file_question_v1alpha1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuestionRequest); i {
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
		file_question_v1alpha1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuestionResponse); i {
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
		file_question_v1alpha1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuestionRequest); i {
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
		file_question_v1alpha1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuestionResponse); i {
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
			RawDescriptor: file_question_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_question_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_question_v1alpha1_service_proto_depIdxs,
		MessageInfos:      file_question_v1alpha1_service_proto_msgTypes,
	}.Build()
	File_question_v1alpha1_service_proto = out.File
	file_question_v1alpha1_service_proto_rawDesc = nil
	file_question_v1alpha1_service_proto_goTypes = nil
	file_question_v1alpha1_service_proto_depIdxs = nil
}
