package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IClustersClientV1 interface {
	GetClusters(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ClusterV1], error)

	GetClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error)

	GetClusterByTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error)

	CreateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error)

	UpdateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error)

	DeleteClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error)

	AddTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error)

	RemoveTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error)
}
