// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: system/menu.proto

package system

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListMenusRequest is the request for the MenuService.ListMenus method.
type ListMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource id, for example, "shelves/shelf1".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The current page number.
	Current int32 `protobuf:"varint,2,opt,name=current,proto3" json:"current,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The no_paging is used to disable pagination.
	NoPaging bool `protobuf:"varint,5,opt,name=no_paging,json=noPaging,proto3" json:"no_paging,omitempty"`
	// The only_count is the query parameter for set only to query the total number
	OnlyCount bool `protobuf:"varint,6,opt,name=only_count,json=onlyCount,proto3" json:"only_count,omitempty"`
}

func (x *ListMenusRequest) Reset() {
	*x = ListMenusRequest{}
	mi := &file_system_menu_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenusRequest) ProtoMessage() {}

func (x *ListMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenusRequest.ProtoReflect.Descriptor instead.
func (*ListMenusRequest) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{0}
}

func (x *ListMenusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListMenusRequest) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ListMenusRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMenusRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListMenusRequest) GetNoPaging() bool {
	if x != nil {
		return x.NoPaging
	}
	return false
}

func (x *ListMenusRequest) GetOnlyCount() bool {
	if x != nil {
		return x.OnlyCount
	}
	return false
}

// ListMenusResponse is the response for the MenuService.ListMenus method.
type ListMenusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field id should match the noun "menus" in the method id.  There
	// will be a maximum number of items returned based on the page_size field
	// in the request.
	Menus []*Menu `protobuf:"bytes,1,rep,name=menus,proto3" json:"menus,omitempty"`
	// The total number of items in the list.
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	// The current page number.
	Current int32 `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,5,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information about this response.
	Extra *anypb.Any `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *ListMenusResponse) Reset() {
	*x = ListMenusResponse{}
	mi := &file_system_menu_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMenusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenusResponse) ProtoMessage() {}

func (x *ListMenusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenusResponse.ProtoReflect.Descriptor instead.
func (*ListMenusResponse) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{1}
}

func (x *ListMenusResponse) GetMenus() []*Menu {
	if x != nil {
		return x.Menus
	}
	return nil
}

func (x *ListMenusResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListMenusResponse) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ListMenusResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMenusResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListMenusResponse) GetExtra() *anypb.Any {
	if x != nil {
		return x.Extra
	}
	return nil
}

// GetMenuRequest is the request for the MenuService.GetMenu method.
type GetMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field will contain id of the resource requested, for example:
	// "shelves/shelf1/menus/menu2"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMenuRequest) Reset() {
	*x = GetMenuRequest{}
	mi := &file_system_menu_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuRequest) ProtoMessage() {}

func (x *GetMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuRequest.ProtoReflect.Descriptor instead.
func (*GetMenuRequest) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{2}
}

func (x *GetMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetMenuResponse is the response for the MenuService.GetMenu method.
type GetMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field id should match the Noun in the method id.
	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *GetMenuResponse) Reset() {
	*x = GetMenuResponse{}
	mi := &file_system_menu_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuResponse) ProtoMessage() {}

func (x *GetMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuResponse.ProtoReflect.Descriptor instead.
func (*GetMenuResponse) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{3}
}

func (x *GetMenuResponse) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

// CreateMenuRequest is the request for the MenuService.CreateMenu method.
type CreateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource id where the menu is to be created.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The menu id to use for this menu.
	MenuId string `protobuf:"bytes,3,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	// The menu resource to create.
	// The field id should match the Noun in the method id.
	Menu *Menu `protobuf:"bytes,2,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *CreateMenuRequest) Reset() {
	*x = CreateMenuRequest{}
	mi := &file_system_menu_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuRequest) ProtoMessage() {}

func (x *CreateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuRequest.ProtoReflect.Descriptor instead.
func (*CreateMenuRequest) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{4}
}

func (x *CreateMenuRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateMenuRequest) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

func (x *CreateMenuRequest) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

// CreateMenuResponse is the response for the MenuService.CreateMenu method.
type CreateMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *CreateMenuResponse) Reset() {
	*x = CreateMenuResponse{}
	mi := &file_system_menu_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuResponse) ProtoMessage() {}

func (x *CreateMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuResponse.ProtoReflect.Descriptor instead.
func (*CreateMenuResponse) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{5}
}

func (x *CreateMenuResponse) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

// UpdateMenuRequest is the request for the MenuService.UpdateMenu method.
type UpdateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The menu resource which replaces the resource on the server.
	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *UpdateMenuRequest) Reset() {
	*x = UpdateMenuRequest{}
	mi := &file_system_menu_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuRequest) ProtoMessage() {}

func (x *UpdateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuRequest.ProtoReflect.Descriptor instead.
func (*UpdateMenuRequest) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMenuRequest) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

// UpdateMenuResponse is the response for the MenuService.UpdateMenu method.
type UpdateMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu *Menu `protobuf:"bytes,1,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *UpdateMenuResponse) Reset() {
	*x = UpdateMenuResponse{}
	mi := &file_system_menu_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuResponse) ProtoMessage() {}

