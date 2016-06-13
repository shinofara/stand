package location

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/shinofara/stand/config"
	"os"
	"sort"
)

type S3 struct {
	Config *config.Config
	cli    *s3.S3
}

func NewS3(cfg *config.Config) *S3 {

	s3Config := cfg.StorageConfig.S3Config
	cre := credentials.NewStaticCredentials(
		s3Config.AccessKeyID,
		s3Config.SecretAccessKey,
		"")

	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(s3Config.Region),
	})

	s := &S3{
		Config: cfg,
		cli:    cli,
	}

	return s
}

func (s *S3) Save(localDir string, filename string) error {
	tmpPath := localDir + "/" + filename
	file, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Config.StorageConfig.S3Config.BucketName),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return err
	}

	if err := os.RemoveAll(tmpPath); err != nil {
		return err
	}

	return nil
}

func (s *S3) Clean() error {
	resp, _ := s.findAll()

	var keys []string
	for _, obj := range resp.Contents {
		keys = append(keys, *obj.Key)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	var num int64 = 0
	for _, key := range keys {
		if num > s.Config.StorageConfig.LifeCyrcle {
			// delete
			if err := s.delete(key); err != nil {
				return err
			}
		}

		num++
	}
	return nil
}

func (s *S3) findAll() (*s3.ListObjectsOutput, error) {
	return s.cli.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.Config.StorageConfig.S3Config.BucketName),
	})

}

func (s *S3) delete(key string) error {
	_, err := s.cli.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.Config.StorageConfig.S3Config.BucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return err
	}

	return nil
}
