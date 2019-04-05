package models

// BackupType : full or incremental
type BackupType = int

const (
	full        BackupType = iota
	incremental BackupType = iota
)

// Backup : Client representation of a stored backup
type Backup struct {
	Timestamp int64      `json:"timestamp"`
	Type      BackupType `json:"type"`
}
