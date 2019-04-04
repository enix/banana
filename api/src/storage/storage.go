package storage

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ObjectStorage : Wrapper of AWS S3 client
type ObjectStorage struct {
	Session *session.Session
	Client  *s3.S3
}
