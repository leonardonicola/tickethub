package bucket

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/leonardonicola/tickethub/config"
)

var (
	s3Client *s3.S3
	logger   *config.Logger
)

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
