package fxxk

import (
	"fmt"
	"testing"
)

/*
231. 2 的幂
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。

示例 1：
输入：n = 1 输出：true 解释：20 = 1
示例 2：
输入：n = 16 输出：true 解释：24 = 16
示例 3：
输入：n = 3 输出：false
示例 4：
输入：n = 4 输出：true
示例 5：
输入：n = 5 输出：false

提示：
-2^31 <= n <= 2^31 - 1

进阶：你能够不使用循环/递归解决此问题吗？
*/
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		//isPowerOfTwo(-n)
		return false
	}

	//n &(n-1) 消除了末尾的1.如果消除了这个1 之后为0，那就是2的幂。不然就不是
	return n&(n-1) == 0
}

func TestAl231(t *testing.T) {
	//fmt.Println(isPowerOfTwo(1))
	//fmt.Println(isPowerOfTwo(2))
	//fmt.Println(isPowerOfTwo(4))
	//fmt.Println(isPowerOfTwo(5))
	fmt.Println(isPowerOfTwo(6))
	fmt.Println(isPowerOfTwo(-16))
}
