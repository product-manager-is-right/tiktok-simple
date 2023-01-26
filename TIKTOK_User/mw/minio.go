package mw

import (
	"GoProject/configs"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

func UploadFile(BucketName string, objectName string, reader io.Reader, objectsize int64, fileType string) error {
	ctx := context.Background()
	// Minio 对象存储初始化
	minioClient, err := minio.New("120.25.2.146:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(configs.AccessKeyId, configs.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Printf("mistaken in store info in minio")
	}
	n, err := minioClient.PutObject(ctx, BucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: fileType,
	})
	if err != nil {
		log.Printf("upload %s of size %d failed, %s", BucketName, objectsize, err)
		return err
	}
	log.Printf("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}
