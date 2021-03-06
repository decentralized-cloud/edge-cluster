// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: edge-cluster-commons.proto

package edgecluster

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//*
// The different error types
type Error int32

const (
	// Indicates the operation was successful
	Error_NO_ERROR Error = 0
	// Indicates the operation fail with unknown error
	Error_UNKNOWN Error = 1
	// Indicates the edge cluster already exists
	Error_EDGE_CLUSTER_ALREADY_EXISTS Error = 2
	// Indicates the edge cluster does not exist
	Error_EDGE_CLUSTER_NOT_FOUND Error = 3
	// Indicates the provided values for he operation were invalid
	Error_BAD_REQUEST Error = 4
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "NO_ERROR",
		1: "UNKNOWN",
		2: "EDGE_CLUSTER_ALREADY_EXISTS",
		3: "EDGE_CLUSTER_NOT_FOUND",
		4: "BAD_REQUEST",
	}
	Error_value = map[string]int32{
		"NO_ERROR":                    0,
		"UNKNOWN":                     1,
		"EDGE_CLUSTER_ALREADY_EXISTS": 2,
		"EDGE_CLUSTER_NOT_FOUND":      3,
		"BAD_REQUEST":                 4,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_commons_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_edge_cluster_commons_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{0}
}

//*
// The different sorting direction
type SortingDirection int32

const (
	// Indicates result data must be sorted from low to high sequence
	SortingDirection_ASCENDING SortingDirection = 0
	// Indicates result data must be sorted from high to low sequence
	SortingDirection_DESCENDING SortingDirection = 1
)

// Enum value maps for SortingDirection.
var (
	SortingDirection_name = map[int32]string{
		0: "ASCENDING",
		1: "DESCENDING",
	}
	SortingDirection_value = map[string]int32{
		"ASCENDING":  0,
		"DESCENDING": 1,
	}
)

func (x SortingDirection) Enum() *SortingDirection {
	p := new(SortingDirection)
	*p = x
	return p
}

func (x SortingDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_commons_proto_enumTypes[1].Descriptor()
}

func (SortingDirection) Type() protoreflect.EnumType {
	return &file_edge_cluster_commons_proto_enumTypes[1]
}

func (x SortingDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingDirection.Descriptor instead.
func (SortingDirection) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{1}
}

//*
// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
type ConditionStatus int32

const (
	ConditionStatus_ConditionTrue    ConditionStatus = 0
	ConditionStatus_ConditionFalse   ConditionStatus = 1
	ConditionStatus_ConditionUnknown ConditionStatus = 2
)

// Enum value maps for ConditionStatus.
var (
	ConditionStatus_name = map[int32]string{
		0: "ConditionTrue",
		1: "ConditionFalse",
		2: "ConditionUnknown",
	}
	ConditionStatus_value = map[string]int32{
		"ConditionTrue":    0,
		"ConditionFalse":   1,
		"ConditionUnknown": 2,
	}
)

func (x ConditionStatus) Enum() *ConditionStatus {
	p := new(ConditionStatus)
	*p = x
	return p
}

func (x ConditionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConditionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_commons_proto_enumTypes[2].Descriptor()
}

func (ConditionStatus) Type() protoreflect.EnumType {
	return &file_edge_cluster_commons_proto_enumTypes[2]
}

func (x ConditionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConditionStatus.Descriptor instead.
func (ConditionStatus) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{2}
}

//*
//  Protocol defines network protocols.
type Protocol int32

const (
	// TCP protocol
	Protocol_TCP Protocol = 0
	// UDP protocol
	Protocol_UDP Protocol = 1
	// SCTP protocol
	Protocol_SCTP Protocol = 2
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
		2: "SCTP",
	}
	Protocol_value = map[string]int32{
		"TCP":  0,
		"UDP":  1,
		"SCTP": 2,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_commons_proto_enumTypes[3].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_edge_cluster_commons_proto_enumTypes[3]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{3}
}

//*
// The pagination information compatible with graphql-relay connection definition, for more information visit:
// https://facebook.github.io/relay/graphql/connections.htm
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasFirst  bool   `protobuf:"varint,1,opt,name=hasFirst,proto3" json:"hasFirst,omitempty"`
	First     int32  `protobuf:"varint,2,opt,name=first,proto3" json:"first,omitempty"`
	HasAfter  bool   `protobuf:"varint,3,opt,name=hasAfter,proto3" json:"hasAfter,omitempty"`
	After     string `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
	HasLast   bool   `protobuf:"varint,5,opt,name=hasLast,proto3" json:"hasLast,omitempty"`
	Last      int32  `protobuf:"varint,6,opt,name=last,proto3" json:"last,omitempty"`
	HasBefore bool   `protobuf:"varint,7,opt,name=hasBefore,proto3" json:"hasBefore,omitempty"`
	Before    string `protobuf:"bytes,8,opt,name=before,proto3" json:"before,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetHasFirst() bool {
	if x != nil {
		return x.HasFirst
	}
	return false
}

