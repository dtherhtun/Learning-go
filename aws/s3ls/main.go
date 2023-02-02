package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		log.Fatal(err)
	}
	s3Svc := s3.New(sess)

	results, err := s3Svc.ListBuckets(nil)
	if err != nil {
		log.Fatal("Unable to get bucket list")
	}
	fmt.Println("--- Buckets ---")
	for _, b := range results.Buckets {
		fmt.Printf("Bucket: %s \n", aws.StringValue(b.Name))
	}
}
