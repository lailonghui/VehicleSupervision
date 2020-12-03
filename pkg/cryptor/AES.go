package cryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

// AES加密
func AesEncrypt(orig string, key string) string {
	var block cipher.Block
	var k []byte

	// 转成字节数组
	origData := []byte(orig)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	if strings.EqualFold(key, "") {
		// 默认使用192bit的秘钥
		k = GetDefault192BitKey()
		block, _ = aes.NewCipher(k)
	} else {
		k = []byte(key)
		block, _ = aes.NewCipher(k)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

// AES解密
func AesDecrypt(cryted string, key string) string {
	var block cipher.Block
	var k []byte

	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	if strings.EqualFold(key, "") {
		// 默认使用192bit的秘钥 24位
		k = GetDefault192BitKey()
		block, _ = aes.NewCipher(k)
	} else {
		k = []byte(key)
		block, _ = aes.NewCipher(k)
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = pKCS7UnPadding(orig)
	return string(orig)
}

// 添加补码
// AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func pKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去除补码
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 获取128bit的秘钥
func GetDefault128BitKey() []byte {
	return []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
}

// 获取128bit的秘钥
func GetDefault192BitKey() []byte {
	return []byte{
		0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
		0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b,
	}
}

// 获取128bit的秘钥
func GetDefault256BitKey() []byte {
	return []byte{
		0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
		0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
	}
}
