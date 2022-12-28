package version1

import "context"

type AttachmentsNullClientV1 struct {
}

func NewAttachmentsNullClientV1() *AttachmentsNullClientV1 {
	return &AttachmentsNullClientV1{}
}

func (c *AttachmentsNullClientV1) GetAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error) {
	return nil, nil
}

func (c *AttachmentsNullClientV1) AddAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
	return make([]*BlobAttachmentV1, 0), nil
}

func (c *AttachmentsNullClientV1) UpdateAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, oldIds []string, newIds []string) ([]*BlobAttachmentV1, error) {
	return make([]*BlobAttachmentV1, 0), nil
}

func (c *AttachmentsNullClientV1) RemoveAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
	return make([]*BlobAttachmentV1, 0), nil
}

func (c *AttachmentsNullClientV1) DeleteAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error) {
	return nil, nil
}
