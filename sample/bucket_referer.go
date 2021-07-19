package sample

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
)

func BucketRefererSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	var referers = []string{
		"http://www.unicloud.com",
	}

	err = sc.SetReferer(bucketName, referers)
	if err != nil {
		HandleError(err)
	}

	r, err := sc.GetReferer(bucketName)
	fmt.Println(r)

	err = sc.SetReferer(bucketName, []string{})
	if err != nil {
		HandleError(err)
	}

	r2, err := sc.GetReferer(bucketName)
	fmt.Println(r2)

	fmt.Printf("BucketRefererSample Run Success !\n\n")
}
