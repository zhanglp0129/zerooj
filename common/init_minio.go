package common

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool `json:",default=false"`
	BucketName      string
	Location        string `json:",optional"`
}

// InitMinio 初始化minio
func InitMinio(c MinioConfig) (*minio.Client, error) {
	// 创建客户端
	minioClient, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
		Secure: c.UseSSL,
		Region: c.Location,
	})
	if err != nil {
		return nil, err
	}

	// 创建存储桶
	exists, err := minioClient.BucketExists(context.Background(), c.BucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), c.BucketName, minio.MakeBucketOptions{
			Region: c.Location,
		})
		if err != nil {
			return nil, err
		}

		// 设置可匿名读
		prefix := "readonly/*"
		err = setReadonly(minioClient, c.BucketName, prefix)
		if err != nil {
			return nil, err
		}
	}

	return minioClient, nil
}

func setReadonly(mc *minio.Client, bucketName, prefix string) error {
	policy := fmt.Sprintf(`{"Version": "2012-10-17","Statement": [{"Effect": "Allow","Principal": {"AWS": ["*"]},"Action": ["s3:GetObject"],"Resource": ["arn:aws:s3:::%s/%s"]}]}`, bucketName, prefix)
	err := mc.SetBucketPolicy(context.Background(), bucketName, policy)

	return err
}
