package fxxk

import (
	"fmt"
	"sort"
	"testing"
)

/*
354. 俄罗斯套娃信封问题
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
注意：不允许旋转信封。

示例 1：
输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
示例 2：
输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1

提示：
1 <= envelopes.length <= 105
envelopes[i].length == 2
1 <= wi, hi <= 105
*/

/*
leetcode测试用例太变态了，dp已经过不了了。
*/
func maxEnvelopes2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	var res354 int
	dp354_1 := make([]int, len(envelopes))
	for i := range envelopes {
		dp354_1[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp354_1[i] = getMax(dp354_1[i], dp354_1[j]+1)
			}
		}
		res354 = getMax(res354, dp354_1[i])
	}
	return res354
}

/*
LIS的二分查找思路
*/
func maxEnvelopes1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	dp354 := make([]int, 0)
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(dp354, h); i < len(dp354) {
			dp354[i] = h
		} else {
			dp354 = append(dp354, h)
		}
	}
	return len(dp354)
}

func TestAl354(t *testing.T) {
	//fmt.Printf("%v\n", minDistance("horse", "ros"))
	fmt.Printf("%v\n", maxEnvelopes1([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	//fmt.Printf("%v\n", maxEnvelopes([][]int{{1, 1}, {1, 1}}))
	//fmt.Printf("%v\n", maxEnvelopes([][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}))
	//fmt.Printf("%v\n", maxEnvelopes([][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}))
}