func (x *UpdateMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuResponse.ProtoReflect.Descriptor instead.
func (*UpdateMenuResponse) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMenuResponse) GetMenu() *Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

// DeleteMenuRequest is the request for the MenuService.DeleteMenu method.
type DeleteMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource id of the menu to be deleted, for example:
	// "shelves/shelf1/menus/menu2"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMenuRequest) Reset() {
	*x = DeleteMenuRequest{}
	mi := &file_system_menu_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuRequest) ProtoMessage() {}

func (x *DeleteMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuRequest.ProtoReflect.Descriptor instead.
func (*DeleteMenuRequest) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteMenuResponse is the response for the MenuService.DeleteMenu method.
type DeleteMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// or api.v1.services.system.Menu menu = 1; or google.protobuf.Empty empty = 1;
	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *DeleteMenuResponse) Reset() {
	*x = DeleteMenuResponse{}
	mi := &file_system_menu_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuResponse) ProtoMessage() {}

func (x *DeleteMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_menu_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuResponse.ProtoReflect.Descriptor instead.
func (*DeleteMenuResponse) Descriptor() ([]byte, []int) {
	return file_system_menu_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMenuResponse) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

var File_system_menu_proto protoreflect.FileDescriptor

var file_system_menu_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x6f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x6f, 0x6e, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe8, 0x01, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04,
	0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x76,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x6e, 0x75, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04,
	0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x45,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x46, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6d,
	0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x23, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xa0, 0x05, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x75, 0x41,
	0x50, 0x49, 0x12, 0x7b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12,
	0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x12,
	0x7a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73,
	0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22,
	0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x6d, 0x65, 0x6e,
	0x75, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x3a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x1a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x79, 0x73, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x69, 0x64, 0x7d, 0x12, 0x83, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x6d,
	0x65, 0x6e, 0x75, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xbe, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x42, 0x09, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x19, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0xa2, 0x02, 0x04, 0x41, 0x56, 0x53, 0x53, 0xaa, 0x02, 0x16, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0xca, 0x02, 0x16, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xe2, 0x02, 0x22, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x3a, 0x3a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_system_menu_proto_rawDescOnce sync.Once
	file_system_menu_proto_rawDescData = file_system_menu_proto_rawDesc
)

func file_system_menu_proto_rawDescGZIP() []byte {
	file_system_menu_proto_rawDescOnce.Do(func() {
		file_system_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_menu_proto_rawDescData)
	})
	return file_system_menu_proto_rawDescData
}

