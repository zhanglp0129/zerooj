package storagetype

import (
	"bytes"
	"context"
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

// GetContent 获取实际存储内容
func GetContent(st StorageType, value []byte, options ...Option) ([]byte, error) {
	cfg := loadConfig(options...)

	// 判断存储类型
	switch st {
	case DirectStorage:
		return value, nil
	case ObjectStorage:
		reader, err := useObjectStorage(cfg, string(value))
		if err != nil {
			return nil, err
		}
		return io.ReadAll(reader)
	case CloudStorage:
		reader, err := useCloudStorage(cfg, string(value))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return io.ReadAll(reader)
	default:
		return nil, constant.NotSupportStorageType
	}

}

// GetContentReader 获取实际存储内容，返回reader
func GetContentReader(st StorageType, value []byte, options ...Option) (io.Reader, error) {
	cfg := loadConfig(options...)

	// 判断存储类型
	switch st {
	case DirectStorage:
		return bytes.NewBuffer(value), nil
	case ObjectStorage:
		return useObjectStorage(cfg, string(value))
	case CloudStorage:
		reader, err := useCloudStorage(cfg, string(value))
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

func useObjectStorage(c *config, objectName string) (io.Reader, error) {
	obj, err := c.mc.GetObject(context.Background(), c.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func useCloudStorage(c *config, url string) (io.ReadCloser, error) {
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
