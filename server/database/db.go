package database

import (
	"context"
	"fmt"

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

	// endpoint := os.Getenv("MINIO_ENDPOINT")    // E.g., "localhost:9000"
	// accessKey := os.Getenv("MINIO_ACCESS_KEY") // E.g., "minioaccesskey"
	// secretKey := os.Getenv("MINIO_SECRET_KEY") // E.g., "miniosecretkey"
	// bucket := os.Getenv("MINIO_BUCKET")        // E.g., "my-bucket"

	/**
	* Test playground provided by minIO
	 */
	// endpoint := "play.min.io"
	// accessKey := "Q3AM3UQ867SPQQA43P2F"
	// secretKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	// useSSL := true
	// bucket := "testbucket"
	// MINIO_ENDPOINT="localhost:9000"
	// MINIO_PORT= 9000
	// MINIO_ACCESSKEY="AKIAIOSFODNN7EXAMPLE"
	// MINIO_SECRETKEY="wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	// MINIO_BUCKET="dev-minio"
	endpoint := "localhost:9000"
	accessKey := "AKIAIOSFODNN7EXAMPLE"
	secretKey := "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	useSSL := false
	bucket := "testbucket"

	// Initialize MinIO client
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL, // Set to true because using HTTPS
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
