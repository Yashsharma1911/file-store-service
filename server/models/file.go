package models

// FileMetadata represents the metadata of a file stored in MinIO
type FileMetadata struct {
	FileName string `json:"file_name" db:"file_name"` // File name
	Size     int64  `json:"size" db:"size"`           // File size in bytes
}
