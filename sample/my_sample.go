package sample

func MySample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

}
