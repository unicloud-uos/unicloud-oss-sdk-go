package sample

import (
	"fmt"

	"github.com/journeymidnight/yig/test/go/lib"
	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
)

func BucketACLSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucket()

	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	err = sc.PutBucketAcl(bucketName, lib.ObjectCannedACLPublicRead)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.GetBucketAcl(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)

	fmt.Printf("BucketACLSample Run Success!\n\n")
}
