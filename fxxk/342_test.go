package fxxk

import (
	"fmt"
	"testing"
)

//给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
//整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x
//示例 1： 输入：n = 16 输出：true
//示例 2：输入：n = 5 输出：false
//示例 3：输入：n = 1 输出：true

//-2^31 <= n <= 2^31 - 1

func isPowerOfFour(n int) bool {
	//math.MaxInt32 math.MinInt32
	//int8  类型大小为 1 字节 int16 类型大小为 2 字节 int32 类型大小为 4 字节 int64 类型大小为 8 字节
	//int 位数是和操作系统相关的，32位操作系统，int是4字节，64位是8字节

	//if n < 0 {
	//	isPowerOfFour(-n)
	//}
	//4*-4 不是4的幂，不要神经了，和算阶乘的不一样，那个是底数为负数，简化成全正的操作

	//傻算
	temp := 1
	for temp < n {
		temp *= 4
	}
	return temp == n
}

//这个题和“2的幂”“3的幂”一样，有不用循环和递归就能直接判断的方法，同样十分巧妙，属于二进制/位运算的应用。
//4的幂的数，都是这样的数：100、10000、1000000……（4、16、64）
//观察规律，可以发现4的幂需要满足以下条件：
//1.最高位是1，其余都为0；
//2.最高位的1应该在奇数位上，比如：100的1在第三位上；
//那么对应的判断方法为：
//用 num & (num-1) 可以判断条件1，比如：100(4) & 011(3) == 0，结果为 0 说明符合条件1；
//是否在奇数位可以用 0xaaaaaaaa 判断，16 进制的 a 是 1010，比如：0100(4) & 1010(a) == 0，结果为 0 说明最高位 1 在奇数位上；

func isPowerOfFour1(n int) bool {
	//前提 -2^31 <= n <= 2^31 - 1
	//其实只要看正数部分，2^31-1
	//2^2-1=3 二进制表示0100-0001=0011
	//2^31-1 最大值 {31个1}
	//想通过位运算，判断是否在奇数位上为1，0100(4) & 1010(a)==0
	//那只需要八个a就可以与上最大值
	//return n > 3 && n&(n-1) == 0 && 0xaaaaaaaa&n == 0
	//n=1是4的0次幂，结果是true的
	return n > 0 && n&(n-1) == 0 && 0xaaaaaaaa&n == 0
}

func isPowerOfFour2(n int) bool {
	// 不能像3的幂那样傻算了，2也能被整除，但是不是4的幂 所以取余（取模）是有缺陷的
	return n > 0 && 4294967296%n == 0
}

func TestAl342(t *testing.T) {
	fmt.Print(isPowerOfFour(5))
	fmt.Print(isPowerOfFour1(5))
	fmt.Print(isPowerOfFour2(5))
}
