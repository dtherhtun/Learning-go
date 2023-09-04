package listBuckets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	client = s3.NewFromConfig(cfg)
}

func ListBuckets() ([]*string, error) {
	buckets := make([]*string, 0, 10)

	res, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	for _, bucket := range res.Buckets {
		buckets = append(buckets, bucket.Name)
	}

	return buckets, nil
}
