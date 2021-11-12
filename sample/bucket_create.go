package sample

import (
	"fmt"
	"github.com/journeymidnight/yig/test/go/lib"
	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
)

func MakeBucketSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Delete a bucket
	err = sc.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Make bucket with ACL
	err = sc.MakeBucketWithAcl(bucketName, lib.BucketCannedACLPublicRead)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketAcl(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)

	err = sc.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Printf("CreateBucketSample Run Success!\n\n")
}