func (x *Pagination) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *Pagination) GetHasAfter() bool {
	if x != nil {
		return x.HasAfter
	}
	return false
}

func (x *Pagination) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *Pagination) GetHasLast() bool {
	if x != nil {
		return x.HasLast
	}
	return false
}

func (x *Pagination) GetLast() int32 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *Pagination) GetHasBefore() bool {
	if x != nil {
		return x.HasBefore
	}
	return false
}

func (x *Pagination) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

//*
// Defines the pair of values that are used to determine how the result data should be sorted.
type SortingOptionPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the field on
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// THe sorting direction
	Direction SortingDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=edgecluster.SortingDirection" json:"direction,omitempty"`
}

func (x *SortingOptionPair) Reset() {
	*x = SortingOptionPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortingOptionPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortingOptionPair) ProtoMessage() {}

func (x *SortingOptionPair) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortingOptionPair.ProtoReflect.Descriptor instead.
func (*SortingOptionPair) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{1}
}

func (x *SortingOptionPair) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SortingOptionPair) GetDirection() SortingDirection {
	if x != nil {
		return x.Direction
	}
	return SortingDirection_ASCENDING
}

//*
// Standard edge cluster object's metadata.
type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The namespace
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObjectMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectMeta) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// PortStatus represents the error condition of a service port
type PortStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Port is the port number of the service port of which status is recorded here
	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	// Protocol is the protocol of the service port of which status is recorded here
	Protocol Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=edgecluster.Protocol" json:"protocol,omitempty"`
	// Error is to record the problem with the service port
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PortStatus) Reset() {
	*x = PortStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortStatus) ProtoMessage() {}

func (x *PortStatus) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortStatus.ProtoReflect.Descriptor instead.
func (*PortStatus) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{3}
}

func (x *PortStatus) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PortStatus) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_TCP
}

func (x *PortStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// LoadBalancerIngress represents the status of a load-balancer ingress point:
// traffic intended for the service should be sent to an ingress point.
type LoadBalancerIngress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP is set for load-balancer ingress points that are IP based
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Hostname is set for load-balancer ingress points that are DNS based
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// PortStatus is a list of records of service ports
	// If used, every port defined in the service should have an entry in it
	PortStatus []*PortStatus `protobuf:"bytes,3,rep,name=portStatus,proto3" json:"portStatus,omitempty"`
}

func (x *LoadBalancerIngress) Reset() {
	*x = LoadBalancerIngress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerIngress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerIngress) ProtoMessage() {}

func (x *LoadBalancerIngress) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerIngress.ProtoReflect.Descriptor instead.
func (*LoadBalancerIngress) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{4}
}

func (x *LoadBalancerIngress) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LoadBalancerIngress) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *LoadBalancerIngress) GetPortStatus() []*PortStatus {
	if x != nil {
		return x.PortStatus
	}
	return nil
}

// LoadBalancerStatus represents the status of a load-balancer.
type LoadBalancerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ingress is a list containing ingress points for the load-balancer.
	// Traffic intended for the service should be sent to these ingress points.
	Ingress []*LoadBalancerIngress `protobuf:"bytes,1,rep,name=Ingress,proto3" json:"Ingress,omitempty"`
}

func (x *LoadBalancerStatus) Reset() {
	*x = LoadBalancerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerStatus) ProtoMessage() {}

func (x *LoadBalancerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerStatus.ProtoReflect.Descriptor instead.
func (*LoadBalancerStatus) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{5}
}

func (x *LoadBalancerStatus) GetIngress() []*LoadBalancerIngress {
	if x != nil {
		return x.Ingress
	}
	return nil
}

//*
// Declares the most recently observed status of the service
type ServiceCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the condition.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status is the status of the condition.
	Status ConditionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=edgecluster.ConditionStatus" json:"status,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=LastTransitionTime,proto3" json:"LastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServiceCondition) Reset() {
	*x = ServiceCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_commons_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceCondition) ProtoMessage() {}

func (x *ServiceCondition) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_commons_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceCondition.ProtoReflect.Descriptor instead.
func (*ServiceCondition) Descriptor() ([]byte, []int) {
	return file_edge_cluster_commons_proto_rawDescGZIP(), []int{6}
}

