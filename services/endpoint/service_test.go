package endpoint_test

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	businessMock "github.com/decentralized-cloud/edge-cluster/services/business/mock"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint"
	gokitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEndpointCreatorService(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	RegisterFailHandler(Fail)
	RunSpecs(t, "Endpoint Creator Service Tests")
}

var _ = Describe("Endpoint Creator Service Tests", func() {
	var (
		mockCtrl            *gomock.Controller
		sut                 endpoint.EndpointCreatorContract
		mockBusinessService *businessMock.MockBusinessContract
		ctx                 context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockBusinessService = businessMock.NewMockBusinessContract(mockCtrl)
		sut, _ = endpoint.NewEndpointCreatorService(mockBusinessService)
		ctx = context.WithValue(context.Background(), models.ContextKeyParsedToken, models.ParsedToken{Email: cuid.New() + "@test.com"})
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate EndpointCreatorService", func() {
		When("edge cluster business service is not provided and NewEndpointCreatorService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := endpoint.NewEndpointCreatorService(nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("businessService", "", err)
			})
		})

		When("all dependencies are resolved and NewEndpointCreatorService is called", func() {
			It("should instantiate the new EndpointCreatorService", func() {
				service, err := endpoint.NewEndpointCreatorService(mockBusinessService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("CreateEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.CreateEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint gokitendpoint.Endpoint
				request  business.CreateEdgeClusterRequest
				response business.CreateEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.CreateEdgeClusterEndpoint()
				request = business.CreateEdgeClusterRequest{
					UserEmail: cuid.New() + "@test.com",
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					},
				}

				response = business.CreateEdgeClusterResponse{
					EdgeClusterID: cuid.New(),
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					},
					Cursor: cuid.New(),
				}
			})

			Context("CreateEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.CreateEdgeClusterRequest{
							EdgeCluster: models.EdgeCluster{
								Name: "",
							}}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service CreateEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							CreateEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.CreateEdgeClusterRequest) (*business.CreateEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service CreateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service CreateEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("ReadEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.ReadEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint gokitendpoint.Endpoint
				request  business.ReadEdgeClusterRequest
				response business.ReadEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.ReadEdgeClusterEndpoint()
				request = business.ReadEdgeClusterRequest{
					UserEmail:     cuid.New() + "@test.com",
					EdgeClusterID: cuid.New(),
				}

				response = business.ReadEdgeClusterResponse{
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					},
				}
			})

			Context("ReadEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.ReadEdgeClusterRequest{
							EdgeClusterID: "",
						}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service ReadEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							ReadEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.ReadEdgeClusterRequest) (*business.ReadEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service ReadEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							ReadEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service ReadEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
							EXPECT().
							ReadEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("UpdateEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.UpdateEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint gokitendpoint.Endpoint
				request  business.UpdateEdgeClusterRequest
				response business.UpdateEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.UpdateEdgeClusterEndpoint()
				request = business.UpdateEdgeClusterRequest{
					UserEmail:     cuid.New() + "@test.com",
					EdgeClusterID: cuid.New(),
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					}}

				response = business.UpdateEdgeClusterResponse{
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					},
					Cursor: cuid.New(),
				}
			})

			Context("UpdateEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.UpdateEdgeClusterRequest{
							EdgeClusterID: "",
							EdgeCluster: models.EdgeCluster{
								Name: "",
							}}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service UpdateEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							UpdateEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.UpdateEdgeClusterRequest) (*business.UpdateEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service UpdateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							UpdateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service UpdateEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
							EXPECT().
							UpdateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("DeleteEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.DeleteEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint gokitendpoint.Endpoint
				request  business.DeleteEdgeClusterRequest
				response business.DeleteEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.DeleteEdgeClusterEndpoint()
				request = business.DeleteEdgeClusterRequest{
					UserEmail:     cuid.New() + "@test.com",
					EdgeClusterID: cuid.New(),
				}

				response = business.DeleteEdgeClusterResponse{}
			})

			Context("DeleteEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.DeleteEdgeClusterRequest{
							EdgeClusterID: "",
						}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service DeleteEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							DeleteEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.DeleteEdgeClusterRequest) (*business.DeleteEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service DeleteEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							DeleteEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service DeleteEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
							EXPECT().
							DeleteEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("ListEdgeClustersEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.ListEdgeClustersEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint       gokitendpoint.Endpoint
				edgeClusterIDs []string
				projectIDs     []string
				request        business.ListEdgeClustersRequest
				response       business.ListEdgeClustersResponse
			)

			BeforeEach(func() {
				endpoint = sut.ListEdgeClustersEndpoint()
				edgeClusterIDs = []string{}
				for idx := 0; idx < rand.Intn(20)+1; idx++ {
					edgeClusterIDs = append(edgeClusterIDs, cuid.New())
				}

				projectIDs = []string{}
				for idx := 0; idx < rand.Intn(20)+1; idx++ {
					projectIDs = append(projectIDs, cuid.New())
				}

				request = business.ListEdgeClustersRequest{
					UserEmail: cuid.New() + "@test.com",
					Pagination: common.Pagination{
						After:  convertStringToPointer(cuid.New()),
						First:  convertIntToPointer(rand.Intn(1000)),
						Before: convertStringToPointer(cuid.New()),
						Last:   convertIntToPointer(rand.Intn(1000)),
					},
					SortingOptions: []common.SortingOptionPair{
						{
							Name:      cuid.New(),
							Direction: common.Ascending,
						},
						{
							Name:      cuid.New(),
							Direction: common.Descending,
						},
					},
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
				}

				edgeClusters := []models.EdgeClusterWithCursor{}

				for idx := 0; idx < rand.Intn(20)+1; idx++ {
					edgeClusters = append(edgeClusters, models.EdgeClusterWithCursor{
						EdgeClusterID: cuid.New(),
						EdgeCluster: models.EdgeCluster{
							Name:          cuid.New(),
							ProjectID:     cuid.New(),
							ClusterSecret: cuid.New(),
							ClusterType:   models.K3S,
						},
						Cursor: cuid.New(),
					})
				}

				response = business.ListEdgeClustersResponse{
					HasPreviousPage: (rand.Intn(10) % 2) == 0,
					HasNextPage:     (rand.Intn(10) % 2) == 0,
					TotalCount:      rand.Int63n(1000),
					EdgeClusters:    edgeClusters,
				}
			})

			Context("ListEdgeClustersEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ListEdgeClustersResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ListEdgeClustersResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service ListEdgeClusters method", func() {
						mockBusinessService.
							EXPECT().
							ListEdgeClusters(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.ListEdgeClustersRequest) (*business.ListEdgeClustersResponse, error) {
									Ω(mappedRequest.Pagination).Should(Equal(request.Pagination))
									Ω(mappedRequest.SortingOptions).Should(Equal(request.SortingOptions))
									Ω(mappedRequest.EdgeClusterIDs).Should(Equal(request.EdgeClusterIDs))
									Ω(mappedRequest.ProjectIDs).Should(Equal(request.ProjectIDs))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ListEdgeClustersResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service ListEdgeClusters returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							ListEdgeClusters(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service ListEdgeClusters returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
							EXPECT().
							ListEdgeClusters(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})
})

func assertArgumentNilError(expectedArgumentName, expectedMessage string, err error) {
	Ω(commonErrors.IsArgumentNilError(err)).Should(BeTrue())

	var argumentNilErr commonErrors.ArgumentNilError
	_ = errors.As(err, &argumentNilErr)

	if expectedArgumentName != "" {
		Ω(argumentNilErr.ArgumentName).Should(Equal(expectedArgumentName))
	}

	if expectedMessage != "" {
		Ω(strings.Contains(argumentNilErr.Error(), expectedMessage)).Should(BeTrue())
	}
}

func assertArgumentError(expectedArgumentName, expectedMessage string, err error, nestedErr error) {
	Ω(commonErrors.IsArgumentError(err)).Should(BeTrue())

	var argumentErr commonErrors.ArgumentError
	_ = errors.As(err, &argumentErr)

	Ω(argumentErr.ArgumentName).Should(Equal(expectedArgumentName))
	Ω(strings.Contains(argumentErr.Error(), expectedMessage)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func convertStringToPointer(str string) *string {
	return &str
}

func convertIntToPointer(i int) *int {
	return &i
}
