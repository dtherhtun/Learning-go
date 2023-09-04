package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dtherhtun/Learning-go/ops/aws/s3Upload"
)

func main() {
	file := flag.String("file", "", "File to upload")
	bucket := flag.String("bucket", "", "s3 bucket")
	prefix := flag.String("prefix", "", "prefix")
	flag.Parse()

	if *file == "" || *bucket == "" || *prefix == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "Error: --file, --bucket and --prefix are required\n")
		os.Exit(1)
	}

	fmt.Printf("Uploading %s to %s/%s\n", *file, *bucket, *prefix)
	if err := s3Upload.Upload(s3Upload.Client, file, bucket, prefix); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
