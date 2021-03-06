// Package grpc implements functions to expose edge-cluster service endpoint using GRPC protocol.
package grpc

import (
	"context"
	"fmt"

	edgeClusterGRPCContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// decodeCreateEdgeClusterRequest decodes CreateEdgeCluster request message from GRPC object to business object
// context: Mandatory The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeCreateEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.CreateEdgeClusterRequest)

	edgeCluster, err := mapToEdgeCluster(castedRequest.EdgeCluster)
	if err != nil {
		return nil, err
	}

	return &business.CreateEdgeClusterRequest{
		EdgeCluster: edgeCluster,
	}, nil
}

// encodeCreateEdgeClusterResponse encodes CreateEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeCreateEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.CreateEdgeClusterResponse)

	if castedResponse.Err == nil {
		edgeCluster, err := mapFromEdgeCluster(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.CreateEdgeClusterResponse{
			Error:         edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeClusterID: castedResponse.EdgeClusterID,
			EdgeCluster:   edgeCluster,
			Cursor:        castedResponse.Cursor,
		}, nil
	}

	return &edgeClusterGRPCContract.CreateEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeReadEdgeClusterRequest decodes ReadEdgeCluster request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeReadEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ReadEdgeClusterRequest)

	return &business.ReadEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
	}, nil
}

// encodeReadEdgeClusterResponse encodes ReadEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeReadEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ReadEdgeClusterResponse)

	if castedResponse.Err == nil {
		edgeCluster, err := mapFromEdgeCluster(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
			Error:           edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster:     edgeCluster,
			ProvisionDetail: mapFromProvisionDetails(castedResponse.ProvisionDetails),
		}, nil
	}

	return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeUpdateEdgeClusterRequest decodes UpdateEdgeCluster request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeUpdateEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.UpdateEdgeClusterRequest)

	edgeCluster, err := mapToEdgeCluster(castedRequest.EdgeCluster)
	if err != nil {
		return nil, err
	}

	return &business.UpdateEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		EdgeCluster:   edgeCluster,
	}, nil
}

// encodeUpdateEdgeClusterResponse encodes UpdateEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeUpdateEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.UpdateEdgeClusterResponse)

	if castedResponse.Err == nil {
		edgeCluster, err := mapFromEdgeCluster(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.UpdateEdgeClusterResponse{
			Error:       edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: edgeCluster,
			Cursor:      castedResponse.Cursor,
		}, nil
	}

	return &edgeClusterGRPCContract.UpdateEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeDeleteEdgeClusterRequest decodes DeleteEdgeCluster request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeDeleteEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.DeleteEdgeClusterRequest)

	return &business.DeleteEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
	}, nil
}

// encodeDeleteEdgeClusterResponse encodes DeleteEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeDeleteEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.DeleteEdgeClusterResponse)
	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.DeleteEdgeClusterResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
		}, nil
	}

	return &edgeClusterGRPCContract.DeleteEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClustersRequest decodes ListEdgeClusters request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrongw
func decodeListEdgeClustersRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClustersRequest)
	sortingOptions := []common.SortingOptionPair{}

	if len(castedRequest.SortingOptions) > 0 {
		sortingOptions = funk.Map(
			castedRequest.SortingOptions,
			func(sortingOption *edgeClusterGRPCContract.SortingOptionPair) common.SortingOptionPair {
				direction := common.Ascending

				if sortingOption.Direction == edgeClusterGRPCContract.SortingDirection_DESCENDING {
					direction = common.Descending
				}

				return common.SortingOptionPair{
					Name:      sortingOption.Name,
					Direction: direction,
				}
			}).([]common.SortingOptionPair)
	}

	pagination := common.Pagination{}

	if castedRequest.Pagination.HasAfter {
		pagination.After = &castedRequest.Pagination.After
	}

	if castedRequest.Pagination.HasFirst {
		first := int(castedRequest.Pagination.First)
		pagination.First = &first
	}

	if castedRequest.Pagination.HasBefore {
		pagination.Before = &castedRequest.Pagination.Before
	}

	if castedRequest.Pagination.HasLast {
		last := int(castedRequest.Pagination.Last)
		pagination.Last = &last
	}

	return &business.ListEdgeClustersRequest{
		Pagination:     pagination,
		EdgeClusterIDs: castedRequest.EdgeClusterIDs,
		ProjectIDs:     castedRequest.ProjectIDs,
		SortingOptions: sortingOptions,
	}, nil
}

