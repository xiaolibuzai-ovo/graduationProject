package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type CloudStorageClient interface {
	Upload(file *multipart.File, filename string) error
}

type cloudStorageClient struct {
	client *storage.Client
	bucket string
}

func NewCloudStorageClient(bucket string) (CloudStorageClient, error) {
	var client *storage.Client
	key := os.Getenv("GCP_SERVICE_ACCOUNT_KEY")
	if len(key) <= 0 {
		return nil, fmt.Errorf("GCP_SERVICE_ACCOUNT_KEY is not set")
	}
	var err error
	bytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	client, err = storage.NewClient(context.Background(), option.WithCredentialsJSON(bytes))
	if err != nil {
		return nil, err
	}
	return &cloudStorageClient{
		client: client,
		bucket: bucket,
	}, nil
}

func (c *cloudStorageClient) Upload(file *multipart.File, filename string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()

	bucket := c.client.Bucket(c.bucket)
	sw := bucket.Object(filename).NewWriter(ctx)
	if _, err := io.Copy(sw, *file); err != nil {
		return err
	}
	if err := sw.Close(); err != nil {
		return err
	}
	return nil
}
