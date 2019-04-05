package models

// BackupContainer : Client representation of a storage bucket
type BackupContainer struct {
	Name     *string      `json:"name"`
	Contents []BackupTree `json:"contents"`
}