// encodeListEdgeClustersResponse encodes ListEdgeClusters response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClustersResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClustersResponse)
	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClustersResponse{
			Error:           edgeClusterGRPCContract.Error_NO_ERROR,
			HasPreviousPage: castedResponse.HasPreviousPage,
			HasNextPage:     castedResponse.HasNextPage,
			TotalCount:      castedResponse.TotalCount,
			EdgeClusters: funk.Map(castedResponse.EdgeClusters, func(edgeCluster models.EdgeClusterWithCursor) *edgeClusterGRPCContract.EdgeClusterWithCursor {
				mappedEdgeCluster, _ := mapFromEdgeCluster(edgeCluster.EdgeCluster)

				return &edgeClusterGRPCContract.EdgeClusterWithCursor{
					EdgeClusterID:   edgeCluster.EdgeClusterID,
					EdgeCluster:     mappedEdgeCluster,
					Cursor:          edgeCluster.Cursor,
					ProvisionDetail: mapFromProvisionDetails(edgeCluster.ProvisionDetails),
				}
			}).([]*edgeClusterGRPCContract.EdgeClusterWithCursor),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClustersResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClusterNodesRequest decodes ListEdgeClusterNodes request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeListEdgeClusterNodesRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClusterNodesRequest)

	return &business.ListEdgeClusterNodesRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
	}, nil
}

// encodeListEdgeClusterNodesResponse encodes ListEdgeClusterNodes response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClusterNodesResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClusterNodesResponse)

	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClusterNodesResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
			Nodes: mapFromNodeStatus(castedResponse.Nodes),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClusterNodesResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClusterPodsRequest decodes ListEdgeClusterPods request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeListEdgeClusterPodsRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClusterPodsRequest)

	return &business.ListEdgeClusterPodsRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		Namespace:     castedRequest.Namespace,
		NodeName:      castedRequest.NodeName,
	}, nil
}

// encodeListEdgeClusterPodsResponse encodes ListEdgeClusterPods response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClusterPodsResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClusterPodsResponse)

	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClusterPodsResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
			Pods:  mapFromPods(castedResponse.Pods),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClusterPodsResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClusterServicesRequest decodes ListEdgeClusterServices request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeListEdgeClusterServicesRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClusterServicesRequest)

	return &business.ListEdgeClusterServicesRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		Namespace:     castedRequest.Namespace,
	}, nil
}

// encodeListEdgeClusterServicesResponse encodes ListEdgeClusterServices response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClusterServicesResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClusterServicesResponse)

	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClusterServicesResponse{
			Error:    edgeClusterGRPCContract.Error_NO_ERROR,
			Services: mapFromServices(castedResponse.Services),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClusterServicesResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

func mapError(err error) edgeClusterGRPCContract.Error {
	if commonErrors.IsUnknownError(err) {
		return edgeClusterGRPCContract.Error_UNKNOWN
	}

	if commonErrors.IsAlreadyExistsError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_ALREADY_EXISTS
	}

	if commonErrors.IsNotFoundError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_NOT_FOUND
	}

	if commonErrors.IsArgumentNilError(err) || commonErrors.IsArgumentError(err) {
		return edgeClusterGRPCContract.Error_BAD_REQUEST
	}

	return edgeClusterGRPCContract.Error_UNKNOWN
}

func mapToEdgeCluster(grpcEdgeCluster *edgeClusterGRPCContract.EdgeCluster) (edgeCluster models.EdgeCluster, err error) {
	var clusterType models.ClusterType

	if grpcEdgeCluster.ClusterType == edgeClusterGRPCContract.ClusterType_K3S {
		clusterType = models.K3S
	} else {
		err = fmt.Errorf("cluster type is not supported: %v", grpcEdgeCluster.ClusterType)

		return
	}

	edgeCluster = models.EdgeCluster{
		ProjectID:     grpcEdgeCluster.ProjectID,
		Name:          grpcEdgeCluster.Name,
		ClusterSecret: grpcEdgeCluster.ClusterSecret,
		ClusterType:   clusterType,
	}

	return
}

func mapFromEdgeCluster(edgeCluster models.EdgeCluster) (grpcEdgeCluster *edgeClusterGRPCContract.EdgeCluster, err error) {
	var clusterType edgeClusterGRPCContract.ClusterType

	if edgeCluster.ClusterType == models.K3S {
		clusterType = edgeClusterGRPCContract.ClusterType_K3S
	} else {
		err = fmt.Errorf("cluster type is not supported: %v", edgeCluster.ClusterType)

		return
	}

	grpcEdgeCluster = &edgeClusterGRPCContract.EdgeCluster{
		ProjectID:     edgeCluster.ProjectID,
		Name:          edgeCluster.Name,
		ClusterSecret: edgeCluster.ClusterSecret,
		ClusterType:   clusterType,
	}

	return
}

