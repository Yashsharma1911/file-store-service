package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIOClient holds the MinIO client instance and the bucket name.
type MinIOClient struct {
	Client *minio.Client
	Bucket string
}

// NewMinIOClient initializes and returns a new MinIO client.
func NewMinIOClient() (*MinIOClient, error) {
	fmt.Println("Database initializing...")

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ROOT_USER")
	secretKey := os.Getenv("MINIO_ROOT_PASSWORD")
	bucket := os.Getenv("MINIO_BUCKET")
	useSSL := os.Getenv("MINIO_USE_SSL")

	ssl, _ := strconv.ParseBool(useSSL)
	// Initialize MinIO client
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl, // Set to true because using HTTPS
	})
	if err != nil {
		return nil, fmt.Errorf("could not create MinIO client: %v", err)
	}

	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		fmt.Errorf("error checking bucket existence: %v", err)
		return nil, err
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Errorf("error creating bucket: %v", err)
			return nil, err
		}
		fmt.Printf("DB created successfully")
	}

	fmt.Println("Connected with existing DB")

	// Return MinIO client
	return &MinIOClient{
		Client: client,
		Bucket: bucket,
	}, nil
}
