package fxxk

import (
	"fmt"
	"testing"
)

/*
66. 加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：
输入：digits = [0]
输出：[1]

提示：
1 <= digits.length <= 100
0 <= digits[i] <= 9
*/

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			break
		}
		for digits[i] == 9 {
			//进位
			digits[i] = 0
			if i == 0 {
				//前面没有元素了，得重新构造个数组
				tmp := make([]int, len(digits)+1)
				tmp[0] = digits[i] + 1
				copy(tmp[1:], digits)
				digits = tmp
			}
		}
	}
	return digits
}

func TestAl66(t *testing.T) {
	//fmt.Printf("%v", plusOne([]int{1, 9}))
	fmt.Printf("%v", plusOne([]int{1, 2, 3}))
	//fmt.Printf("%v", plusOne([]int{9, 9}))
}
