package utils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/leonardonicola/tickethub/config"
)

func DeleteObjectFromBucket(bucket string, identifier string) error {

	svc := config.GetS3Client()
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(identifier)})
	if err != nil {
		logger.Errorf("EVENT(create) - rollback s3 insertion: %v", err)
		return fmt.Errorf("EVENT(create) - rollback s3 insertion: %v", err)
	}
	return nil
}
