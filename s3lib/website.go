package s3lib

import (
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/service/s3"
)

func (s3client *S3Client) PutBucketWebsite(bucketName string) error {
	params := &s3.PutBucketWebsiteInput{
		Bucket: aws.String(bucketName),
		WebsiteConfiguration: &s3.WebsiteConfiguration{
			IndexDocument: &s3.IndexDocument{
				Suffix: aws.String("index.html"),
			},
			ErrorDocument: &s3.ErrorDocument{
				Key: aws.String("error.html"),
			},
		},
	}
	_, err := s3client.Client.PutBucketWebsite(params)
	if err != nil {
		return err
	}
	return nil
}

func (s3client *S3Client) GetBucketWebsite(bucketName string) error {
	params := &s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucketName),
	}
	_, err := s3client.Client.GetBucketWebsite(params)
	if err != nil {
		return err
	}
	return nil
}
func (s3client *S3Client) DeleteBucketWebsite(bucketName string) error {
	params := &s3.DeleteBucketWebsiteInput{
		Bucket: aws.String(bucketName),
	}
	_, err := s3client.Client.DeleteBucketWebsite(params)
	if err != nil {
		return err
	}
	return nil
}
