package build

import (
	clients1 "github.com/pip-services-content2/client-attachments-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type AttachmentsClientFactory struct {
	*cbuild.Factory
}

func NewAttachmentsClientFactory() *AttachmentsClientFactory {
	c := &AttachmentsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-attachments", "client", "null", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-attachments", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewAttachmentsNullClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewAttachmentsCommandableHttpClientV1)

	return c
}
