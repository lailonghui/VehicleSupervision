package cryptor

import "crypto/sha256"

// SHA256编码
func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return string(h.Sum(nil))
}
