package models

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

// BackupContainer : Client representation of a bucket
type BackupContainer struct {
	Name     *string      `json:"name"`
	Contents []BackupTree `json:"contents"`
}

// NewBackupContainer : Instanciate a backup container from an S3 bucket
func NewBackupContainer(name *string) *BackupContainer {
	return &BackupContainer{
		Name: name,
	}
}

// NewFilledBackupContainer : Instanciate a backup container from an S3 bucket
//														and fill it with the given backup trees
func NewFilledBackupContainer(name *string, subtrees []BackupTree) *BackupContainer {
	return &BackupContainer{
		Name:     name,
		Contents: subtrees,
	}
}

// NewBackupContainerList : Instanciate a backup container list from an S3 bucket list
func NewBackupContainerList(list *s3.ListBucketsOutput) []BackupContainer {
	formattedList := make([]BackupContainer, len(list.Buckets))

	for index, elem := range list.Buckets {
		formattedList[index] = *NewBackupContainer(elem.Name)
	}

	return formattedList
}
