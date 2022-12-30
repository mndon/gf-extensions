package uss

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
)

type Uss struct {
	endpoint          string
	accessKeyId       string
	accessKeySecret   string
	Bucket            string
	BucketInternalUrl string
	BucketUrl         string
	BucketCdn         string
}

func NewUss(endpoint string, accessKeyId string, accessKeySecret string, bucket string, bucketInternalUrl string, bucketUrl string, bucketCdn string) *Uss {
	return &Uss{
		endpoint:          endpoint,
		accessKeyId:       accessKeyId,
		accessKeySecret:   accessKeySecret,
		Bucket:            bucket,
		BucketInternalUrl: bucketInternalUrl,
		BucketUrl:         bucketUrl,
		BucketCdn:         bucketCdn,
	}
}

// DeleteObj 删除文件
func (u *Uss) DeleteObj(objName string) error {
	bucket, err := u.getBucket()
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(objName)
	if err != nil {
		return err
	}
	return nil
}

// SignUrl
// @Description: 签名上传路径
// @receiver u
// @param objName
// @return url
// @return err
func (u *Uss) SignUrl(objName string) (url string, err error) {
	bucket, err := u.getBucket()
	if err != nil {
		return "", err
	}
	// 预签
	url, err = bucket.SignURL(objName, oss.HTTPPut, 300)
	if err != nil {
		return "", err
	}
	// 映射到自有域名
	if u.BucketUrl != "" {
		return strings.Replace(url, u.BucketInternalUrl, u.BucketUrl, 1), nil
	}
	return url, nil
}

// IsObjExist
// @Description: 文件是否存在
// @receiver u
// @param bucketName
// @param objName
// @return exist
// @return err
func (u *Uss) IsObjExist(objName string) (exist bool, err error) {
	bucket, err := u.getBucket()
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	exist, err = bucket.IsObjectExist(objName)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (u *Uss) GetObjUrl(objName string, cdn bool) (url string, err error) {
	exist, err := u.IsObjExist(objName)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("object not exist")
	}
	if cdn {
		return u.BucketCdn + objName, nil
	}
	return u.BucketUrl, nil
}

func (u *Uss) getBucket() (*oss.Bucket, error) {
	client, err := oss.New(u.endpoint, u.accessKeyId, u.accessKeySecret)
	if err != nil {
		return nil, err
	}
	// New bucket
	bucket, err := client.Bucket(u.Bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
