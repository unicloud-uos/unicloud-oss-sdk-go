package sample

import (
	"fmt"
	"strings"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
	"github.com/unicloud-uos/uos-sdk-go/service/s3"
)

func ObjectACLSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Test ObjectACL public-read
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("ObjectACLSample"))
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObjectAcl(bucketName, objectKey, s3.BucketCannedACLPublicRead)
	if err != nil {
		HandleError(err)
	}
	out, err := sc.GetObjectAcl(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)
	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// Test ObjectACL public-read-write
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("ObjectACLSample"))
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObjectAcl(bucketName, objectKey, s3.BucketCannedACLPublicReadWrite)
	if err != nil {
		HandleError(err)
	}
	out, err = sc.GetObjectAcl(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Get Bucket ACL:", out)
	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	fmt.Printf("ObjectACLSample Run Success!\n\n")
}
