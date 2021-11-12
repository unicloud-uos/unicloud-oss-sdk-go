package sample

import (
	"fmt"
	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
	"strings"
)

func DeleteObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	// Delete single key
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// 1. Delete an object
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("DeleteObjectSample"))
	if err != nil {
		HandleError(err)
	}

	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// Delete file that not exists will not failed
	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// 2. Delete multiple objects
	err = sc.PutObject(bucketName, objectKey+"1", strings.NewReader("DeleteObjectSample"))
	if err != nil {
		HandleError(err)
	}

	err = sc.PutObject(bucketName, objectKey+"2", strings.NewReader("DeleteObjectSample"))
	if err != nil {
		HandleError(err)
	}

	keys, err := sc.DeleteObjects(bucketName, objectKey+"1", objectKey+"2")
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Delete keys: ", keys)

	fmt.Printf("DeleteObjectSample Run Success !\n\n")
}
