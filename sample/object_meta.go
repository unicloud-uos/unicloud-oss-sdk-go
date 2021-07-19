package sample

import "fmt"

func ObjectMetaSample() {
	DeleteTestBucketAndObject()

	defer DeleteTestBucketAndObject()

	// TODO: Set Custom Meta
	fmt.Printf("ObjectMetaSample Run Success !\n\n")
}
