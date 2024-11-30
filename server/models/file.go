package models

// FileMetadata represents the metadata of a file stored in MinIO
type FileMetadata struct {
	FileName string `json:"file_name" db:"file_name"`
	Size     int64  `json:"size" db:"size"`
	CheckSum string `json:"check_sum" db:"check_sum"`
}
