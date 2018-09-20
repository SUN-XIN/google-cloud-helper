// +build go1.8

package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func NewClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

func GetPath(bucketName, fileName string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)
}
