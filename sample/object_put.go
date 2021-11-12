package sample

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
)

func PutObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// 1. Put a string object
	err = sc.PutObject(bucketName, objectKey, strings.NewReader("PutObjectSample"))
	if err != nil {
		HandleError(err)
	}

	err = sc.DeleteObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// 2. Put a file
	f, err := os.Open(localFilePath)
	defer f.Close()
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObject(bucketName, objectKey, f)
	if err != nil {
		HandleError(err)
	}
	out, err := sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	out.Close()

	fmt.Printf("NewObjectSample Run Success !\n\n")
}

func PutObjectWithForbidOverwrite() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	err = sc.PutObject(bucketName, objectKey, strings.NewReader("PutObjectWithForbidOverwrite"))
	if err != nil {
		HandleError(err)
	}
	out, err := sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	b, _ := ioutil.ReadAll(out)
	fmt.Println(string(b))

	//set forbid overwrite
	_, err = sc.PutObjectWithForbidOverwrite(bucketName, objectKey, strings.NewReader("OverwriteValue"), true)
	if err == nil {
		HandleError(errors.New("should be error"))
	}

	out, err = sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	b, _ = ioutil.ReadAll(out)
	fmt.Println(string(b))

	_, err = sc.PutObjectWithForbidOverwrite(bucketName, objectKey, strings.NewReader("OverwriteValue"), false)
	if err != nil {
		HandleError(err)
	}
	out, err = sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}
	b, _ = ioutil.ReadAll(out)
	fmt.Println(string(b))
	out.Close()

	fmt.Printf("NewObjectSample Run Success !\n\n")
}
