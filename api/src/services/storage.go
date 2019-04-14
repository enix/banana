package services

import (
	"errors"
	"os"

	"enix.io/banana/src/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Storage : Use this API to interact with object storage
var Storage *ObjectStorage

// ObjectStorage : Wrapper of AWS S3 client
type ObjectStorage struct {
	Session *session.Session
	Client  *s3.S3
}

// Connect : Open a connection using the specified configuration
func (o *ObjectStorage) Connect(endpoint string, creds *credentials.Credentials) {
	config := &aws.Config{
		Credentials:      creds,
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	}

	o.Session = session.New(config)
	o.Client = s3.New(o.Session)
}

// ListBuckets : List buckets on remote
func (o *ObjectStorage) ListBuckets() (*s3.ListBucketsOutput, error) {
	return o.Client.ListBuckets(nil)
}

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

// OpenStorageConnection : Initialize and test connection with object storage
func OpenStorageConnection() error {
	accessToken, err := Vault.GetStorageAccessToken()
	if err != nil {
		return err
	}
	secretToken, err := Vault.GetStorageSecretToken()
	if err != nil {
		return err
	}

	Storage = &ObjectStorage{}
	Storage.Connect(
		os.Getenv("API_ENDPOINT"),
		credentials.NewStaticCredentials(accessToken, secretToken, ""),
	)

	_, err = Storage.ListBuckets()
	if err != nil {
		return errors.New("fatal: failed to list buckets from remote. configuration error?")
	}

	logger.Log("etablished connection with object storage")
	return nil
}