func (x *ServiceCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServiceCondition) GetStatus() ConditionStatus {
	if x != nil {
		return x.Status
	}
	return ConditionStatus_ConditionTrue
}

func (x *ServiceCondition) GetLastTransitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *ServiceCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ServiceCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_edge_cluster_commons_proto protoreflect.FileDescriptor

var file_edge_cluster_commons_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x64,
	0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0a, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68,
	0x61, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x68, 0x61, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x68, 0x61, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x68,
	0x61, 0x73, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x68, 0x61, 0x73, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x22, 0x64, 0x0a, 0x11, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x69, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x7a, 0x0a, 0x13, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x50,
	0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x22, 0xda, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x4a, 0x0a, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x70, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x2a,
	0x31, 0x0a, 0x10, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x2a, 0x4e, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x75, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x6c, 0x73, 0x65, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x02, 0x2a, 0x26, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x43, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x43, 0x54, 0x50, 0x10, 0x02, 0x42, 0x0d, 0x5a, 0x0b, 0x65, 0x64,
	0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_edge_cluster_commons_proto_rawDescOnce sync.Once
	file_edge_cluster_commons_proto_rawDescData = file_edge_cluster_commons_proto_rawDesc
)

func file_edge_cluster_commons_proto_rawDescGZIP() []byte {
	file_edge_cluster_commons_proto_rawDescOnce.Do(func() {
		file_edge_cluster_commons_proto_rawDescData = protoimpl.X.CompressGZIP(file_edge_cluster_commons_proto_rawDescData)
	})
	return file_edge_cluster_commons_proto_rawDescData
}

var file_edge_cluster_commons_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_edge_cluster_commons_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_edge_cluster_commons_proto_goTypes = []interface{}{
	(Error)(0),                    // 0: edgecluster.Error
	(SortingDirection)(0),         // 1: edgecluster.SortingDirection
	(ConditionStatus)(0),          // 2: edgecluster.ConditionStatus
	(Protocol)(0),                 // 3: edgecluster.Protocol
	(*Pagination)(nil),            // 4: edgecluster.Pagination
	(*SortingOptionPair)(nil),     // 5: edgecluster.SortingOptionPair
	(*ObjectMeta)(nil),            // 6: edgecluster.ObjectMeta
	(*PortStatus)(nil),            // 7: edgecluster.PortStatus
	(*LoadBalancerIngress)(nil),   // 8: edgecluster.LoadBalancerIngress
	(*LoadBalancerStatus)(nil),    // 9: edgecluster.LoadBalancerStatus
	(*ServiceCondition)(nil),      // 10: edgecluster.ServiceCondition
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_edge_cluster_commons_proto_depIdxs = []int32{
	1,  // 0: edgecluster.SortingOptionPair.direction:type_name -> edgecluster.SortingDirection
	3,  // 1: edgecluster.PortStatus.protocol:type_name -> edgecluster.Protocol
	7,  // 2: edgecluster.LoadBalancerIngress.portStatus:type_name -> edgecluster.PortStatus
	8,  // 3: edgecluster.LoadBalancerStatus.Ingress:type_name -> edgecluster.LoadBalancerIngress
	2,  // 4: edgecluster.ServiceCondition.status:type_name -> edgecluster.ConditionStatus
	11, // 5: edgecluster.ServiceCondition.LastTransitionTime:type_name -> google.protobuf.Timestamp
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_edge_cluster_commons_proto_init() }
func file_edge_cluster_commons_proto_init() {
	if File_edge_cluster_commons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_edge_cluster_commons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_edge_cluster_commons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortingOptionPair); i {
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
		file_edge_cluster_commons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
		file_edge_cluster_commons_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortStatus); i {
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
		file_edge_cluster_commons_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerIngress); i {
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
		file_edge_cluster_commons_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerStatus); i {
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
		file_edge_cluster_commons_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceCondition); i {
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
			RawDescriptor: file_edge_cluster_commons_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_edge_cluster_commons_proto_goTypes,
		DependencyIndexes: file_edge_cluster_commons_proto_depIdxs,
		EnumInfos:         file_edge_cluster_commons_proto_enumTypes,
		MessageInfos:      file_edge_cluster_commons_proto_msgTypes,
	}.Build()
	File_edge_cluster_commons_proto = out.File
	file_edge_cluster_commons_proto_rawDesc = nil
	file_edge_cluster_commons_proto_goTypes = nil
	file_edge_cluster_commons_proto_depIdxs = nil
}
