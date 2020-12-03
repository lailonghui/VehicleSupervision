package util

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
)

// 验证图片文件格式
func ValidImageSuffix(suffix string) (ok bool) {
	switch strings.ToLower(suffix) {
	case "svg", "jpg", "jpeg", "png":
		ok = true
	}
	return
}

// 字符串拼接
func StringSplit(data ...string) (str string, err error) {
	var buffer bytes.Buffer
	loop := len(data)
	for i := 0; i < loop; i++ {
		buffer.WriteString(data[i])
	}
	str = buffer.String()
	return
}

// 判断是否包含中文
func IsContainChinese(str string) (ok bool) {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
