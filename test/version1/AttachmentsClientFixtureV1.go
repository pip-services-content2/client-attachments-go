package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-attachments-go/version1"
	"github.com/stretchr/testify/assert"
)

type AttachmentsClientFixtureV1 struct {
	Client version1.IAttachmentsClientV1

	REFERENCE1 *version1.ReferenceV1
	REFERENCE2 *version1.ReferenceV1
}

func NewAttachmentsClientFixtureV1(client version1.IAttachmentsClientV1) *AttachmentsClientFixtureV1 {
	return &AttachmentsClientFixtureV1{
		Client:     client,
		REFERENCE1: version1.NewReferenceV1("000000000000000000000011", "goal", "Goal 1"),
		REFERENCE2: version1.NewReferenceV1("000000000000000000000012", "goal", "Goal 2"),
	}
}

func (c *AttachmentsClientFixtureV1) clear() {
	for _, ref := range []*version1.ReferenceV1{c.REFERENCE1, c.REFERENCE2} {
		c.Client.RemoveAttachments(context.Background(), "", ref, []string{"1", "2", "3"})
	}
}

func (c *AttachmentsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Add attachments
	_, err := c.Client.AddAttachments(context.Background(), "123", c.REFERENCE1, []string{"1", "2"})
	assert.Nil(t, err)

	// Add other attachments
	_, err = c.Client.AddAttachments(context.Background(), "123", c.REFERENCE2, []string{"2", "3"})
	assert.Nil(t, err)

	// Check attachments has references
	item, err := c.Client.GetAttachmentById(context.Background(), "123", "2")
	assert.Nil(t, err)

	assert.NotNil(t, item)
	assert.Len(t, item.References, 2)

	// Remove reference
	_, err = c.Client.UpdateAttachments(context.Background(),
		"123",
		c.REFERENCE1,
		[]string{"1", "2"}, []string{"1"},
	)
	assert.Nil(t, err)

	// Remove another reference
	_, err = c.Client.RemoveAttachments(context.Background(),
		"123",
		c.REFERENCE1,
		[]string{"1"},
	)
	assert.Nil(t, err)

	// Remove attachments
	item, err = c.Client.DeleteAttachmentById(context.Background(), "123", "1")
	assert.Nil(t, err)
	assert.Nil(t, item)

	// Try to get deleted attachments
	item, err = c.Client.GetAttachmentById(context.Background(), "123", "2")
	assert.Nil(t, err)

	assert.NotNil(t, item)
	assert.Len(t, item.References, 1)

	reference := item.References[0]

	assert.Equal(t, "000000000000000000000012", reference.Id)
}
