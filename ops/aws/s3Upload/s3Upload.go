package s3Upload

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var Client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	Client = s3.NewFromConfig(cfg)
}

func Upload(client *s3.Client, file, bucket, prefix *string) error {

	body, err := os.ReadFile(*file)
	if err != nil {
		return fmt.Errorf("file could not be opened: %w", err)
	}

	key := *prefix + path.Base(*file)

	params := &s3.PutObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	}

	if _, err := client.PutObject(context.TODO(), params); err != nil {
		return fmt.Errorf("error uploading to s3: %w", err)
	}

	return nil
}
