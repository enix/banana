package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CreateBucket : Create a bucket
func (o *ObjectStorage) CreateBucket(name string) error {
	params := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}

	_, err := o.Client.CreateBucket(params)
	return err
}

// ListBuckets : List buckets on remote
func (o *ObjectStorage) ListBuckets() (*s3.ListBucketsOutput, error) {
	return o.Client.ListBuckets(nil)
}
