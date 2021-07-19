package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
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

	// TODO: Make bucket with ACL

	// Delete a bucket
	err = sc.DeleteBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	fmt.Printf("CreateBucketSample Run Success!\n\n")
}
