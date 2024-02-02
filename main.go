package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.
type BucketBasics struct {
	S3Client *s3.Client
}

// CreateBucket creates a bucket with the specified name in the specified Region.
func (basics BucketBasics) CreateBucket(name string, region string) error {
	_, err := basics.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(name),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		log.Printf("Couldn't create bucket %v in Region %v. Here's why: %v\n",
			name, region, err)
	}
	return err
}
func main() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	count := 10
	fmt.Printf("Let's list up to %v buckets for your account.\n", count)
	result, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		return
	}
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count] {
			fmt.Printf("\t%v\n", *bucket.Name)
		}
	}

	// b := BucketBasics {S3Client:s3Client,}
	// err = b.CreateBucket("test-bucket-12-jk","us-west-2")
	// if err != nil {
	// 	fmt.Printf("couldnt make bucket. Here's why: %v\n", err)
	// 	return
	// }

}
	
	// // Create S3 service client
	// svc := s3.New(sess)
	
	// // Initialize a session in us-west-2 that the SDK will use to load
	// // credentials from the shared credentials file ~/.aws/credentials.
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-west-2")},
	// )

	// // Create S3 service client
	// svc := s3.New(sess)

	// _, err = svc.CreateBucket(&s3.CreateBucketInput{
	// 	Bucket: &bucket,
	// })


	// if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &bucket}); err != nil {
	// 	log.Printf("Failed to wait for bucket to exist %s, %s\n", bucket, err)
	// 	return
	// }

	// _, err = svc.PutObject(&s3.PutObjectInput{
	// 	Body:   strings.NewReader("Hello World!"),
	// 	Bucket: &bucket,
	// 	Key:    &key,
	// })
	// if err != nil {
	// 	log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err)
	// 	return
	// }

	// log.Printf("Successfully created bucket %s and uploaded data with key %s\n", bucket, key)