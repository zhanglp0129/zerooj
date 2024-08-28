package storagetype

import (
	"fmt"
	"testing"
	"zerooj/common"
)

func TestGetDirectContent(t *testing.T) {
	content, err := GetContent(DirectStorage, []byte("hello, world"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func TestGetObjectContent(t *testing.T) {
	c := common.MinioConfig{
		Endpoint:        "127.0.0.1:9000",
		AccessKeyID:     "M7DXjku6iWHzWE9pAcWw",
		SecretAccessKey: "bPGm73a9eqxRI3Sqn42AgTm4OftqPskYpnnEQ0eY",
		BucketName:      "zerooj",
	}

	mc, err := common.InitMinio(c)
	if err != nil {
		panic(err)
	}
	content, err := GetContent(ObjectStorage, []byte("hello.txt"), WithObject(mc, c.BucketName))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func TestGetCloudContent(t *testing.T) {
	content, err := GetContent(CloudStorage, []byte("https://www.baidu.com"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
