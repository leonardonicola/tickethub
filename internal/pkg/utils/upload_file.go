package utils

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	bucket "github.com/leonardonicola/tickethub/internal/pkg/s3"
)

func UploadFileToBucket(bucketName string, fileHeader *multipart.FileHeader) (*UploadedFile, error) {

	svc := bucket.GetS3Client()
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
	_, err = svc.PutObject(&s3.PutObjectInput{
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
