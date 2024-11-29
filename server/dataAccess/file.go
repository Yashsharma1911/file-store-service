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

/**
* Creates new instance of FileData Access
* File data access is struct which allows to perform operations with MinIo
* using MinIo client sdk
 */
func NewFileDataAccess(minioClient *database.MinIOClient) *FileDataAccess {
	return &FileDataAccess{
		MinIOClient: minioClient,
	}
}

// UploadFile uploads a file to MinIO and returns file metadata.
func (f *FileDataAccess) UploadFile(ctx context.Context, file []byte, fileName string) (*models.FileMetadata, error) {
	_, err := f.MinIOClient.Client.PutObject(ctx, f.MinIOClient.Bucket, fileName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not upload file: %v", err)
	}

	fileMetadata := &models.FileMetadata{
		FileName: fileName,
		Size:     int64(len(file)),
	}

	return fileMetadata, nil
}

// GetFileMetadata retrieves metadata for a file from MinIO.
func (f *FileDataAccess) GetFileMetadata(ctx context.Context, fileName string) (*models.FileMetadata, error) {
	object, err := f.MinIOClient.Client.StatObject(ctx, f.MinIOClient.Bucket, fileName, minio.StatObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("file does not exist: %v", err)
	}

	fileMetadata := &models.FileMetadata{
		FileName: object.Key,
		Size:     object.Size,
	}

	return fileMetadata, nil
}

// DeleteFile removes a file from MinIO bucket.
func (f *FileDataAccess) DeleteFile(ctx context.Context, fileName string) error {
	_, errorFindingFile := f.GetFileMetadata(ctx, fileName)
	if errorFindingFile != nil {
		return fmt.Errorf("file does not exist")
	}
	err := f.MinIOClient.Client.RemoveObject(ctx, f.MinIOClient.Bucket, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("could not delete file: %v", err)
	}

	return nil
}

// UpdateFile deletes the old file and uploads a new one to MinIO.
func (f *FileDataAccess) UpdateFile(ctx context.Context, oldFileName, newFileName string, newFile []byte) (*models.FileMetadata, error) {
	// If file already exit delete before updating it
	_ = f.DeleteFile(ctx, oldFileName)

	// UploadFile methods update existing file and if it is not present
	// it will store the given file
	_, err := f.UploadFile(ctx, newFile, newFileName)
	if err != nil {
		return nil, fmt.Errorf("could not upload new file: %v", err)
	}

	fileMetadata := &models.FileMetadata{
		FileName: newFileName,
		Size:     int64(len(newFile)),
	}

	return fileMetadata, nil
}

// ListFiles retrieves a list of all file metadata in the MinIO bucket.
func (f *FileDataAccess) ListFiles(ctx context.Context) ([]*models.FileMetadata, error) {
	var files []*models.FileMetadata

	// Find all the files in the bucket recursively
	objectCh := f.MinIOClient.Client.ListObjects(ctx, f.MinIOClient.Bucket, minio.ListObjectsOptions{
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("could not list objects: %v", object.Err)
		}

		fileMetadata := &models.FileMetadata{
			FileName: object.Key,
			Size:     object.Size,
		}
		files = append(files, fileMetadata)
	}

	return files, nil
}
