package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
)

func BucketACLSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucket()

	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// TODO : Surport Set Bucket CannedACL 'PublicRead'

	out, err := sc.GetBucketAcl(bucketName)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)

	fmt.Printf("BucketACLSample Run Success!\n\n")
}
