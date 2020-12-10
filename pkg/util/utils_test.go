package util

import (
	"fmt"
	"testing"
)

type A struct {
	AA int
	BB int
	CC int
	DD int
}
type AInput struct {
	AA int
	BB int
	CC int
	DD int
}

func TestStructAssign(t *testing.T) {
	a := &A{AA: 1, BB: 2}
	b := &AInput{CC: 3, DD: 4}
	StructAssign(a, b)
	fmt.Println(a)
	// Output: {AA: 1, BB: 2, CC: 3, DD: 4}
}
