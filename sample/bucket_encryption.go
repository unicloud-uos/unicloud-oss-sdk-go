package sample

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
	"github.com/unicloud-uos/uos-sdk-go/aws"
	"github.com/unicloud-uos/uos-sdk-go/service/s3"
)

func BucketEncryptionSample() {
	DeleteTestBucket()
	defer DeleteTestBucket()

	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	encryption := &s3.ServerSideEncryptionConfiguration{
		Rules: []*s3.ServerSideEncryptionRule{
			{
				ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
					SSEAlgorithm: aws.String("AES256"),
				},
			},
		},
	}

	err = sc.PutBucketEncryption(bucketName, encryption)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketEncryption(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket Encryption", out)

	out, err = sc.DeleteBucketEncryption(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Delete Bucket Encryption", out)

	fmt.Printf("BucketLifecycleSample Run Success !\n\n")
}
