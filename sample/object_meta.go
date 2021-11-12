package sample

import (
	"fmt"
	"strings"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
	"github.com/unicloud-uos/uos-sdk-go/aws"
)

func ObjectMetaSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

	// Set Custom Meta
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	c := make(map[string]*string)
	c["a"] = aws.String("b")
	err = sc.PutObjectWithMeta(bucketName, objectKey, strings.NewReader("ObjectMetaSample"), c)
	if err != nil {
		HandleError(err)
	}
	fmt.Printf("ObjectMetaSample Run Success !\n\n")
}
