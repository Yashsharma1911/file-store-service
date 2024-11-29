package database

import (
	"fmt"
	"os"

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
	endpoint := os.Getenv("MINIO_ENDPOINT")    // E.g., "localhost:9000"
	accessKey := os.Getenv("MINIO_ACCESS_KEY") // E.g., "minioaccesskey"
	secretKey := os.Getenv("MINIO_SECRET_KEY") // E.g., "miniosecretkey"
	bucket := os.Getenv("MINIO_BUCKET")        // E.g., "my-bucket"

	// If any environment variables are missing, you can hardcode them temporarily for testing purposes
	if endpoint == "" {
		endpoint = "localhost:9000" // default MinIO endpoint for local setup
	}
	if accessKey == "" {
		accessKey = "minioaccesskey" // default access key for local MinIO
	}
	if secretKey == "" {
		secretKey = "miniosecretkey" // default secret key for local MinIO
	}
	if bucket == "" {
		bucket = "my-bucket" // default bucket name for local MinIO
	}
	// Initialize MinIO client
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // Set to true if using HTTPS
	})
	if err != nil {
		return nil, fmt.Errorf("could not create MinIO client: %v", err)
	}

	// exists, err := client.BucketExists(context.Background(), bucket)
	// if err != nil {
	// 	log.Fatalf("Error checking bucket existence: %v", err)
	// 	return nil, err
	// }
	// if !exists {
	// 	err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
	// 	if err != nil {
	// 		log.Fatalf("Error creating bucket: %v", err)
	// 		return nil, err
	// 	}
	// 	log.Printf("Bucket created successfully")
	// }

	// Return MinIO client
	return &MinIOClient{
		Client: client,
		Bucket: bucket,
	}, nil
}
