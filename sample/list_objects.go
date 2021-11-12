package sample

import (
	"fmt"
	"strings"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/s3lib"
	"github.com/unicloud-uos/uos-sdk-go/service/s3"
)

func ListObjectsSample() {
	var keys = []string{
		objectKey + "1",
		objectKey + "2",
		objectKey + "3",
		objectKey + "4",
		objectKey + "/1-1",
		objectKey + "/1-2",
		objectKey + "/2-1",
		objectKey + "/2-2",
	}
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	for _, k := range keys {
		sc.DeleteObject(bucketName, k)
	}
	DeleteTestBucketAndObject()

	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	for _, k := range keys {
		err := sc.PutObject(bucketName, k, strings.NewReader(k))
		if err != nil {
			HandleError(err)
		}
	}

	_, _, _, err = sc.ListObjects(bucketName, objectKey+"/", "/", 1000)
	if err != nil {
		HandleError(err)
	}

	for _, k := range keys {
		err := sc.DeleteObject(bucketName, k)
		if err != nil {
			HandleError(err)
		}
	}
	DeleteTestBucketAndObject()
	fmt.Printf("ListObjectsSample Run Success !\n\n")
}

func ListObjectVersionsSample() {
	var keys = []string{
		objectKey + "1",
		objectKey + "2",
		objectKey + "3",
		objectKey + "4",
		objectKey + "/1-1",
		objectKey + "/1-2",
		objectKey + "/2-1",
		objectKey + "/2-2",
	}
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	defer func() {
		sc.DeleteObjectVersion(bucketName, objectKey, "null")
		out, _ := sc.ListObjectVersions(bucketName, "", "", "", 1000)
		for _, v := range out.CommonPrefixes {
			fmt.Println("Prefix:", v)
		}
		for _, v := range out.DeleteMarkers {
			fmt.Println("DeleteMarkers:", v)
		}
		for _, v := range out.Versions {
			sc.DeleteObjectVersion(bucketName, objectKey, *v.VersionId)
		}
		sc.DeleteBucket(bucketName)
	}()

	for _, k := range keys {
		sc.DeleteObjectVersion(bucketName, k, "nill")
	}
	DeleteTestBucketAndObject()

	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	for _, k := range keys {
		err := sc.PutObject(bucketName, k, strings.NewReader(k))
		if err != nil {
			HandleError(err)
		}
	}
	objects, _, _, err := sc.ListObjects(bucketName, "", "/", 1000)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("ListObjects:", objects)

	err = sc.PutBucketVersion(bucketName, s3.BucketVersioningStatusEnabled)
	if err != nil {
		HandleError(err)
	}

	for _, k := range keys {
		err := sc.PutObject(bucketName, k, strings.NewReader(k))
		if err != nil {
			HandleError(err)
		}
	}

	for _, k := range keys {
		err := sc.DeleteObject(bucketName, k)
		if err != nil {
			HandleError(err)
		}
	}

	out, err := sc.ListObjectVersions(bucketName, objectKey+"/", "/", "", 1000)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Prefix: ", *out.Prefix, "\nListObjects: ", out.Versions, "\nDeleteMarker: ", out.DeleteMarkers)
	for _, k := range out.Versions {
		_, err := sc.DeleteObjectVersion(bucketName, *k.Key, *k.VersionId)
		if err != nil {
			HandleError(err)
		}
	}
	for _, k := range out.DeleteMarkers {
		_, err := sc.DeleteObjectVersion(bucketName, *k.Key, *k.VersionId)
		if err != nil {
			HandleError(err)
		}
	}
	out, err = sc.ListObjectVersions(bucketName, "", "/", "", 1000)
	if err != nil {
		HandleError(err)
	}
	fmt.Println("Prefix: ", *out.Prefix, "\nListObjects: ", out.Versions, "\nDeleteMarker: ", out.DeleteMarkers)
	for _, k := range out.Versions {
		_, err := sc.DeleteObjectVersion(bucketName, *k.Key, *k.VersionId)
		if err != nil {
			HandleError(err)
		}
	}
	for _, k := range out.DeleteMarkers {
		_, err := sc.DeleteObjectVersion(bucketName, *k.Key, *k.VersionId)
		if err != nil {
			HandleError(err)
		}
	}
	sc.DeleteBucket(bucketName)
	fmt.Printf("ListObjectVersionsSample Run Success !\n\n")
}
