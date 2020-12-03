package cryptor

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5编码
func MD5Encode(data string) (sign string) {
	h := md5.New() // 初始化HASH对象
	h.Write([]byte(data))
	sign = hex.EncodeToString(h.Sum(nil))
	return
}
