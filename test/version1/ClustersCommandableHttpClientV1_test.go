package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-clusters-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type clustersCommandableHttpClientV1Test struct {
	client  *version1.ClustersCommandableHttpClientV1
	fixture *ClustersClientFixtureV1
}

func newClustersCommandableHttpClientV1Test() *clustersCommandableHttpClientV1Test {
	return &clustersCommandableHttpClientV1Test{}
}

func (c *clustersCommandableHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewClustersCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewClustersClientFixtureV1(c.client)
}

func (c *clustersCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newClustersCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}
