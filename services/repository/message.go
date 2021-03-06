// Package repository implements different repository services required by the edge-cluster service
package repository

import (
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/micro-business/go-core/common"
)

// CreateEdgeClusterRequest contains the request to create a new edge cluster
type CreateEdgeClusterRequest struct {
	UserEmail   string
	EdgeCluster models.EdgeCluster
}

// CreateEdgeClusterResponse contains the result of creating a new edge cluster
type CreateEdgeClusterResponse struct {
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
	Cursor        string
}

// ReadEdgeClusterRequest contains the request to read an existing edge cluster
type ReadEdgeClusterRequest struct {
	UserEmail     string
	EdgeClusterID string
}

// ReadEdgeClusterResponse contains the result of reading an existing edge cluster
type ReadEdgeClusterResponse struct {
	EdgeCluster models.EdgeCluster
}

// UpdateEdgeClusterRequest contains the request to update an existing edge cluster
type UpdateEdgeClusterRequest struct {
	UserEmail     string
	EdgeClusterID string
	EdgeCluster   models.EdgeCluster
}

// UpdateEdgeClusterResponse contains the result of updating an existing edge cluster
type UpdateEdgeClusterResponse struct {
	EdgeCluster models.EdgeCluster
	Cursor      string
}

// DeleteEdgeClusterRequest contains the request to delete an existing edge cluster
type DeleteEdgeClusterRequest struct {
	UserEmail     string
	EdgeClusterID string
}

// DeleteEdgeClusterResponse contains the result of deleting an existing edge cluster
type DeleteEdgeClusterResponse struct {
}

// ListEdgeClustersRequest contains the filter criteria to look for existing projects
type ListEdgeClustersRequest struct {
	UserEmail      string
	Pagination     common.Pagination
	SortingOptions []common.SortingOptionPair
	EdgeClusterIDs []string
	ProjectIDs     []string
}

// ListEdgeClustersResponse contains the list of the projects that matched the result
type ListEdgeClustersResponse struct {
	HasPreviousPage bool
	HasNextPage     bool
	TotalCount      int64
	EdgeClusters    []models.EdgeClusterWithCursor
}
