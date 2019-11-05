// Package repository implements different repository services required by the edge-cluster service
package repository

import (
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/micro-business/go-core/common"
)

// CreateEdgeClusterRequest contains the request to create a new edge cluster
type CreateEdgeClusterRequest struct {
	EdgeCluster models.EdgeCluster
}

// CreateEdgeClusterResponse contains the result of creating a new edge cluster
type CreateEdgeClusterResponse struct {
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// ReadEdgeClusterRequest contains the request to read an existing edge cluster
type ReadEdgeClusterRequest struct {
	EdgeClusterID string
}

// ReadEdgeClusterResponse contains the result of reading an existing edge cluster
type ReadEdgeClusterResponse struct {
	EdgeCluster models.EdgeCluster
}

// UpdateEdgeClusterRequest contains the request to update an existing edge cluster
type UpdateEdgeClusterRequest struct {
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// UpdateEdgeClusterResponse contains the result of updating an existing edge cluster
type UpdateEdgeClusterResponse struct {
	EdgeCluster models.EdgeCluster
}

// DeleteEdgeClusterRequest contains the request to delete an existing edge cluster
type DeleteEdgeClusterRequest struct {
	EdgeClusterID string
}

// DeleteEdgeClusterResponse contains the result of deleting an existing edge cluster
type DeleteEdgeClusterResponse struct {
}

// SearchRequest contains the filter criteria to look for existing tenants
type SearchRequest struct {
	Pagination     common.Pagination
	SortingOptions []common.SortingOptionPair
	EdgeClusterIDs []string
}

// SearchResponse contains the list of the tenants that matched the result
type SearchResponse struct {
	HasPreviousPage bool
	HasNextPage     bool
	EdgeClusters    []models.EdgeClusterWithCursor
}
