package fxxk

import (
	"fmt"
	"testing"
)

func reverse(x int) int {
	y := 0
	for x != 0 {
		//	B % A 求余
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

func TestReverseInt(t *testing.T) {
	fmt.Printf("%v", reverse(123))
}
