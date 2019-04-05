package models

// BackupTree : Client representation of all backups for a given directory
type BackupTree struct {
	Name     *string  `json:"name"`
	Contents []Backup `json:"contents"`
}
