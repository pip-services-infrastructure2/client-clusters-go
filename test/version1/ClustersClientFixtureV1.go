package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-infrastructure2/client-clusters-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type ClustersClientFixtureV1 struct {
	Client   version1.IClustersClientV1
	CLUSTER1 *version1.ClusterV1
	CLUSTER2 *version1.ClusterV1
}

func NewClustersClientFixtureV1(client version1.IClustersClientV1) *ClustersClientFixtureV1 {
	return &ClustersClientFixtureV1{
		Client: client,
		CLUSTER1: &version1.ClusterV1{
			Id:             "1",
			Name:           "Cluster #1",
			Type:           "root",
			Active:         true,
			ApiHost:        "api.mycluster1.com",
			ServicePorts:   map[string]uint16{"myservice1": 30001, "myservice2": 30002},
			MaxTenantCount: 1,
			TenantsCount:   1,
			ActiveTenants:  []string{"1"},
		},
		CLUSTER2: &version1.ClusterV1{
			Id:              "2",
			Name:            "Cluster #2",
			Type:            "tenants",
			Active:          true,
			ApiHost:         "api.mycluster2.com",
			ServicePorts:    map[string]uint16{"myservice1": 30001, "myservice2": 30002},
			MaxTenantCount:  10,
			TenantsCount:    4,
			ActiveTenants:   []string{"2", "3"},
			InactiveTenants: []string{"4"},
		},
	}
}

func (c *ClustersClientFixtureV1) TestCrudOperations(t *testing.T) {
	// Create one cluster
	cluster1, err := c.Client.CreateCluster(context.Background(), "123", c.CLUSTER1)
	assert.Nil(t, err)

	assert.NotNil(t, cluster1)
	assert.Equal(t, cluster1.Name, c.CLUSTER1.Name)
	assert.Equal(t, cluster1.ApiHost, c.CLUSTER1.ApiHost)

	// Create another cluster
	cluster2, err := c.Client.CreateCluster(context.Background(), "123", c.CLUSTER2)
	assert.Nil(t, err)

	assert.NotNil(t, cluster2)
	assert.Equal(t, cluster2.Name, c.CLUSTER2.Name)
	assert.Equal(t, cluster2.ApiHost, c.CLUSTER2.ApiHost)

	// Get all clusters
	clusters, err := c.Client.GetClusters(context.Background(), "123", nil, data.NewPagingParams(0, 5, false))
	assert.Nil(t, err)

	assert.True(t, len(clusters.Data) >= 2)

	// Update the cluster
	cluster1.Active = false
	cluster1.MaxTenantCount = 2
	cluster1.TenantsCount = 2

	cluster, err := c.Client.UpdateCluster(context.Background(), "123", cluster1)
	assert.Nil(t, err)

	assert.NotNil(t, cluster)
	assert.False(t, cluster.Active)
	assert.Equal(t, cluster.MaxTenantCount, 2)
	assert.Equal(t, cluster.TenantsCount, 2)
	assert.False(t, cluster.Open)

	cluster1 = cluster

	// Add tenant to cluster
	cluster, err = c.Client.AddTenant(context.Background(), "123", "5")
	assert.Nil(t, err)

	assert.NotNil(t, cluster)
	assert.True(t, cluster.Active)
	assert.Contains(t, cluster.ActiveTenants, "5")

	// Get cluster by tenant
	cluster, err = c.Client.GetClusterByTenant(context.Background(), "123", "5")
	assert.Nil(t, err)

	assert.NotNil(t, cluster)
	assert.True(t, cluster.Active)
	assert.Contains(t, cluster.ActiveTenants, "5")

	// Remove tenant from cluster
	cluster, err = c.Client.RemoveTenant(context.Background(), "123", "5")
	assert.Nil(t, err)

	assert.NotNil(t, cluster)
	assert.NotContains(t, cluster.ActiveTenants, "5")

	// Delete cluster
	_, err = c.Client.DeleteClusterById(context.Background(), "123", cluster1.Id)
	assert.Nil(t, err)

	// Try to get delete cluster
	cluster, err = c.Client.GetClusterById(context.Background(), "123", cluster1.Id)

	assert.Nil(t, err)
	assert.Nil(t, cluster)
}
