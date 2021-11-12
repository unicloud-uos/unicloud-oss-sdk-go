package sample

import (
	"fmt"
	"os"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
)

func HeadObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Put a file
	f, err := os.Open(localFilePath)
	defer f.Close()
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObject(bucketName, objectKey, f)
	if err != nil {
		HandleError(err)
	}

	out, err := sc.HeadObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("HeadObject result: ", out)

	fmt.Printf("HeadObjectSample Run Success !\n\n")
}