func mapFromNodeStatus(nodes []models.EdgeClusterNode) []*edgeClusterGRPCContract.EdgeClusterNode {
	return funk.Map(nodes, func(node models.EdgeClusterNode) *edgeClusterGRPCContract.EdgeClusterNode {
		conditions := funk.Map(node.Node.Status.Conditions, func(condition v1.NodeCondition) *edgeClusterGRPCContract.NodeCondition {
			return &edgeClusterGRPCContract.NodeCondition{
				Type:               edgeClusterGRPCContract.NodeConditionType(edgeClusterGRPCContract.NodeConditionType_value[string(condition.Type)]),
				Status:             edgeClusterGRPCContract.ConditionStatus(edgeClusterGRPCContract.ConditionStatus_value[string(condition.Status)]),
				LastHeartbeatTime:  &timestamppb.Timestamp{Seconds: condition.LastHeartbeatTime.Time.Unix()},
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			}
		}).([]*edgeClusterGRPCContract.NodeCondition)

		addresses := funk.Map(node.Node.Status.Addresses, func(address v1.NodeAddress) *edgeClusterGRPCContract.EdgeClusterNodeAddress {
			return &edgeClusterGRPCContract.EdgeClusterNodeAddress{
				NodeAddressType: edgeClusterGRPCContract.NodeAddressType(edgeClusterGRPCContract.NodeAddressType_value[string(address.Type)]),
				Address:         address.Address,
			}
		}).([]*edgeClusterGRPCContract.EdgeClusterNodeAddress)

		nodeInfo := edgeClusterGRPCContract.NodeSystemInfo{
			MachineID:               node.Node.Status.NodeInfo.MachineID,
			SystemUUID:              node.Node.Status.NodeInfo.SystemUUID,
			BootID:                  node.Node.Status.NodeInfo.BootID,
			KernelVersion:           node.Node.Status.NodeInfo.KernelVersion,
			OsImage:                 node.Node.Status.NodeInfo.OSImage,
			ContainerRuntimeVersion: node.Node.Status.NodeInfo.ContainerRuntimeVersion,
			KubeletVersion:          node.Node.Status.NodeInfo.KubeletVersion,
			KubeProxyVersion:        node.Node.Status.NodeInfo.KubeProxyVersion,
			OperatingSystem:         node.Node.Status.NodeInfo.OperatingSystem,
			Architecture:            node.Node.Status.NodeInfo.Architecture,
		}

		return &edgeClusterGRPCContract.EdgeClusterNode{
			Metadata: &edgeClusterGRPCContract.ObjectMeta{
				Id:        string(node.Node.UID),
				Name:      node.Node.Name,
				Namespace: node.Node.Namespace,
			},
			Status: &edgeClusterGRPCContract.NodeStatus{
				Conditions: conditions,
				Addresses:  addresses,
				NodeInfo:   &nodeInfo,
			},
		}
	}).([]*edgeClusterGRPCContract.EdgeClusterNode)
}

func mapFromPods(pods []models.EdgeClusterPod) []*edgeClusterGRPCContract.EdgeClusterPod {
	return funk.Map(pods, func(pod models.EdgeClusterPod) *edgeClusterGRPCContract.EdgeClusterPod {
		conditions := funk.Map(pod.Pod.Status.Conditions, func(condition v1.PodCondition) *edgeClusterGRPCContract.PodCondition {
			return &edgeClusterGRPCContract.PodCondition{
				Type:               edgeClusterGRPCContract.PodConditionType(edgeClusterGRPCContract.PodConditionType_value[string(condition.Type)]),
				Status:             edgeClusterGRPCContract.ConditionStatus(edgeClusterGRPCContract.ConditionStatus_value[string(condition.Status)]),
				LastProbeTime:      &timestamppb.Timestamp{Seconds: condition.LastProbeTime.Time.Unix()},
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			}
		}).([]*edgeClusterGRPCContract.PodCondition)

		return &edgeClusterGRPCContract.EdgeClusterPod{
			Metadata: mapFromObjectMeta(pod.Pod.ObjectMeta),
			Spec: &edgeClusterGRPCContract.PodSpec{
				NodeName: pod.Pod.Spec.NodeName,
			},
			Status: &edgeClusterGRPCContract.PodStatus{
				HostIP:     pod.Pod.Status.HostIP,
				PodIP:      pod.Pod.Status.PodIP,
				Conditions: conditions,
			},
		}
	}).([]*edgeClusterGRPCContract.EdgeClusterPod)
}

