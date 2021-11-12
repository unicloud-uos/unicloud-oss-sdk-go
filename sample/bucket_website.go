package sample

import (
	"fmt"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
)

func BucketWebsiteSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	//put bucket website
	err = sc.PutBucketWebsite(bucketName, "index.html", "error.html")
	if err != nil {
		HandleError(err)
	}

	//Get Bucket Website
	err = sc.GetBucketWebsite(bucketName)
	if err != nil {
		HandleError(err)
	}

	//Delete Bucket Website
	err = sc.DeleteBucketWebsite(bucketName)
	if err != nil {
		HandleError(err)
	}

	fmt.Printf("BucketWebsiteSample Run Success !\n\n")
}
