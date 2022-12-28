package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-attachments-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type AttachmentsCommandableHttpClientV1 struct {
	client  *version1.AttachmentsCommandableHttpClientV1
	fixture *AttachmentsClientFixtureV1
}

func newAttachmentsCommandableHttpClientV1() *AttachmentsCommandableHttpClientV1 {
	return &AttachmentsCommandableHttpClientV1{}
}

func (c *AttachmentsCommandableHttpClientV1) setup(t *testing.T) *AttachmentsClientFixtureV1 {
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

	c.client = version1.NewAttachmentsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewAttachmentsClientFixtureV1(c.client)

	return c.fixture
}

func (c *AttachmentsCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpCrudOperations(t *testing.T) {
	c := newAttachmentsCommandableHttpClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
