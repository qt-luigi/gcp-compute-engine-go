package main

import (
	"context"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

// UploadFromFilename はアップロードされたファイルをCloud Storageへ保存します。
func UploadFromFilename(f multipart.File, fh *multipart.FileHeader, bucketName, filename string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	nw := client.Bucket(bucketName).Object(filename).NewWriter(ctx)
	//nw.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	nw.ContentType = fh.Header.Get("Content-Type")
	nw.CacheControl = "public, max-age=86400"
	if _, err = io.Copy(nw, f); err != nil {
		return err
	}
	if err := nw.Close(); err != nil {
		return err
	}
	return nil
}
