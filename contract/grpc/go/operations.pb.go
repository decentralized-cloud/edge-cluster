// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: operations.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_operations_proto protoreflect.FileDescriptor

var file_operations_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a,
	0x1b, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64,
	0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x64,
	0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc1, 0x06, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x62, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x45, 0x64, 0x67, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x45,
	0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x24,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x73, 0x12,
	0x27, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x74, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_operations_proto_goTypes = []interface{}{
	(*CreateEdgeClusterRequest)(nil),        // 0: edgecluster.CreateEdgeClusterRequest
	(*ReadEdgeClusterRequest)(nil),          // 1: edgecluster.ReadEdgeClusterRequest
	(*UpdateEdgeClusterRequest)(nil),        // 2: edgecluster.UpdateEdgeClusterRequest
	(*DeleteEdgeClusterRequest)(nil),        // 3: edgecluster.DeleteEdgeClusterRequest
	(*ListEdgeClustersRequest)(nil),         // 4: edgecluster.ListEdgeClustersRequest
	(*ListEdgeClusterNodesRequest)(nil),     // 5: edgecluster.ListEdgeClusterNodesRequest
	(*ListEdgeClusterPodsRequest)(nil),      // 6: edgecluster.ListEdgeClusterPodsRequest
	(*ListEdgeClusterServicesRequest)(nil),  // 7: edgecluster.ListEdgeClusterServicesRequest
	(*CreateEdgeClusterResponse)(nil),       // 8: edgecluster.CreateEdgeClusterResponse
	(*ReadEdgeClusterResponse)(nil),         // 9: edgecluster.ReadEdgeClusterResponse
	(*UpdateEdgeClusterResponse)(nil),       // 10: edgecluster.UpdateEdgeClusterResponse
	(*DeleteEdgeClusterResponse)(nil),       // 11: edgecluster.DeleteEdgeClusterResponse
	(*ListEdgeClustersResponse)(nil),        // 12: edgecluster.ListEdgeClustersResponse
	(*ListEdgeClusterNodesResponse)(nil),    // 13: edgecluster.ListEdgeClusterNodesResponse
	(*ListEdgeClusterPodsResponse)(nil),     // 14: edgecluster.ListEdgeClusterPodsResponse
	(*ListEdgeClusterServicesResponse)(nil), // 15: edgecluster.ListEdgeClusterServicesResponse
}
var file_operations_proto_depIdxs = []int32{
	0,  // 0: edgecluster.Service.CreateEdgeCluster:input_type -> edgecluster.CreateEdgeClusterRequest
	1,  // 1: edgecluster.Service.ReadEdgeCluster:input_type -> edgecluster.ReadEdgeClusterRequest
	2,  // 2: edgecluster.Service.UpdateEdgeCluster:input_type -> edgecluster.UpdateEdgeClusterRequest
	3,  // 3: edgecluster.Service.DeleteEdgeCluster:input_type -> edgecluster.DeleteEdgeClusterRequest
	4,  // 4: edgecluster.Service.ListEdgeClusters:input_type -> edgecluster.ListEdgeClustersRequest
	5,  // 5: edgecluster.Service.ListEdgeClusterNodes:input_type -> edgecluster.ListEdgeClusterNodesRequest
	6,  // 6: edgecluster.Service.ListEdgeClusterPods:input_type -> edgecluster.ListEdgeClusterPodsRequest
	7,  // 7: edgecluster.Service.ListEdgeClusterServices:input_type -> edgecluster.ListEdgeClusterServicesRequest
	8,  // 8: edgecluster.Service.CreateEdgeCluster:output_type -> edgecluster.CreateEdgeClusterResponse
	9,  // 9: edgecluster.Service.ReadEdgeCluster:output_type -> edgecluster.ReadEdgeClusterResponse
	10, // 10: edgecluster.Service.UpdateEdgeCluster:output_type -> edgecluster.UpdateEdgeClusterResponse
	11, // 11: edgecluster.Service.DeleteEdgeCluster:output_type -> edgecluster.DeleteEdgeClusterResponse
	12, // 12: edgecluster.Service.ListEdgeClusters:output_type -> edgecluster.ListEdgeClustersResponse
	13, // 13: edgecluster.Service.ListEdgeClusterNodes:output_type -> edgecluster.ListEdgeClusterNodesResponse
	14, // 14: edgecluster.Service.ListEdgeClusterPods:output_type -> edgecluster.ListEdgeClusterPodsResponse
	15, // 15: edgecluster.Service.ListEdgeClusterServices:output_type -> edgecluster.ListEdgeClusterServicesResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_operations_proto_init() }
func file_operations_proto_init() {
	if File_operations_proto != nil {
		return
	}
	file_edge_cluster_messages_proto_init()
	file_edge_cluster_node_messages_proto_init()
	file_edge_cluster_pod_messages_proto_init()
	file_edge_cluster_service_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operations_proto_goTypes,
		DependencyIndexes: file_operations_proto_depIdxs,
	}.Build()
	File_operations_proto = out.File
	file_operations_proto_rawDesc = nil
	file_operations_proto_goTypes = nil
	file_operations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// CreateEdgeCluster creates a new edge cluster
	// request: The request to create a new edge cluster
	// Returns the result of creating new edge cluster
	CreateEdgeCluster(ctx context.Context, in *CreateEdgeClusterRequest, opts ...grpc.CallOption) (*CreateEdgeClusterResponse, error)
	// ReadEdgeCluster read an exsiting edge cluster
	// request: The request to read an esiting edge cluster
	// Returns the result of reading an existing edge cluster
	ReadEdgeCluster(ctx context.Context, in *ReadEdgeClusterRequest, opts ...grpc.CallOption) (*ReadEdgeClusterResponse, error)
	// UpdateEdgeCluster update an exsiting edge cluster
	// request: The request to update an esiting edge cluster
	// Returns the result of updateing an existing edge cluster
	UpdateEdgeCluster(ctx context.Context, in *UpdateEdgeClusterRequest, opts ...grpc.CallOption) (*UpdateEdgeClusterResponse, error)
	// DeleteEdgeCluster delete an exsiting edge cluster
	// request: The request to delete an esiting edge cluster
	// Returns the result of deleting an existing edge cluster
	DeleteEdgeCluster(ctx context.Context, in *DeleteEdgeClusterRequest, opts ...grpc.CallOption) (*DeleteEdgeClusterResponse, error)
	// ListEdgeClusters returns the list of edge clusters that matched the criteria
	// request: The request contains the search criteria
	// Returns the list of edge clusters that matched the criteria
	ListEdgeClusters(ctx context.Context, in *ListEdgeClustersRequest, opts ...grpc.CallOption) (*ListEdgeClustersResponse, error)
	// ListEdgeClusterNodes lists an existing edge cluster nodes details
	// request: The request to list an existing edge cluster nodes details
	// Returns an existing edge cluster nodes details
	ListEdgeClusterNodes(ctx context.Context, in *ListEdgeClusterNodesRequest, opts ...grpc.CallOption) (*ListEdgeClusterNodesResponse, error)
	// ListEdgeClusterPods lists an existing edge cluster pods details
	// request: The request to list an existing edge cluster pods details
	// Returns an existing edge cluster pods details
	ListEdgeClusterPods(ctx context.Context, in *ListEdgeClusterPodsRequest, opts ...grpc.CallOption) (*ListEdgeClusterPodsResponse, error)
	// ListEdgeClusterServices lists an existing edge cluster services details
	// request: The request to list an existing edge cluster services details
	// Returns an existing edge cluster services details
	ListEdgeClusterServices(ctx context.Context, in *ListEdgeClusterServicesRequest, opts ...grpc.CallOption) (*ListEdgeClusterServicesResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateEdgeCluster(ctx context.Context, in *CreateEdgeClusterRequest, opts ...grpc.CallOption) (*CreateEdgeClusterResponse, error) {
	out := new(CreateEdgeClusterResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/CreateEdgeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ReadEdgeCluster(ctx context.Context, in *ReadEdgeClusterRequest, opts ...grpc.CallOption) (*ReadEdgeClusterResponse, error) {
	out := new(ReadEdgeClusterResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/ReadEdgeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateEdgeCluster(ctx context.Context, in *UpdateEdgeClusterRequest, opts ...grpc.CallOption) (*UpdateEdgeClusterResponse, error) {
	out := new(UpdateEdgeClusterResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/UpdateEdgeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteEdgeCluster(ctx context.Context, in *DeleteEdgeClusterRequest, opts ...grpc.CallOption) (*DeleteEdgeClusterResponse, error) {
	out := new(DeleteEdgeClusterResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/DeleteEdgeCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListEdgeClusters(ctx context.Context, in *ListEdgeClustersRequest, opts ...grpc.CallOption) (*ListEdgeClustersResponse, error) {
	out := new(ListEdgeClustersResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/ListEdgeClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListEdgeClusterNodes(ctx context.Context, in *ListEdgeClusterNodesRequest, opts ...grpc.CallOption) (*ListEdgeClusterNodesResponse, error) {
	out := new(ListEdgeClusterNodesResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/ListEdgeClusterNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListEdgeClusterPods(ctx context.Context, in *ListEdgeClusterPodsRequest, opts ...grpc.CallOption) (*ListEdgeClusterPodsResponse, error) {
	out := new(ListEdgeClusterPodsResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/ListEdgeClusterPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListEdgeClusterServices(ctx context.Context, in *ListEdgeClusterServicesRequest, opts ...grpc.CallOption) (*ListEdgeClusterServicesResponse, error) {
	out := new(ListEdgeClusterServicesResponse)
	err := c.cc.Invoke(ctx, "/edgecluster.Service/ListEdgeClusterServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// CreateEdgeCluster creates a new edge cluster
	// request: The request to create a new edge cluster
	// Returns the result of creating new edge cluster
	CreateEdgeCluster(context.Context, *CreateEdgeClusterRequest) (*CreateEdgeClusterResponse, error)
	// ReadEdgeCluster read an exsiting edge cluster
	// request: The request to read an esiting edge cluster
	// Returns the result of reading an existing edge cluster
	ReadEdgeCluster(context.Context, *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error)
	// UpdateEdgeCluster update an exsiting edge cluster
	// request: The request to update an esiting edge cluster
	// Returns the result of updateing an existing edge cluster
	UpdateEdgeCluster(context.Context, *UpdateEdgeClusterRequest) (*UpdateEdgeClusterResponse, error)
	// DeleteEdgeCluster delete an exsiting edge cluster
	// request: The request to delete an esiting edge cluster
	// Returns the result of deleting an existing edge cluster
	DeleteEdgeCluster(context.Context, *DeleteEdgeClusterRequest) (*DeleteEdgeClusterResponse, error)
	// ListEdgeClusters returns the list of edge clusters that matched the criteria
	// request: The request contains the search criteria
	// Returns the list of edge clusters that matched the criteria
	ListEdgeClusters(context.Context, *ListEdgeClustersRequest) (*ListEdgeClustersResponse, error)
	// ListEdgeClusterNodes lists an existing edge cluster nodes details
	// request: The request to list an existing edge cluster nodes details
	// Returns an existing edge cluster nodes details
	ListEdgeClusterNodes(context.Context, *ListEdgeClusterNodesRequest) (*ListEdgeClusterNodesResponse, error)
	// ListEdgeClusterPods lists an existing edge cluster pods details
	// request: The request to list an existing edge cluster pods details
	// Returns an existing edge cluster pods details
	ListEdgeClusterPods(context.Context, *ListEdgeClusterPodsRequest) (*ListEdgeClusterPodsResponse, error)
	// ListEdgeClusterServices lists an existing edge cluster services details
	// request: The request to list an existing edge cluster services details
	// Returns an existing edge cluster services details
	ListEdgeClusterServices(context.Context, *ListEdgeClusterServicesRequest) (*ListEdgeClusterServicesResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) CreateEdgeCluster(context.Context, *CreateEdgeClusterRequest) (*CreateEdgeClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEdgeCluster not implemented")
}
func (*UnimplementedServiceServer) ReadEdgeCluster(context.Context, *ReadEdgeClusterRequest) (*ReadEdgeClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEdgeCluster not implemented")
}
func (*UnimplementedServiceServer) UpdateEdgeCluster(context.Context, *UpdateEdgeClusterRequest) (*UpdateEdgeClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEdgeCluster not implemented")
}
func (*UnimplementedServiceServer) DeleteEdgeCluster(context.Context, *DeleteEdgeClusterRequest) (*DeleteEdgeClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEdgeCluster not implemented")
}
func (*UnimplementedServiceServer) ListEdgeClusters(context.Context, *ListEdgeClustersRequest) (*ListEdgeClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEdgeClusters not implemented")
}
func (*UnimplementedServiceServer) ListEdgeClusterNodes(context.Context, *ListEdgeClusterNodesRequest) (*ListEdgeClusterNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEdgeClusterNodes not implemented")
}
func (*UnimplementedServiceServer) ListEdgeClusterPods(context.Context, *ListEdgeClusterPodsRequest) (*ListEdgeClusterPodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEdgeClusterPods not implemented")
}
func (*UnimplementedServiceServer) ListEdgeClusterServices(context.Context, *ListEdgeClusterServicesRequest) (*ListEdgeClusterServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEdgeClusterServices not implemented")
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_CreateEdgeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEdgeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateEdgeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/CreateEdgeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateEdgeCluster(ctx, req.(*CreateEdgeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ReadEdgeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEdgeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ReadEdgeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/ReadEdgeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ReadEdgeCluster(ctx, req.(*ReadEdgeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateEdgeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEdgeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateEdgeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/UpdateEdgeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateEdgeCluster(ctx, req.(*UpdateEdgeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteEdgeCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEdgeClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteEdgeCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/DeleteEdgeCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteEdgeCluster(ctx, req.(*DeleteEdgeClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListEdgeClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEdgeClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListEdgeClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/ListEdgeClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListEdgeClusters(ctx, req.(*ListEdgeClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListEdgeClusterNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEdgeClusterNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListEdgeClusterNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/ListEdgeClusterNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListEdgeClusterNodes(ctx, req.(*ListEdgeClusterNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListEdgeClusterPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEdgeClusterPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListEdgeClusterPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/ListEdgeClusterPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListEdgeClusterPods(ctx, req.(*ListEdgeClusterPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListEdgeClusterServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEdgeClusterServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ListEdgeClusterServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgecluster.Service/ListEdgeClusterServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ListEdgeClusterServices(ctx, req.(*ListEdgeClusterServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "edgecluster.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEdgeCluster",
			Handler:    _Service_CreateEdgeCluster_Handler,
		},
		{
			MethodName: "ReadEdgeCluster",
			Handler:    _Service_ReadEdgeCluster_Handler,
		},
		{
			MethodName: "UpdateEdgeCluster",
			Handler:    _Service_UpdateEdgeCluster_Handler,
		},
		{
			MethodName: "DeleteEdgeCluster",
			Handler:    _Service_DeleteEdgeCluster_Handler,
		},
		{
			MethodName: "ListEdgeClusters",
			Handler:    _Service_ListEdgeClusters_Handler,
		},
		{
			MethodName: "ListEdgeClusterNodes",
			Handler:    _Service_ListEdgeClusterNodes_Handler,
		},
		{
			MethodName: "ListEdgeClusterPods",
			Handler:    _Service_ListEdgeClusterPods_Handler,
		},
		{
			MethodName: "ListEdgeClusterServices",
			Handler:    _Service_ListEdgeClusterServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operations.proto",
}
