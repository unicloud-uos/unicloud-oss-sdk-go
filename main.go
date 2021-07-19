package main

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/sample"
)

func main() {
	// Read S3 config
	sample.ReadConfig()

	//sample.MakeBucketSample()
	//sample.ListBucketsSample()
	//sample.BucketACLSample()
	//sample.BucketLifecycleSample()
	//sample.BucketRefererSample()
	//sample.BucketLoggingSample()
	//sample.BucketCORSSample()
	//sample.BucketEncryptionSample()
	//
	//sample.PutObjectSample()
	//sample.PutObjectWithForbidOverwrite()
	//sample.GetObjectSample()
	//sample.ListObjectsSample()
	//sample.DeleteObjectSample()
	//sample.AppendObjectSample()
	//sample.ObjectACLSample()
	//sample.ObjectMetaSample()
	//sample.PutEncryptObjectWithSSECSample()
	//sample.PutEncryptObjectWithSSES3Sample()
	sample.MultiPartUploadSample()
	sample.MultiPartUploadSampleWithForbidOverwrite()


	//sample.CopyObjectSample()
	//sample.CopyObjectWithForbidOverwriteSample()
	//sample.ArchiveSample()
	//sample.MySample()

	fmt.Println("All samples completed !")
}
