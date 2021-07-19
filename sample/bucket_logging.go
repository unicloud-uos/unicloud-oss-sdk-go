package sample

import "fmt"

func BucketLoggingSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	// TODO: SetBucketLogging(bucketName, logBucketName, "prefix")

	// TODO: GetBucketLogging(bucketName)

	// TODO: DeleteBucketLogging(bucketName)

	fmt.Printf("BucketLoggingSample Run Success !\n\n")
}
