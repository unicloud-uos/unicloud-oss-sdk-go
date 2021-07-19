package sample

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	endpoint      string
	accessKey     string
	secretKey     string
	bucketName    string
	objectKey     string
	localFilePath string
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("sample")
	viper.ReadInConfig()
	endpoint = viper.GetString("s3.endpoint")
	accessKey = viper.GetString("s3.accessKey")
	secretKey = viper.GetString("s3.secretKey")
	bucketName = viper.GetString("s3.bucketName")
	objectKey = viper.GetString("s3.objectKey")
	localFilePath = viper.GetString("s3.localFilePath")
	fmt.Printf("Read Config:\n"+
		"endpoint: %s\n"+
		"accessKey: %s\n"+
		"secretKey: %s\n"+
		"bucketName: %s\n"+
		"objectKey: %s\n"+
		"localFilePath: %s\n", endpoint, accessKey, secretKey, bucketName, objectKey, localFilePath)
}
