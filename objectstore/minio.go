package objectstore

import (
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinIO() *minio.Client {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyId := os.Getenv("MINIO_ACCESSID")
	secretAccessKey := os.Getenv("MINIO_ACCESSKEY")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_SECURE"))

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
