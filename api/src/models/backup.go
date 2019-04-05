package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
)

// Backup : Client representation of a stored backup
type Backup struct {
	Time string `json:"time"`
	Type string `json:"type"`
}

// NewBackup : Instanciate a backup from an AWS object
//						 As there is multiple objects for a single backup,
//						 this call will return nil without error if
//						 the given object is not a backup manifest
func NewBackup(obj *s3.Object) (*Backup, error) {
	backup := &Backup{}
	metadata := strings.Split(strings.Split(*obj.Key, "/")[1], ".")
	var timeIndex int

	fmt.Println(metadata)
	if metadata[0] == "duplicity-full" {
		if len(metadata) < 4 {
			return nil, errors.New("invalid full backup name found")
		}

		if metadata[2] != "manifest" {
			return nil, nil
		}

		backup.Type = "full"
		timeIndex = 1
	} else if metadata[0] == "duplicity-inc" {
		if len(metadata) < 6 {
			return nil, errors.New("invalid inc backup name found")
		}

		if metadata[4] != "manifest" {
			return nil, nil
		}

		backup.Type = "incremental"
		timeIndex = 3
	} else {
		return nil, nil
	}

	// date, err := time.Parse("20060102T150405Z", metadata[timeIndex])
	// if err != nil {
	// 	return nil, err
	// }

	backup.Time = metadata[timeIndex]
	return backup, nil
}

// NewBackupList : Instanciate a backup list from an object list
func NewBackupList(list *s3.ListObjectsOutput) ([]Backup, error) {
	formattedList := make([]Backup, 0, len(list.Contents))

	for _, elem := range list.Contents {
		backup, err := NewBackup(elem)
		if err != nil {
			return nil, err
		}
		if backup == nil {
			continue
		}

		formattedList = append(formattedList, *backup)
	}

	return formattedList, nil
}
