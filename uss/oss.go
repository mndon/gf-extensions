package uss

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/mndon/gf-extensions/config"
)

type sOss struct {
	endpoint      string
	accessID      string
	accessKey     string
	defaultBucket string
}

var insOss *sOss

func init() {
	ctx := context.TODO()
	insOss = &sOss{
		endpoint:      config.GetValueFromConfigWithPanic(ctx, "oss.endpoint").String(),
		accessID:      config.GetValueFromConfigWithPanic(ctx, "oss.accessID").String(),
		accessKey:     config.GetValueFromConfigWithPanic(ctx, "oss.accessKey").String(),
		defaultBucket: config.GetValueFromConfigWithPanic(ctx, "oss.defaultBucket").String(),
	}

}

func Oss() *sOss {
	return insOss
}

// DeleteObj 删除文件
func (c *sOss) DeleteObj(bucketName string, objName string) error {
	if bucketName == "" {
		bucketName = c.defaultBucket
	}
	client, err := oss.New(c.endpoint, c.accessID, c.accessKey)
	if err != nil {
		return err
	}
	// New bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.DeleteObject(objName)
	if err != nil {
		return err
	}
	return nil
}

// SignUrl 签名上传路径
func (c *sOss) SignUrl(bucketName string, objName string) (url string, err error) {
	fmt.Println(c)
	if bucketName == "" {
		bucketName = c.defaultBucket
	}
	client, err := oss.New(c.endpoint, c.accessID, c.accessKey)
	if err != nil {
		return "", err
	}
	// New bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	url, err = bucket.SignURL(objName, oss.HTTPPut, 300)
	if err != nil {
		return "", err
	}
	return url, nil
}

// IsObjectExist 文件是否存在
func (c *sOss) IsObjectExist(bucketName string, objName string) (exist bool, err error) {
	fmt.Println(c)
	if bucketName == "" {
		bucketName = c.defaultBucket
	}
	client, err := oss.New(c.endpoint, c.accessID, c.accessKey)
	if err != nil {
		return false, err
	}
	// New bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return false, err
	}
	exist, err = bucket.IsObjectExist(objName)
	if err != nil {
		return false, err
	}
	return exist, nil
}
