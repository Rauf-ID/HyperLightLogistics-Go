//
// This file is part of HyperLightLogistics-Go.
//
// HyperLightLogistics-Go is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// HyperLightLogistics-Go is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with HyperLightLogistics-Go.  If not, see <https://www.gnu.org/licenses/>.
//
// Copyright (C) 2024 Rauf Agaguliev

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: deliveryOptions.proto

package proto

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity  uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_deliveryOptions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryOptions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_deliveryOptions_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DeliveryAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	State   string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Street  string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	Zipcode string `protobuf:"bytes,5,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
}

func (x *DeliveryAddress) Reset() {
	*x = DeliveryAddress{}
	mi := &file_deliveryOptions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryAddress) ProtoMessage() {}

func (x *DeliveryAddress) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryOptions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryAddress.ProtoReflect.Descriptor instead.
func (*DeliveryAddress) Descriptor() ([]byte, []int) {
	return file_deliveryOptions_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *DeliveryAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *DeliveryAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *DeliveryAddress) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *DeliveryAddress) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

type DeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId      uint64           `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	DeliveryAddress *DeliveryAddress `protobuf:"bytes,2,opt,name=deliveryAddress,proto3" json:"deliveryAddress,omitempty"`
	Products        []*Product       `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *DeliveryRequest) Reset() {
	*x = DeliveryRequest{}
	mi := &file_deliveryOptions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRequest) ProtoMessage() {}

func (x *DeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryOptions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRequest.ProtoReflect.Descriptor instead.
func (*DeliveryRequest) Descriptor() ([]byte, []int) {
	return file_deliveryOptions_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryRequest) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *DeliveryRequest) GetDeliveryAddress() *DeliveryAddress {
	if x != nil {
		return x.DeliveryAddress
	}
	return nil
}

func (x *DeliveryRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type DeliveryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	DeliveryTime string  `protobuf:"bytes,2,opt,name=deliveryTime,proto3" json:"deliveryTime,omitempty"`
	Price        float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *DeliveryOptions) Reset() {
	*x = DeliveryOptions{}
	mi := &file_deliveryOptions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryOptions) ProtoMessage() {}

func (x *DeliveryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryOptions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryOptions.ProtoReflect.Descriptor instead.
func (*DeliveryOptions) Descriptor() ([]byte, []int) {
	return file_deliveryOptions_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryOptions) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DeliveryOptions) GetDeliveryTime() string {
	if x != nil {
		return x.DeliveryTime
	}
	return ""
}

func (x *DeliveryOptions) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductDeliveryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId       uint64             `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	DeliveryOptions []*DeliveryOptions `protobuf:"bytes,2,rep,name=deliveryOptions,proto3" json:"deliveryOptions,omitempty"`
}

func (x *ProductDeliveryOptions) Reset() {
	*x = ProductDeliveryOptions{}
	mi := &file_deliveryOptions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductDeliveryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDeliveryOptions) ProtoMessage() {}

func (x *ProductDeliveryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryOptions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDeliveryOptions.ProtoReflect.Descriptor instead.
func (*ProductDeliveryOptions) Descriptor() ([]byte, []int) {
	return file_deliveryOptions_proto_rawDescGZIP(), []int{4}
}

func (x *ProductDeliveryOptions) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductDeliveryOptions) GetDeliveryOptions() []*DeliveryOptions {
	if x != nil {
		return x.DeliveryOptions
	}
	return nil
}

type DeliveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductDeliveryOptions `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *DeliveryResponse) Reset() {
	*x = DeliveryResponse{}
	mi := &file_deliveryOptions_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryResponse) ProtoMessage() {}

func (x *DeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryOptions_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryResponse.ProtoReflect.Descriptor instead.
func (*DeliveryResponse) Descriptor() ([]byte, []int) {
	return file_deliveryOptions_proto_rawDescGZIP(), []int{5}
}

func (x *DeliveryResponse) GetProducts() []*ProductDeliveryOptions {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_deliveryOptions_proto protoreflect.FileDescriptor

var file_deliveryOptions_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x9f, 0x01,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x40, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x5f, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x78, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4d, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x67, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x18, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x24, 0x50, 0x01, 0x5a, 0x20, 0x48, 0x79, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x47, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deliveryOptions_proto_rawDescOnce sync.Once
	file_deliveryOptions_proto_rawDescData = file_deliveryOptions_proto_rawDesc
)

func file_deliveryOptions_proto_rawDescGZIP() []byte {
	file_deliveryOptions_proto_rawDescOnce.Do(func() {
		file_deliveryOptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_deliveryOptions_proto_rawDescData)
	})
	return file_deliveryOptions_proto_rawDescData
}

var file_deliveryOptions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_deliveryOptions_proto_goTypes = []any{
	(*Product)(nil),                // 0: proto.Product
	(*DeliveryAddress)(nil),        // 1: proto.DeliveryAddress
	(*DeliveryRequest)(nil),        // 2: proto.DeliveryRequest
	(*DeliveryOptions)(nil),        // 3: proto.DeliveryOptions
	(*ProductDeliveryOptions)(nil), // 4: proto.ProductDeliveryOptions
	(*DeliveryResponse)(nil),       // 5: proto.DeliveryResponse
}
var file_deliveryOptions_proto_depIdxs = []int32{
	1, // 0: proto.DeliveryRequest.deliveryAddress:type_name -> proto.DeliveryAddress
	0, // 1: proto.DeliveryRequest.products:type_name -> proto.Product
	3, // 2: proto.ProductDeliveryOptions.deliveryOptions:type_name -> proto.DeliveryOptions
	4, // 3: proto.DeliveryResponse.products:type_name -> proto.ProductDeliveryOptions
	2, // 4: proto.DeliveryOptionsService.CalculateDeliveryOptions:input_type -> proto.DeliveryRequest
	5, // 5: proto.DeliveryOptionsService.CalculateDeliveryOptions:output_type -> proto.DeliveryResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_deliveryOptions_proto_init() }
func file_deliveryOptions_proto_init() {
	if File_deliveryOptions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deliveryOptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deliveryOptions_proto_goTypes,
		DependencyIndexes: file_deliveryOptions_proto_depIdxs,
		MessageInfos:      file_deliveryOptions_proto_msgTypes,
	}.Build()
	File_deliveryOptions_proto = out.File
	file_deliveryOptions_proto_rawDesc = nil
	file_deliveryOptions_proto_goTypes = nil
	file_deliveryOptions_proto_depIdxs = nil
}
