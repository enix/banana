package storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

// ListTopLevelObjectsInBucket : Returns the list of "top level" objects in bucket
//															 assuming they're '/' delimited
func (o *ObjectStorage) ListTopLevelObjectsInBucket(bucket *string) (*s3.ListObjectsOutput, error) {
	delimiter := "/"
	listObjectInput := &s3.ListObjectsInput{
		Bucket:    bucket,
		Delimiter: &delimiter,
	}

	return o.Client.ListObjects(listObjectInput)
}

// ListObjectsWithPrefixInBucket : Returns the list of objects with the specified prefix
//																 assuming they're '/' delimited
func (o *ObjectStorage) ListObjectsWithPrefixInBucket(bucket *string, prefix *string) (*s3.ListObjectsOutput, error) {
	listObjectInput := &s3.ListObjectsInput{
		Bucket: bucket,
		Prefix: prefix,
	}

	return o.Client.ListObjects(listObjectInput)
}
