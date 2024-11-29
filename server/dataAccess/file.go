package dataAccess

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Yashsharma1911/file-store-service/server/database"
	"github.com/Yashsharma1911/file-store-service/server/models"
	"github.com/minio/minio-go/v7"
)

// FileDataAccess implementation for MinIO
type FileDataAccess struct {
	MinIOClient *database.MinIOClient
}

// NewFileDataAccess creates a new instance of FileDataAccess.
func NewFileDataAccess(minioClient *database.MinIOClient) *FileDataAccess {
	return &FileDataAccess{
		MinIOClient: minioClient,
	}
}

// UploadFile uploads a file to MinIO and returns file metadata.
func (f *FileDataAccess) UploadFile(ctx context.Context, file []byte, fileName string) (*models.FileMetadata, error) {
	// Upload file to MinIO
	_, err := f.MinIOClient.Client.PutObject(ctx, f.MinIOClient.Bucket, fileName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not upload file: %v", err)
	}

	// Return file metadata
	fileMetadata := &models.FileMetadata{
		FileName: fileName,
		Size:     int64(len(file)),
	}

	return fileMetadata, nil
}

// GetFileMetadata retrieves metadata for a file from MinIO.
func (f *FileDataAccess) GetFileMetadata(ctx context.Context, fileName string) (*models.FileMetadata, error) {
	// Check if the file exists in the bucket
	_, err := f.MinIOClient.Client.StatObject(ctx, f.MinIOClient.Bucket, fileName, minio.StatObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not retrieve file metadata: %v", err)
	}

	// Return file metadata
	fileMetadata := &models.FileMetadata{
		FileName: fileName,
	}

	return fileMetadata, nil
}

// DeleteFile removes a file from MinIO.
func (f *FileDataAccess) DeleteFile(ctx context.Context, fileName string) error {
	// Remove the file from MinIO bucket
	err := f.MinIOClient.Client.RemoveObject(ctx, f.MinIOClient.Bucket, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("could not delete file: %v", err)
	}

	return nil
}

// UpdateFile deletes the old file and uploads a new one to MinIO.
func (f *FileDataAccess) UpdateFile(ctx context.Context, oldFileName, newFileName string, newFile []byte) (*models.FileMetadata, error) {
	// First, delete the old file
	err := f.DeleteFile(ctx, oldFileName)
	if err != nil {
		return nil, fmt.Errorf("could not delete old file: %v", err)
	}

	// Now upload the new file
	_, err = f.UploadFile(ctx, newFile, newFileName)
	if err != nil {
		return nil, fmt.Errorf("could not upload new file: %v", err)
	}

	// Return new file metadata
	fileMetadata := &models.FileMetadata{
		FileName: newFileName,
		Size:     int64(len(newFile)),
	}

	return fileMetadata, nil
}

// ListFiles retrieves a list of all file metadata in the MinIO bucket.
func (f *FileDataAccess) ListFiles(ctx context.Context) ([]*models.FileMetadata, error) {
	// Initialize an empty slice to store file metadata
	var files []*models.FileMetadata

	// List objects in the MinIO bucket
	objectCh := f.MinIOClient.Client.ListObjects(ctx, f.MinIOClient.Bucket, minio.ListObjectsOptions{
		Recursive: true, // Set to true to list all objects in the bucket recursively
	})

	// Iterate over the objects and retrieve file metadata
	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("could not list objects: %v", object.Err)
		}

		// Append file metadata to the list
		fileMetadata := &models.FileMetadata{
			FileName: object.Key,
			Size:     object.Size,
		}
		files = append(files, fileMetadata)
	}

	// Return the list of file metadata
	return files, nil
}
