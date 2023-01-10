package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ClustersNullClientV1 struct {
}

func NewClustersNullClientV1() *ClustersNullClientV1 {
	return &ClustersNullClientV1{}
}

func (c *ClustersNullClientV1) GetClusters(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ClusterV1], error) {
	return *data.NewEmptyDataPage[*ClusterV1](), nil
}

func (c *ClustersNullClientV1) GetClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error) {
	return nil, nil
}

func (c *ClustersNullClientV1) GetClusterByTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	return nil, nil
}

func (c *ClustersNullClientV1) CreateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error) {
	return cluster, nil
}

func (c *ClustersNullClientV1) UpdateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error) {
	return cluster, nil
}

func (c *ClustersNullClientV1) DeleteClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error) {
	return nil, nil
}

func (c *ClustersNullClientV1) AddTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	return nil, nil
}

func (c *ClustersNullClientV1) RemoveTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	return nil, nil
}
