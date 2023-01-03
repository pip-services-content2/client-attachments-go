package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-attachments-go/version1"
)

type AttachmentsMockClientV1 struct {
	client  *version1.AttachmentsMockClientV1
	fixture *AttachmentsClientFixtureV1
}

func newAttachmentsMockClientV1() *AttachmentsMockClientV1 {
	return &AttachmentsMockClientV1{}
}

func (c *AttachmentsMockClientV1) setup(t *testing.T) *AttachmentsClientFixtureV1 {
	c.client = version1.NewAttachmentsMockClientV1()
	c.fixture = NewAttachmentsClientFixtureV1(c.client)
	return c.fixture
}

func (c *AttachmentsMockClientV1) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newAttachmentsMockClientV1()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