var file_system_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_system_menu_proto_goTypes = []any{
	(*ListMenusRequest)(nil),   // 0: api.v1.services.system.ListMenusRequest
	(*ListMenusResponse)(nil),  // 1: api.v1.services.system.ListMenusResponse
	(*GetMenuRequest)(nil),     // 2: api.v1.services.system.GetMenuRequest
	(*GetMenuResponse)(nil),    // 3: api.v1.services.system.GetMenuResponse
	(*CreateMenuRequest)(nil),  // 4: api.v1.services.system.CreateMenuRequest
	(*CreateMenuResponse)(nil), // 5: api.v1.services.system.CreateMenuResponse
	(*UpdateMenuRequest)(nil),  // 6: api.v1.services.system.UpdateMenuRequest
	(*UpdateMenuResponse)(nil), // 7: api.v1.services.system.UpdateMenuResponse
	(*DeleteMenuRequest)(nil),  // 8: api.v1.services.system.DeleteMenuRequest
	(*DeleteMenuResponse)(nil), // 9: api.v1.services.system.DeleteMenuResponse
	(*Menu)(nil),               // 10: api.v1.services.system.Menu
	(*anypb.Any)(nil),          // 11: google.protobuf.Any
	(*emptypb.Empty)(nil),      // 12: google.protobuf.Empty
}
var file_system_menu_proto_depIdxs = []int32{
	10, // 0: api.v1.services.system.ListMenusResponse.menus:type_name -> api.v1.services.system.Menu
	11, // 1: api.v1.services.system.ListMenusResponse.extra:type_name -> google.protobuf.Any
	10, // 2: api.v1.services.system.GetMenuResponse.menu:type_name -> api.v1.services.system.Menu
	10, // 3: api.v1.services.system.CreateMenuRequest.menu:type_name -> api.v1.services.system.Menu
	10, // 4: api.v1.services.system.CreateMenuResponse.menu:type_name -> api.v1.services.system.Menu
	10, // 5: api.v1.services.system.UpdateMenuRequest.menu:type_name -> api.v1.services.system.Menu
	10, // 6: api.v1.services.system.UpdateMenuResponse.menu:type_name -> api.v1.services.system.Menu
	12, // 7: api.v1.services.system.DeleteMenuResponse.empty:type_name -> google.protobuf.Empty
	0,  // 8: api.v1.services.system.MenuAPI.ListMenus:input_type -> api.v1.services.system.ListMenusRequest
	2,  // 9: api.v1.services.system.MenuAPI.GetMenu:input_type -> api.v1.services.system.GetMenuRequest
	4,  // 10: api.v1.services.system.MenuAPI.CreateMenu:input_type -> api.v1.services.system.CreateMenuRequest
	6,  // 11: api.v1.services.system.MenuAPI.UpdateMenu:input_type -> api.v1.services.system.UpdateMenuRequest
	8,  // 12: api.v1.services.system.MenuAPI.DeleteMenu:input_type -> api.v1.services.system.DeleteMenuRequest
	1,  // 13: api.v1.services.system.MenuAPI.ListMenus:output_type -> api.v1.services.system.ListMenusResponse
	3,  // 14: api.v1.services.system.MenuAPI.GetMenu:output_type -> api.v1.services.system.GetMenuResponse
	5,  // 15: api.v1.services.system.MenuAPI.CreateMenu:output_type -> api.v1.services.system.CreateMenuResponse
	7,  // 16: api.v1.services.system.MenuAPI.UpdateMenu:output_type -> api.v1.services.system.UpdateMenuResponse
	9,  // 17: api.v1.services.system.MenuAPI.DeleteMenu:output_type -> api.v1.services.system.DeleteMenuResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_system_menu_proto_init() }
func file_system_menu_proto_init() {
	if File_system_menu_proto != nil {
		return
	}
	file_system_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_menu_proto_goTypes,
		DependencyIndexes: file_system_menu_proto_depIdxs,
		MessageInfos:      file_system_menu_proto_msgTypes,
	}.Build()
	File_system_menu_proto = out.File
	file_system_menu_proto_rawDesc = nil
	file_system_menu_proto_goTypes = nil
	file_system_menu_proto_depIdxs = nil
}