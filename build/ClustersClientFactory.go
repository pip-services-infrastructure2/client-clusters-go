package build

import (
	clients1 "github.com/pip-services-infrastructure2/client-clusters-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type ClustersClientFactory struct {
	*cbuild.Factory
}

func NewClustersClientFactory() *ClustersClientFactory {
	c := &ClustersClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-clusters", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-clusters", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-clusters", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewClustersNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewClustersMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewClustersCommandableHttpClientV1)

	return c
}
