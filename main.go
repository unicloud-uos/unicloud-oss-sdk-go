package main

import (
	"fmt"

	"github.com/unicloud-uos/unicloud-oss-sdk-samples-go/sample"
)

func main() {
	// Read S3 config
	sample.ReadConfig()

	sample.MakeBucketSample()
	sample.BucketACLSample()
	sample.ListBucketsSample()
	sample.BucketLifecycleSample()
	sample.BucketRefererSample()
	sample.BucketLoggingSample()
	sample.BucketCORSSample()
	sample.BucketEncryptionSample()

	sample.HeadObjectSample()
	sample.PutObjectSample()
	sample.PutObjectWithForbidOverwrite()
	sample.GetObjectSample()
	sample.GetObjectByRange()
	sample.GetObjectWithCondition()
	sample.ListObjectsSample()
	sample.ListObjectVersionsSample()
	sample.DeleteObjectSample()
	sample.AppendObjectSample()
	sample.ObjectACLSample()
	sample.ObjectMetaSample()
	sample.PutEncryptObjectWithSSECSample()
	sample.PutEncryptObjectWithSSES3Sample()
	sample.MultiPartUploadSample()
	sample.MultiPartUploadSampleWithForbidOverwrite()
	sample.MultiPartDownloadSample()

	sample.CopyObjectSample()
	sample.CopyObjectWithForbidOverwriteSample()

	sample.PreSignedSample()
	sample.CleanEnv()

	fmt.Println("All samples completed !")
}
