package sample

func CleanEnv() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
}
