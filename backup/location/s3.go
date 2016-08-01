package location

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/shinofara/stand/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	cli        *s3.S3
	storageCfg *config.StorageConfig
}

func NewS3(storageCfg *config.StorageConfig) *S3 {

	s3Config := storageCfg.S3Config
	cre := credentials.NewStaticCredentials(
		s3Config.AccessKeyID,
		s3Config.SecretAccessKey,
		"")

	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(s3Config.Region),
	})

	s := &S3{
		cli:        cli,
		storageCfg: storageCfg,
	}

	return s
}

func (s *S3) Save(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.storageCfg.S3Config.BucketName),
		Key:    aws.String(s.makeUploadPath(file)),
		Body:   file,
	})
	if err != nil {
		return err
	}

	if err := os.Remove(filename); err != nil {
		return err
	}

	return nil
}

func (s *S3) Clean() error {
	resp, _ := s.findAll()

	pattern := replacePattern(s.storageCfg.Path)

	var keys []string
	for _, obj := range resp.Contents {
		if matched, _ := path.Match(pattern, *obj.Key); matched {
			keys = append(keys, *obj.Key)
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	var num int64 = 0
	for _, key := range keys {
		if num > s.storageCfg.LifeCyrcle {
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
		Bucket: aws.String(s.storageCfg.S3Config.BucketName),
	})
}

func (s *S3) delete(key string) error {
	_, err := s.cli.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.storageCfg.S3Config.BucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *S3) makeUploadPath(file *os.File) string {
	_, name := filepath.Split(file.Name())
	return fmt.Sprintf("%s/%s", s.storageCfg.Path, name)
}

func replacePattern(str string) string {
	rep := regexp.MustCompile(`^/`)
	path := fmt.Sprintf("%s",
		rep.ReplaceAllString(str, ""),
	)

	if regexp.MustCompile(`/$`).MatchString(str) {
		return fmt.Sprintf("%s*", path)
	}

	return fmt.Sprintf("%s/*", path)
}
