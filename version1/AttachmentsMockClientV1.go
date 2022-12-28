package version1

// import "context"

// type AttachmentsMockClientV1 struct {
// 	attachments []*BlobAttachmentV1
// }

// func NewAttachmentsMockClientV1() *AttachmentsMockClientV1 {
// 	return &AttachmentsMockClientV1{
// 		attachments: make([]*BlobAttachmentV1, 0),
// 	}
// }

// func (c *AttachmentsMockClientV1) GetAttachmentById(ctx context.Context, correlationId string, id string) (result *BlobAttachmentV1, err error) {
// 	for _, v := range c.attachments {
// 		if v.Id == id {
// 			buf := *v
// 			result = &buf
// 			break
// 		}
// 	}
// 	return result, nil
// }

// func (c *AttachmentsMockClientV1) addReference(ctx context.Context, correlationId string, id string, reference *ReferenceV1) *BlobAttachmentV1 {
// 	var item *BlobAttachmentV1

// 	for _, attach := range c.attachments {
// 		if attach.Id == id {
// 			item = attach
// 			break
// 		}
// 	}

// 	if item != nil {
// 		for i, ref := range item.References {
// 			if ref.Id == reference.Id && ref.Type == reference.Type {
// 				c.attachments = append(c.attachments[:i], c.attachments[i+1:]...)
// 			}
// 		}

// 		item.References = append(item.References, reference)
// 	} else {
// 		item = NewBlobAttachmentV1(id, []*ReferenceV1{reference})
// 		c.attachments = append(c.attachments, item)
// 	}

// 	return item
// }

// func (c *AttachmentsMockClientV1) AddAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
// 	attachments := make([]*BlobAttachmentV1, 0)

// 	// Record new references to all blobs
// 	for _, id := range ids {
// 		attachment := c.addReference(ctx, correlationId, id, reference)
// 		if attachment != nil {
// 			attachments = append(c.attachments, attachment)
// 		}
// 	}

// 	// Mark new blobs completed
// 	blobIds := make([]string, 0)
// 	for _, att := range c.attachments {
// 		if att.References != nil && len(att.References) <= 1 {
// 			blobIds = append(blobIds, att.Id)
// 		}
// 	}

// }

// func (c *AttachmentsMockClientV1) UpdateAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, oldIds []string, newIds []string) ([]*BlobAttachmentV1, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (c *AttachmentsMockClientV1) RemoveAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (c *AttachmentsMockClientV1) DeleteAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error) {
// 	panic("not implemented") // TODO: Implement
// }
