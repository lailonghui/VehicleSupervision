package util

import (
	"fmt"
	"math"
)

// 切片分割
func SliceSplitInt64(arr []int64, size int) (arrs [][]int64) {
	aLen := len(arr)
	if aLen < size {
		return [][]int64{arr}
	}
	arrsLen := int(math.Ceil(float64(aLen) / float64(size)))
	fmt.Println(arrsLen)
	for i := 0; i < arrsLen; i++ {
		if i == arrsLen-1 {
			arrs = append(arrs, arr[i*size:aLen])
		} else {
			arrs = append(arrs, arr[i*size:(i+1)*size])
		}
	}
	return arrs
}

func SliceSplitString(arr []string, size int) (arrs [][]string) {
	aLen := len(arr)
	if aLen < size {
		return [][]string{arr}
	}
	arrsLen := int(math.Ceil(float64(aLen) / float64(size)))
	fmt.Println(arrsLen)
	for i := 0; i < arrsLen; i++ {
		if i == arrsLen-1 {
			arrs = append(arrs, arr[i*size:aLen])
		} else {
			arrs = append(arrs, arr[i*size:(i+1)*size])
		}
	}
	return arrs
}
