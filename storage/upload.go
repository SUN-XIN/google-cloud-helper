package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func upload(ctx context.Context,
	client *storage.Client,
	bucketName, fileName string,
	rules []storage.ACLRule,
	data []byte) error {
	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	defer wc.Close()

	//wc.ContentType = "text/plain"
	wc.ACL = rules
	if _, err := wc.Write(data); err != nil {
		return fmt.Errorf("Failed Write: %+v", err)
	}

	return nil
}

// UploadAsPublic write data to storage with the given bucket, path
// ACLRule: AllUsers RoleReader
// return the public link
func UploadAsPublic(ctx context.Context,
	client *storage.Client,
	bucketName, fileName string,
	data []byte) (error, string) {
	return upload(ctx, client, bucketName, fileName,
		[]storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}},
		data), GetPath(bucketName, fileName)
}
