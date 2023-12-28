package fxxk

import (
	"fmt"
	"testing"
)

/*
698. 划分为k个相等的子集
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

示例 1：
输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
示例 2:
输入: nums = [1,2,3,4], k = 3
输出: false

提示：
1 <= k <= len(nums) <= 16
0 < nums[i] < 10000
每个元素的频率在 [1,4] 范围内
*/

/*
和416有点相似，但是又有区别。
416是分成2组，找到一组值为和的一半就行了，剩下一组肯定就是和的另外一半了，但是这题如果照搬就会出问题。
找到和的1/k 是不够的，剩余的未必都能是1/k。所以这题得遍历全，找齐k 个 sum/k的 子集
题目翻译过来是 数组nums 能否分成 k个 和为sum/k的子集

*/
func canPartitionKSubsets(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	if len(nums) == 1 {
		//	k 肯定也等于1
		return true
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k
	//k个target值
	_ = target

	return false
}

func TestAl698(t *testing.T) {
	//fmt.Println(canPartitionKSubsets([]int{9, 5},2))
	//fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4, 5, 3},2))
	//fmt.Println(canPartition1([]int{1, 2, 3, 4, 5, 3},2))
	//fmt.Println(canPartition1([]int{14, 9, 8, 4, 3, 2},2))
	fmt.Println(canPartitionKSubsets([]int{14, 9, 8, 4, 3, 2}, 2))
	//fmt.Println(canPartitionKSubsets([]int{100, 100, 99, 97},2))
	//fmt.Println(canPartitionKSubsets([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97},2))
}