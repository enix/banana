package models

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
)

// BackupTree : Client representation of all backed up directories
//							in a given backup container
type BackupTree struct {
	Name     *string  `json:"name"`
	Contents []Backup `json:"contents"`
}

// NewBackupTree : Instanciate a backup tree from an S3 prefix
func NewBackupTree(name *string) *BackupTree {
	trimmedName := strings.TrimSuffix(*name, "/")
	return &BackupTree{
		Name: &trimmedName,
	}
}

// NewFilledBackupTree : Instanciate a backup tree from an S3 prefix
//											 and fill it with the given backups
func NewFilledBackupTree(name *string, backups []Backup) *BackupTree {
	return &BackupTree{
		Name:     name,
		Contents: backups,
	}
}

// NewBackupTreeList : Instanciate a backup tree list from an object list
func NewBackupTreeList(list *s3.ListObjectsOutput) []BackupTree {
	formattedList := make([]BackupTree, len(list.CommonPrefixes))

	for index, elem := range list.CommonPrefixes {
		formattedList[index] = *NewBackupTree(elem.Prefix)
	}

	return formattedList
}