func mapFromServices(services []models.EdgeClusterService) []*edgeClusterGRPCContract.EdgeClusterService {
	return funk.Map(services, func(service models.EdgeClusterService) *edgeClusterGRPCContract.EdgeClusterService {
		conditions := funk.Map(service.Service.Status.Conditions, func(condition metav1.Condition) *edgeClusterGRPCContract.ServiceCondition {
			return &edgeClusterGRPCContract.ServiceCondition{
				Type:               condition.Type,
				Status:             edgeClusterGRPCContract.ConditionStatus(edgeClusterGRPCContract.ConditionStatus_value[string(condition.Status)]),
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			}
		}).([]*edgeClusterGRPCContract.ServiceCondition)

		ports := funk.Map(service.Service.Spec.Ports, func(port v1.ServicePort) *edgeClusterGRPCContract.ServicePort {
			return &edgeClusterGRPCContract.ServicePort{
				Name:       port.Name,
				Protcol:    edgeClusterGRPCContract.Protocol(edgeClusterGRPCContract.Protocol_value[string(port.Protocol)]),
				Port:       port.Port,
				TargetPort: port.TargetPort.String(),
				NodePort:   port.NodePort,
			}
		}).([]*edgeClusterGRPCContract.ServicePort)

		return &edgeClusterGRPCContract.EdgeClusterService{
			Metadata: mapFromObjectMeta(service.Service.ObjectMeta),
			Spec: &edgeClusterGRPCContract.ServiceSpec{
				Ports:        ports,
				Type:         edgeClusterGRPCContract.ServiceType(edgeClusterGRPCContract.ServiceType_value[string(service.Service.Spec.Type)]),
				ClusterIPs:   service.Service.Spec.ClusterIPs,
				ExternalIPs:  service.Service.Spec.ExternalIPs,
				ExternalName: service.Service.Spec.ExternalName,
			},
			Status: &edgeClusterGRPCContract.ServiceStatus{
				LoadBalancer: mapFromLoadBalancerStatus(service.Service.Status.LoadBalancer),
				Conditions:   conditions,
			},
		}
	}).([]*edgeClusterGRPCContract.EdgeClusterService)
}

func mapFromObjectMeta(objectMeta metav1.ObjectMeta) *edgeClusterGRPCContract.ObjectMeta {
	return &edgeClusterGRPCContract.ObjectMeta{
		Id:        string(objectMeta.UID),
		Name:      objectMeta.Name,
		Namespace: objectMeta.Namespace,
	}
}

func mapFromLoadBalancerStatus(loadBalancerStatus v1.LoadBalancerStatus) (mappedValues *edgeClusterGRPCContract.LoadBalancerStatus) {
	mappedValues = &edgeClusterGRPCContract.LoadBalancerStatus{
		Ingress: funk.Map(loadBalancerStatus.Ingress, func(ingress v1.LoadBalancerIngress) *edgeClusterGRPCContract.LoadBalancerIngress {
			return &edgeClusterGRPCContract.LoadBalancerIngress{
				Ip:         ingress.IP,
				Hostname:   ingress.Hostname,
				PortStatus: mapFromPortStatus(ingress.Ports),
			}
		}).([]*edgeClusterGRPCContract.LoadBalancerIngress),
	}

	return
}

func mapFromPortStatus(ports []v1.PortStatus) []*edgeClusterGRPCContract.PortStatus {
	return funk.Map(ports, func(port v1.PortStatus) *edgeClusterGRPCContract.PortStatus {
		error := ""
		if port.Error != nil {
			error = *port.Error
		}

		return &edgeClusterGRPCContract.PortStatus{
			Port:     port.Port,
			Protocol: edgeClusterGRPCContract.Protocol(edgeClusterGRPCContract.Protocol_value[string(port.Protocol)]),
			Error:    error,
		}
	}).([]*edgeClusterGRPCContract.PortStatus)
}

func mapFromProvisionDetails(from models.ProvisionDetails) (mappedValues *edgeClusterGRPCContract.ProvisionDetail) {
	provisionDetails := edgeClusterGRPCContract.ProvisionDetail{
		KubeConfigContent: from.KubeconfigContent,
	}

	if from.Service != nil {
		provisionDetails.LoadBalancer = mapFromLoadBalancerStatus(from.Service.Status.LoadBalancer)
		provisionDetails.Ports = funk.Map(from.Service.Spec.Ports, func(port v1.ServicePort) int32 {
			return port.Port
		}).([]int32)
	}

	return &provisionDetails
}
