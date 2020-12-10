package util

import (
	"bytes"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

//StructAssign把一个结构体的数据赋值给另一个结构体
//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

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
