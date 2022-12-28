package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type AttachmentsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewAttachmentsCommandableHttpClientV1() *AttachmentsCommandableHttpClientV1 {
	return &AttachmentsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/attachments"),
	}
}

func (c *AttachmentsCommandableHttpClientV1) GetAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"id", id,
	)

	res, err := c.CallCommand(ctx, "get_attachment_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*BlobAttachmentV1](res, correlationId)
}

func (c *AttachmentsCommandableHttpClientV1) AddAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"reference", reference,
		"ids", ids,
	)

	res, err := c.CallCommand(ctx, "add_attachments", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*BlobAttachmentV1](res, correlationId)
}

func (c *AttachmentsCommandableHttpClientV1) UpdateAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, oldIds []string, newIds []string) ([]*BlobAttachmentV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"reference", reference,
		"old_ids", oldIds,
		"new_ids", newIds,
	)

	res, err := c.CallCommand(ctx, "update_attachments", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*BlobAttachmentV1](res, correlationId)
}

func (c *AttachmentsCommandableHttpClientV1) RemoveAttachments(ctx context.Context, correlationId string, reference *ReferenceV1, ids []string) ([]*BlobAttachmentV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"reference", reference,
		"ids", ids,
	)

	res, err := c.CallCommand(ctx, "remove_attachments", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]*BlobAttachmentV1](res, correlationId)
}

func (c *AttachmentsCommandableHttpClientV1) DeleteAttachmentById(ctx context.Context, correlationId string, id string) (*BlobAttachmentV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"id", id,
	)

	res, err := c.CallCommand(ctx, "delete_attachment_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*BlobAttachmentV1](res, correlationId)
}
