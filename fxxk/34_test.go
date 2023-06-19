package fxxk

import (
	"fmt"
	"testing"
)

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
你必须设计并实现时间复杂度为O(log n)的算法解决此问题。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]

提示：
0 <= nums.length <= 105
-109<= nums[i]<= 109
nums是一个非递减数组
-109<= target<= 109
*/

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return res
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			//右边界收缩 找左边界
			right = mid - 1
			//left = mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left != len(nums) && nums[left] == target {
		res[0] = left
	}

	left1, right1 := 0, len(nums)-1
	for left1 <= right1 {
		mid := (right1-left1)/2 + left1
		if nums[mid] == target {
			//左边界收缩 找右边界
			//right = mid - 1
			left1 = mid + 1
		} else if nums[mid] > target {
			right1 = mid - 1
		} else if nums[mid] < target {
			left1 = mid + 1
		}
	}
	if right1 >= 0 && nums[right1] == target {
		res[1] = right1
	}

	return res
}

func TestAl34(t *testing.T) {
	fmt.Printf("%v", searchRange([]int{5, 6, 7, 8, 8, 10}, 8))
	//fmt.Printf("%v", searchRange([]int{1}, 0))
	//fmt.Printf("%v", searchRange([]int{1}, 1))
}
