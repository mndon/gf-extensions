package http

import (
	"crypto/md5"
	"encoding/hex"
)

type utils struct {
}

var insUtils = utils{}

func Utils() *utils {
	return &insUtils
}

// Md5Ency md5加密
func (u utils) md5Ency(data string, salt string) string {
	h := md5.New()
	h.Write([]byte(data + salt))
	result := hex.EncodeToString(h.Sum(nil))
	return result
}
