package bucket

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/leonardonicola/tickethub/config"
)

var (
	s3Client *s3.S3
	logger   *config.Logger
)

type UploadedFile struct {
	Identifier string
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func InitS3Client() {
	logger = config.NewLogger()
	access_key, found := os.LookupEnv("AWS_ACCESS_KEY")
	if !found {
		logger.Fatalf("AWS S3 - access key not found")
		exitErrorf("AWS S3 - access key not found")
	}
	secret_key, found := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !found {
		logger.Fatalf("AWS S3 - secret access key not found")
		exitErrorf("AWS S3 - secret access key not found")
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String("sa-east-1"),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, "")},
	)

	if err != nil {
		logger.Errorf("AWS S3 - error occurred: %v", err)
		exitErrorf("AWS S3 - error occurred: %v", err)
	}
	logger.Debug("Created new S3 client")
	s3Client = s3.New(s3Session)

}

func GetS3Client() *s3.S3 {
	return s3Client
}

func UploadFileToBucket(bucketName string, fileHeader *multipart.FileHeader) (*UploadedFile, error) {

	file, err := fileHeader.Open()

	if err != nil {
		logger.Errorf("Error on file upload: %v", err)
		return nil, fmt.Errorf("file upload error: %v", err)
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)

	// Read the entire file at one time (prolly change to streams?)
	if _, err := file.Read(buffer); err != nil {
		logger.Errorf("Error on file reading: %v", err)
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	randomFilename := uuid.NewString()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(randomFilename),
		Body:        fileBytes,
		ContentType: aws.String(fileType),
	})

	if err != nil {
		logger.Errorf("S3 put object error: %v", err)
		return nil, fmt.Errorf("EVENT(create) S3: %v", err)
	}

	logger.Info("S3 put object had success")

	return &UploadedFile{
		Identifier: randomFilename,
	}, nil

}
func DeleteObjectFromBucket(bucketName, identifier string) error {

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: aws.String(identifier)})
	if err != nil {
		logger.Errorf("EVENT(create) - rollback s3 insertion: %v", err)
		return fmt.Errorf("EVENT(create) - rollback s3 insertion: %v", err)
	}
	return nil
}
