package version1

import "time"

type ClusterV1 struct {
	Id string `json:"id"`

	Name   string `json:"name"`
	Type   string `json:"type"`
	Active bool   `json:"active"`

	MasterNodes []string `json:"master_nodes"`
	SlaveNodes  []string `json:"slave_nodes"`

	ApiHost      string            `json:"api_host"`
	ServicePorts map[string]uint16 `json:"service_port"`

	Maintenance bool      `json:"maintenance"`
	Version     string    `json:"version"`
	UpdateTime  time.Time `json:"update_time"`

	MaxTenantCount  int      `json:"max_tenant_count"`
	TenantsCount    int      `json:"tenants_count"`
	Open            bool     `json:"open"`
	ActiveTenants   []string `json:"active_tenants"`
	InactiveTenants []string `json:"inactive_tenants"`
}
