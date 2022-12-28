package version1

type BlobAttachmentV1 struct {
	Id         string
	References []*ReferenceV1
}

func NewBlobAttachmentV1(id string, references []*ReferenceV1) *BlobAttachmentV1 {
	return &BlobAttachmentV1{
		Id:         id,
		References: references,
	}
}
