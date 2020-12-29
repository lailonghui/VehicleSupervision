package util

import (
	"fmt"
	"testing"
)

func TestSliceSplitInt64(t *testing.T) {
	var arr []int64
	arr = append(arr, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	arrs := SliceSplitInt64(arr, 10)
	fmt.Println(arrs)
}
