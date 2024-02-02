package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func prev_buggy_chat_gpt() {
	region := "us-west-2"

	// Load AWS configuration with the specified region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		os.Exit(1)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Specify the bucket name
	bucketName := "thisisab77-kktb4cket-name123"

	// Specify the S3 regional endpoint for us-west-2
	regionalEndpoint := "s3-" + region + ".amazonaws.com"

	// Create S3 bucket with a custom endpoint
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: s3.BucketLocationConstraint(regionalEndpoint),
		},
	})
	if err != nil {
		fmt.Println("Error creating S3 bucket:", err)
		os.Exit(1)
	}

	fmt.Println("S3 bucket created successfully:", bucketName)
}

func prev_buggy_chat_gpt_region_fix() {
	region := "us-west-2"

	// Load AWS configuration with the specified region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		os.Exit(1)
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// Specify the bucket name
	bucketName := "thisisab77-kktb4cket-name123"

	// Create S3 bucket with a custom region
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		fmt.Println("Error creating S3 bucket:", err)
		os.Exit(1)
	}

	fmt.Println("S3 bucket created successfully:", bucketName)
}
