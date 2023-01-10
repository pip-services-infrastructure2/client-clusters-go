package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type ClustersCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewClustersCommandableHttpClientV1() *ClustersCommandableHttpClientV1 {
	return &ClustersCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/clusters"),
	}
}

func (c *ClustersCommandableHttpClientV1) GetClusters(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ClusterV1], error) {
	res, err := c.CallCommand(ctx, "get_clusters", correlationId, data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *data.NewEmptyDataPage[*ClusterV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*ClusterV1]](res, correlationId)
}

func (c *ClustersCommandableHttpClientV1) GetClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error) {
	res, err := c.CallCommand(ctx, "get_cluster_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"cluster_id", clusterId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ClusterV1](res, correlationId)
}

func (c *ClustersCommandableHttpClientV1) GetClusterByTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	filter := data.NewFilterParamsFromTuples(
		"active", true,
		"tenant_id", tenantId,
	)
	page, err := c.GetClusters(ctx, correlationId, filter, nil)
	if err != nil {
		return nil, err
	}

	if len(page.Data) > 0 {
		return page.Data[0], nil
	}

	return nil, nil
}

func (c *ClustersCommandableHttpClientV1) CreateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error) {
	res, err := c.CallCommand(ctx, "create_cluster", correlationId, data.NewAnyValueMapFromTuples(
		"cluster", cluster,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ClusterV1](res, correlationId)
}

func (c *ClustersCommandableHttpClientV1) UpdateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error) {
	res, err := c.CallCommand(ctx, "update_cluster", correlationId, data.NewAnyValueMapFromTuples(
		"cluster", cluster,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ClusterV1](res, correlationId)
}

func (c *ClustersCommandableHttpClientV1) DeleteClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error) {
	res, err := c.CallCommand(ctx, "delete_cluster_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"cluster_id", clusterId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ClusterV1](res, correlationId)
}

func (c *ClustersCommandableHttpClientV1) AddTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	res, err := c.CallCommand(ctx, "add_tenant", correlationId, data.NewAnyValueMapFromTuples(
		"tenant_id", tenantId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ClusterV1](res, correlationId)
}

func (c *ClustersCommandableHttpClientV1) RemoveTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	res, err := c.CallCommand(ctx, "remove_tenant", correlationId, data.NewAnyValueMapFromTuples(
		"tenant_id", tenantId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ClusterV1](res, correlationId)
}
