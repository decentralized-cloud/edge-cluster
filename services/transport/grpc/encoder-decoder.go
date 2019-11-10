// Package grpc implements functions to expose edge-cluster service endpoint using GRPC protocol.
package grpc

import (
	"context"

	edgeClusterGRPCContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
)

// decodeCreateEdgeClusterRequest decodes CreateEdgeCluster request message from GRPC object to business object
// context: Mandatory The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeCreateEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.CreateEdgeClusterRequest)

	return &business.CreateEdgeClusterRequest{
		EdgeCluster: models.EdgeCluster{
			TenantID:               castedRequest.EdgeCluster.TenantID,
			Name:                   castedRequest.EdgeCluster.Name,
			ClusterPublicIPAddress: castedRequest.EdgeCluster.ClusterPublicIPAddress,
		}}, nil
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
		return &edgeClusterGRPCContract.CreateEdgeClusterResponse{
			Error:         edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeClusterID: castedResponse.EdgeClusterID,
			EdgeCluster: &edgeClusterGRPCContract.EdgeCluster{
				TenantID:               castedResponse.EdgeCluster.TenantID,
				Name:                   castedResponse.EdgeCluster.Name,
				ClusterPublicIPAddress: castedResponse.EdgeCluster.ClusterPublicIPAddress,
			},
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
		return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: &edgeClusterGRPCContract.EdgeCluster{
				TenantID:               castedResponse.EdgeCluster.TenantID,
				Name:                   castedResponse.EdgeCluster.Name,
				ClusterPublicIPAddress: castedResponse.EdgeCluster.ClusterPublicIPAddress,
			},
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

	return &business.UpdateEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		EdgeCluster: models.EdgeCluster{
			TenantID:               castedRequest.EdgeCluster.TenantID,
			Name:                   castedRequest.EdgeCluster.Name,
			ClusterPublicIPAddress: castedRequest.EdgeCluster.ClusterPublicIPAddress,
		}}, nil
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
		return &edgeClusterGRPCContract.UpdateEdgeClusterResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: &edgeClusterGRPCContract.EdgeCluster{
				TenantID:               castedResponse.EdgeCluster.TenantID,
				Name:                   castedResponse.EdgeCluster.Name,
				ClusterPublicIPAddress: castedResponse.EdgeCluster.ClusterPublicIPAddress,
			},
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

// decodeSearchRequest decodes Search request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrongw
func decodeSearchRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.SearchRequest)
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

	return &business.SearchRequest{
		Pagination: common.Pagination{
			After:  castedRequest.Pagination.After,
			First:  int(castedRequest.Pagination.First),
			Before: castedRequest.Pagination.Before,
			Last:   int(castedRequest.Pagination.Last),
		},
		EdgeClusterIDs: castedRequest.EdgeClusterIDs,
		TenantIDs:      castedRequest.TenantIDs,
		SortingOptions: sortingOptions,
	}, nil
}

// encodeSearchResponse encodes Search response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeSearchResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.SearchResponse)
	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.SearchResponse{
			Error:           edgeClusterGRPCContract.Error_NO_ERROR,
			HasPreviousPage: castedResponse.HasPreviousPage,
			HasNextPage:     castedResponse.HasNextPage,
			EdgeClusters: funk.Map(castedResponse.EdgeClusters, func(edgeCluster models.EdgeClusterWithCursor) *edgeClusterGRPCContract.EdgeClusterWithCursor {
				return &edgeClusterGRPCContract.EdgeClusterWithCursor{
					EdgeClusterID: edgeCluster.EdgeClusterID,
					Edgecluster: &edgeClusterGRPCContract.EdgeCluster{
						TenantID:               edgeCluster.EdgeCluster.TenantID,
						Name:                   edgeCluster.EdgeCluster.Name,
						ClusterPublicIPAddress: edgeCluster.EdgeCluster.ClusterPublicIPAddress,
					},
					Cursor: edgeCluster.Cursor,
				}
			}).([]*edgeClusterGRPCContract.EdgeClusterWithCursor),
		}, nil
	}

	return &edgeClusterGRPCContract.SearchResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

func mapError(err error) edgeClusterGRPCContract.Error {
	if business.IsUnknownError(err) {
		return edgeClusterGRPCContract.Error_UNKNOWN
	}

	if business.IsEdgeClusterAlreadyExistsError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_ALREADY_EXISTS
	}

	if business.IsEdgeClusterNotFoundError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_NOT_FOUND
	}

	if commonErrors.IsArgumentNilError(err) || commonErrors.IsArgumentError(err) {
		return edgeClusterGRPCContract.Error_BAD_REQUEST
	}

	panic("Error type undefined.")
}
