package version1

import (
	"context"
)

type IAttachmentsClientV1 interface {
	GetAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error)

	AddAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error)

	UpdateAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, oldIds []string, newIds []string) ([]*BlobAttachmentV1, error)

	RemoveAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error)

	DeleteAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error)
}
