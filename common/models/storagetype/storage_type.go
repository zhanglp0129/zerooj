package storagetype

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"net/http"
	"zerooj/common/constant"
)

// StorageType 存储类型
type StorageType uint32

const (
	// DirectStorage 直接存储，保存的是内容
	DirectStorage StorageType = iota + 1
	// ObjectStorage 对象存储，保存的是对象名称
	ObjectStorage
	// CloudStorage 云存储，保存的是直链URL
	CloudStorage
)

// GetContent 获取实际存储内容。推荐直接存储和云存储使用
func GetContent(st StorageType, value []byte, options ...Option) ([]byte, error) {
	cfg := loadConfig(options...)

	// 判断存储类型
	switch st {
	case DirectStorage:
		return value, nil
	case ObjectStorage:
		reader, err := readObjectStorage(cfg, string(value))
		if err != nil {
			return nil, err
		}
		return io.ReadAll(reader)
	case CloudStorage:
		reader, err := readCloudStorage(cfg, string(value))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return io.ReadAll(reader)
	default:
		return nil, constant.NotSupportStorageType
	}

}

// GetContentReader 获取实际存储内容，返回reader。推荐直接存储和对象存储使用
func GetContentReader(st StorageType, value []byte, options ...Option) (io.Reader, error) {
	cfg := loadConfig(options...)

	// 判断存储类型
	switch st {
	case DirectStorage:
		return bytes.NewBuffer(value), nil
	case ObjectStorage:
		return readObjectStorage(cfg, string(value))
	case CloudStorage:
		reader, err := readCloudStorage(cfg, string(value))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		contentBytes, err := io.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(contentBytes), nil

	default:
		return nil, constant.NotSupportStorageType
	}
}

// Storage 存储内容。如果小于等于255字节，直接存储；大于255字节，对象存储，并返回对象名
func Storage(content []byte, mc *minio.Client, bucketName string) ([]byte, StorageType, error) {
	if len(content) <= 255 {
		return content, DirectStorage, nil
	}

	objectName := uuid.New().String()
	_, err := mc.PutObject(context.Background(), bucketName, objectName, bytes.NewBuffer(content), int64(len(content)), minio.PutObjectOptions{})
	if err != nil {
		return nil, 0, err
	}

	return []byte(objectName), ObjectStorage, nil
}

// StorageObject 存储对象，并返回对象名
func StorageObject(content io.Reader, mc *minio.Client, bucketName string) (string, StorageType, error) {
	objectName := uuid.New().String()
	_, err := mc.PutObject(context.Background(), bucketName, objectName, content, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", 0, err
	}
	return objectName, ObjectStorage, nil
}

func readObjectStorage(c *config, objectName string) (io.Reader, error) {
	obj, err := c.mc.GetObject(context.Background(), c.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func readCloudStorage(c *config, url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// Option 选项
type Option func(c *config)

// WithObject 使用对象存储时，调用该函数
func WithObject(mc *minio.Client, bucket string) Option {
	return func(c *config) {
		c.mc = mc
		c.bucket = bucket
	}
}

type config struct {
	mc     *minio.Client
	bucket string
}

func loadConfig(options ...Option) *config {
	// 加载配置
	var cfg config
	for _, opt := range options {
		opt(&cfg)
	}
	return &cfg
}
