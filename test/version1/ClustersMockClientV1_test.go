package test_version1

import (
	"testing"

	"github.com/pip-services-infrastructure2/client-clusters-go/version1"
)

type clustersMockClientV1Test struct {
	client  *version1.ClustersMockClientV1
	fixture *ClustersClientFixtureV1
}

func newClustersMockClientV1Test() *clustersMockClientV1Test {
	return &clustersMockClientV1Test{}
}

func (c *clustersMockClientV1Test) setup(t *testing.T) {
	c.client = version1.NewClustersMockClientV1()
	c.fixture = NewClustersClientFixtureV1(c.client)
}

func (c *clustersMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newClustersMockClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
