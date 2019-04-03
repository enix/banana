package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ObjectStorage : Wrapper of AWS S3 client
type ObjectStorage struct {
	Session *session.Session
	Client  *s3.S3
}

// Connect : Open a connection using the specified configuration
func (o *ObjectStorage) Connect(endpoint string, creds *credentials.Credentials) error {
	config := &aws.Config{
		Credentials:      creds,
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	}

	o.Session = session.New(config)
	o.Client = s3.New(o.Session)
	return nil
}
