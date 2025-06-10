package pkg

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

const bucketName = "test"

var MinioClient *minio.Client

func InitMinio() {
	ctx := context.Background()
	endpoint := "14.103.149.194:9000"
	accessKeyID := "minio_JnMtax"
	secretAccessKey := "minio_6NzJe2"
	useSSL := false
	var err error
	// Initialize minio client object.
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	location := "us-east-1"

	err = MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MinioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}
func Upload(object string, reader io.Reader, size int64) error {
	_, err := MinioClient.PutObject(context.Background(), bucketName, object, reader, size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
