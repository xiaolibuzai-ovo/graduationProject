package config

import (
	"os"
)

var gcsClient *GcsClient

type GcsClient struct {
	GcsClient        CloudStorageClient
	StoragePublicUrl string
}

func InitGcs() {
	gcsClient = NewGcs()
}

func GetGcsClient() *GcsClient {
	return gcsClient
}

func NewGcs() *GcsClient {
	storagePublicUrl := os.Getenv("STORAGE_PUBLIC_URL")
	if len(storagePublicUrl) <= 0 {
		panic("STORAGE_PUBLIC_URL not set")
	}
	storagePublicBucket := os.Getenv("STORAGE_PUBLIC_BUCKET")
	if len(storagePublicBucket) <= 0 {
		panic("STORAGE_PUBLIC_BUCKET not set")
	}
	publicStorageClient, err := NewCloudStorageClient(storagePublicBucket)
	if err != nil {
		panic(err)
	}
	return &GcsClient{
		GcsClient:        publicStorageClient,
		StoragePublicUrl: storagePublicUrl,
	}
}
