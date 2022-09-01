// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: servers/articles/protobuf_files/articles/articles.proto

package articles

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

	AfterId string `protobuf:"bytes,1,opt,name=after_id,json=afterId,proto3" json:"after_id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Id      string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Page    int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Order   string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	Sort    string `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort,omitempty"`
	Title   string `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	UserId  string `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[0]
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
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAfterId() string {
	if x != nil {
		return x.AfterId
	}
	return ""
}

func (x *Request) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Request) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Request) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *Request) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Request) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	TotalPage int32 `protobuf:"varint,2,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Meta) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Data_Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Data_Users) Reset() {
	*x = Data_Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Users) ProtoMessage() {}

func (x *Data_Users) ProtoReflect() protoreflect.Message {
	mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Users.ProtoReflect.Descriptor instead.
func (*Data_Users) Descriptor() ([]byte, []int) {
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP(), []int{3}
}

func (x *Data_Users) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Data_Users) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Data_Users) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type Data_Articles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string      `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string      `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	User      *Data_Users `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt string      `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt string      `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Data_Articles) Reset() {
	*x = Data_Articles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Articles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Articles) ProtoMessage() {}

func (x *Data_Articles) ProtoReflect() protoreflect.Message {
	mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Articles.ProtoReflect.Descriptor instead.
func (*Data_Articles) Descriptor() ([]byte, []int) {
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP(), []int{4}
}

func (x *Data_Articles) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Data_Articles) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Data_Articles) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Data_Articles) GetUser() *Data_Users {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Data_Articles) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Data_Articles) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Data_Articles) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *Meta            `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Articles []*Data_Articles `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty"`
	Message  *Message         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Response) GetArticles() []*Data_Articles {
	if x != nil {
		return x.Articles
	}
	return nil
}

func (x *Response) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_servers_articles_protobuf_files_articles_articles_proto protoreflect.FileDescriptor

var file_servers_articles_protobuf_files_articles_articles_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x66, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x3b, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x22, 0x23,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x08,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc9,
	0x01, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servers_articles_protobuf_files_articles_articles_proto_rawDescOnce sync.Once
	file_servers_articles_protobuf_files_articles_articles_proto_rawDescData = file_servers_articles_protobuf_files_articles_articles_proto_rawDesc
)

func file_servers_articles_protobuf_files_articles_articles_proto_rawDescGZIP() []byte {
	file_servers_articles_protobuf_files_articles_articles_proto_rawDescOnce.Do(func() {
		file_servers_articles_protobuf_files_articles_articles_proto_rawDescData = protoimpl.X.CompressGZIP(file_servers_articles_protobuf_files_articles_articles_proto_rawDescData)
	})
	return file_servers_articles_protobuf_files_articles_articles_proto_rawDescData
}

var file_servers_articles_protobuf_files_articles_articles_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_servers_articles_protobuf_files_articles_articles_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: articles.Request
	(*Meta)(nil),          // 1: articles.Meta
	(*Message)(nil),       // 2: articles.Message
	(*Data_Users)(nil),    // 3: articles.Data_Users
	(*Data_Articles)(nil), // 4: articles.Data_Articles
	(*Response)(nil),      // 5: articles.Response
}
var file_servers_articles_protobuf_files_articles_articles_proto_depIdxs = []int32{
	3, // 0: articles.Data_Articles.user:type_name -> articles.Data_Users
	1, // 1: articles.Response.meta:type_name -> articles.Meta
	4, // 2: articles.Response.articles:type_name -> articles.Data_Articles
	2, // 3: articles.Response.message:type_name -> articles.Message
	0, // 4: articles.Users.Create:input_type -> articles.Request
	0, // 5: articles.Users.Read:input_type -> articles.Request
	0, // 6: articles.Users.Update:input_type -> articles.Request
	0, // 7: articles.Users.Delete:input_type -> articles.Request
	5, // 8: articles.Users.Create:output_type -> articles.Response
	5, // 9: articles.Users.Read:output_type -> articles.Response
	5, // 10: articles.Users.Update:output_type -> articles.Response
	5, // 11: articles.Users.Delete:output_type -> articles.Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_servers_articles_protobuf_files_articles_articles_proto_init() }
func file_servers_articles_protobuf_files_articles_articles_proto_init() {
	if File_servers_articles_protobuf_files_articles_articles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Users); i {
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
		file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Articles); i {
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
		file_servers_articles_protobuf_files_articles_articles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_servers_articles_protobuf_files_articles_articles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servers_articles_protobuf_files_articles_articles_proto_goTypes,
		DependencyIndexes: file_servers_articles_protobuf_files_articles_articles_proto_depIdxs,
		MessageInfos:      file_servers_articles_protobuf_files_articles_articles_proto_msgTypes,
	}.Build()
	File_servers_articles_protobuf_files_articles_articles_proto = out.File
	file_servers_articles_protobuf_files_articles_articles_proto_rawDesc = nil
	file_servers_articles_protobuf_files_articles_articles_proto_goTypes = nil
	file_servers_articles_protobuf_files_articles_articles_proto_depIdxs = nil
}
