package location

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/shinofara/stand/config"
	"os"
)

type S3 struct {
	Config *config.Config
}

func (s *S3) Save(filename string) error {
	file, err := os.Open(fmt.Sprintf("/tmp/%s", filename))
	if err != nil {
		return err
	}
	defer file.Close()

	s3Config := s.Config.S3Config
	cre := credentials.NewStaticCredentials(
		s3Config.AccessKeyID,
		s3Config.SecretAccessKey,
		"")

	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(s3Config.Region),
	})

	_, err = cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Config.BucketName),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *S3) Clean() error {
	return nil
}
