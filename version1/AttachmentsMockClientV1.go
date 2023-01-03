package version1

import (
	"context"

	blobClients "github.com/pip-services-infrastructure2/client-blobs-go/version1"
)

type AttachmentsMockClientV1 struct {
	attachments []*BlobAttachmentV1
	blobClient  *blobClients.BlobsMockClientV1
}

func NewAttachmentsMockClientV1() *AttachmentsMockClientV1 {
	return &AttachmentsMockClientV1{
		attachments: make([]*BlobAttachmentV1, 0),
		blobClient:  blobClients.NewBlobsMockClientV1(),
	}
}

func (c *AttachmentsMockClientV1) GetAttachmentById(ctx context.Context, correlationId string, id string) (result *BlobAttachmentV1, err error) {
	for _, v := range c.attachments {
		if v.Id == id {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *AttachmentsMockClientV1) addReference(ctx context.Context, correlationId string, id string, reference *ReferenceV1) *BlobAttachmentV1 {
	var item *BlobAttachmentV1

	for _, attach := range c.attachments {
		if attach.Id == id {
			item = attach
			break
		}
	}

	if item != nil {
		for i, ref := range item.References {
			if ref.Id == reference.Id && ref.Type == reference.Type {
				c.attachments = append(c.attachments[:i], c.attachments[i+1:]...)
			}
		}

		item.References = append(item.References, reference)
	} else {
		item = NewBlobAttachmentV1(id, []*ReferenceV1{reference})
		c.attachments = append(c.attachments, item)
	}

	return item
}

func (c *AttachmentsMockClientV1) AddAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
	attachments := make([]*BlobAttachmentV1, 0)

	// Record new references to all blobs
	for _, id := range ids {
		attachment := c.addReference(ctx, correlationId, id, reference)
		if attachment != nil {
			attachments = append(attachments, attachment)
		}
	}

	// Mark new blobs completed
	blobIds := make([]string, 0)
	for _, att := range c.attachments {
		if att.References != nil && len(att.References) <= 1 {
			blobIds = append(blobIds, att.Id)
		}
	}

	if len(blobIds) > 0 {
		err := c.blobClient.MarkBlobsCompleted(ctx, correlationId, blobIds)
		if err != nil {
			return nil, err
		}
	}

	return attachments, nil
}

func (c *AttachmentsMockClientV1) UpdateAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, oldIds []string, newIds []string) ([]*BlobAttachmentV1, error) {
	attachments := make([]*BlobAttachmentV1, 0)

	ids := make([]string, 0)
	for _, oldId := range oldIds {
		include := false
		for _, newId := range newIds {
			if oldId == newId {
				include = true
				break
			}
		}

		if !include {
			ids = append(ids, oldId)
		}
	}

	// Remove obsolete ids
	if len(ids) > 0 {
		removedAttachments, err := c.RemoveAttachments(ctx, correlationId, reference, ids)

		if err != nil {
			return nil, err
		}

		attachments = append(attachments, removedAttachments...)
	}

	ids = make([]string, 0)

	for _, newId := range newIds {
		include := false
		for _, oldId := range oldIds {
			if newId == oldId {
				include = true
				break
			}
		}

		if !include {
			ids = append(ids, newId)
		}
	}

	// Add new ids
	if len(ids) > 0 {
		addAttachments, err := c.AddAttachments(ctx, correlationId, reference, ids)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, addAttachments...)
	}

	return attachments, nil
}

func (c *AttachmentsMockClientV1) RemoveAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
	attachments := make([]*BlobAttachmentV1, 0)

	for _, id := range ids {
		attachment, err := c.removeReference(ctx, correlationId, id, reference)
		if err != nil {
			return nil, err
		}

		if attachment != nil {
			attachments = append(attachments, attachment)

			if attachment.References == nil || len(attachment.References) == 0 {
				_, err := c.DeleteAttachmentById(ctx, correlationId, attachment.Id)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return attachments, nil
}

func (c *AttachmentsMockClientV1) removeReference(ctx context.Context, correlationId string, id string, reference *ReferenceV1) (*BlobAttachmentV1, error) {
	var item *BlobAttachmentV1

	for _, attachment := range c.attachments {
		if attachment.Id == id {
			item = attachment
			break
		}
	}

	if item != nil {
		for i, ref := range item.References {
			if ref.Id == reference.Id && ref.Type == reference.Type {
				item.References = append(item.References[:i], item.References[i+1:]...)
			}
		}
	}

	return item, nil
}

func (c *AttachmentsMockClientV1) DeleteAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error) {
	var attachment *BlobAttachmentV1

	for index, item := range c.attachments {
		if item.Id == id {
			attachment = item
			c.attachments = append(c.attachments[:index], c.attachments[index+1:]...)
			break
		}
	}

	err := c.blobClient.DeleteBlobById(ctx, correlationId, id)
	if err != nil {
		return nil, err
	}

	return attachment, nil
}
