package storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListBuckets : List buckets on remote
func (o *ObjectStorage) ListBuckets() (*s3.ListBucketsOutput, error) {
	return o.Client.ListBuckets(nil)
}
