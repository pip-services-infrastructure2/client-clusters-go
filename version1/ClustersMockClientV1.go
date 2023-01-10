package version1

import (
	"context"
	"strings"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ClustersMockClientV1 struct {
	clusters []*ClusterV1
}

func NewClustersMockClientV1() *ClustersMockClientV1 {
	return &ClustersMockClientV1{
		clusters: make([]*ClusterV1, 0),
	}
}

func (c *ClustersMockClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if i < len(array2) {
				if array1[i] == array2[i] {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
}

func (c *ClustersMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *ClustersMockClientV1) matchSearch(item *ClusterV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Id, search) {
		return true
	}
	if c.matchString(item.Name, search) {
		return true
	}
	return false
}

func (c *ClustersMockClientV1) composeFilter(filter *data.FilterParams) func(*ClusterV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	name, nameOk := filter.GetAsNullableString("name")
	clusterType, clusterTypeOk := filter.GetAsNullableString("type")
	active, activeOk := filter.GetAsNullableBoolean("active")
	open, openOk := filter.GetAsNullableBoolean("open")
	tenantId, tenantIdOk := filter.GetAsNullableString("tenant_id")

	tenantIdsString := filter.GetAsString("tenant_id")
	tenantIds := make([]string, 0)

	// Process ids filter
	if tenantIdsString != "" {
		tenantIds = strings.Split(tenantIdsString, ",")
	}

	return func(item *ClusterV1) bool {
		if idOk && item.Id != id {
			return false
		}
		if tenantIdOk && (item.ActiveTenants == nil || !arrayContains(item.ActiveTenants, tenantId)) {
			return false
		}
		if len(tenantIds) > 0 && !c.contains(tenantIds, item.ActiveTenants) {
			return false
		}
		if nameOk && item.Name != name {
			return false
		}
		if clusterTypeOk && item.Type != clusterType {
			return false
		}
		if activeOk && item.Active != active {
			return false
		}
		if openOk && item.Open != open {
			return false
		}
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		return true
	}
}

func (c *ClustersMockClientV1) GetClusters(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*ClusterV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*ClusterV1, 0)
	for _, v := range c.clusters {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}
	return *data.NewDataPage(items, len(c.clusters)), nil
}

func (c *ClustersMockClientV1) GetClusterById(ctx context.Context, correlationId string, clusterId string) (result *ClusterV1, err error) {
	for _, v := range c.clusters {
		if v.Id == clusterId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func arrayContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func (c *ClustersMockClientV1) GetClusterByTenant(ctx context.Context, correlationId string, tenantId string) (result *ClusterV1, err error) {
	for _, v := range c.clusters {
		if arrayContains(v.ActiveTenants, tenantId) {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *ClustersMockClientV1) CreateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error) {
	if cluster == nil {
		return nil, nil
	}

	if cluster.Id == "" {
		cluster.Id = data.IdGenerator.NextLong()
	}

	if cluster.UpdateTime.IsZero() {
		cluster.UpdateTime = time.Now()
	}

	if !cluster.Active {
		cluster.Active = true
	}

	cluster.Open = cluster.MaxTenantCount > cluster.TenantsCount

	buf := *cluster
	c.clusters = append(c.clusters, &buf)
	return cluster, nil
}

func (c *ClustersMockClientV1) UpdateCluster(ctx context.Context, correlationId string, cluster *ClusterV1) (*ClusterV1, error) {
	if cluster == nil {
		return nil, nil
	}

	cluster.Open = cluster.MaxTenantCount > cluster.TenantsCount

	var index = -1
	for i, v := range c.clusters {
		if v.Id == cluster.Id {
			index = i
			break
		}
	}

	buf := *cluster
	c.clusters[index] = &buf

	return cluster, nil
}

func (c *ClustersMockClientV1) DeleteClusterById(ctx context.Context, correlationId string, clusterId string) (*ClusterV1, error) {
	var index = -1
	for i, v := range c.clusters {
		if v.Id == clusterId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.clusters[index]
	if index < len(c.clusters) {
		c.clusters = append(c.clusters[:index], c.clusters[index+1:]...)
	} else {
		c.clusters = c.clusters[:index]
	}

	return item, nil
}

func (c *ClustersMockClientV1) AddTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	var cluster *ClusterV1
	for _, item := range c.clusters {
		if item.Active && item.Open {
			cluster = item
			break
		}
	}

	if cluster != nil {
		if cluster.ActiveTenants == nil {
			cluster.ActiveTenants = make([]string, 0)
		}
		cluster.ActiveTenants = append(cluster.ActiveTenants, tenantId)
		cluster.TenantsCount++
		cluster.Open = cluster.MaxTenantCount > cluster.TenantsCount
	}

	return cluster, nil
}

func (c *ClustersMockClientV1) RemoveTenant(ctx context.Context, correlationId string, tenantId string) (*ClusterV1, error) {
	var cluster *ClusterV1
	for _, item := range c.clusters {
		if item.ActiveTenants != nil && arrayContains(item.ActiveTenants, tenantId) {
			cluster = item
			break
		}
	}

	activeTenants := make([]string, 0)
	if cluster != nil {
		for _, tenant := range cluster.ActiveTenants {
			if tenant != tenantId {
				activeTenants = append(activeTenants, tenant)
			}
		}
		cluster.ActiveTenants = activeTenants
		cluster.TenantsCount--
		cluster.Open = cluster.MaxTenantCount > cluster.TenantsCount
	}

	return cluster, nil

}
