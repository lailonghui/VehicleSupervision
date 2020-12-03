package cryptor

import (
	"encoding/base64"
)

// base64解码
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// base64编码
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
