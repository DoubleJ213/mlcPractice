package fxxk

import (
	"fmt"
	"testing"
)

/*
438. 找到字符串中所有字母异位词
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 104
s和p仅包含小写字母
*/

func findAnagrams(s string, p string) []int {
	//s 包含 p
	if len(s) < len(p) {
		return nil
	}
	result := make([]int, 0)
	count := make([]int, 26)
	for i := 0; i < len(p); i++ {
		//p[i]byte range 遍历 rune
		count[p[i]-'a']++
	}
	for j := 0; j < len(p); j++ {
		count[s[j]-'a']--
	}
	left := 0
	if countIsNull(count) {
		result = append(result, left)
	}
	for k := len(p); k < len(s); k++ {
		count[s[left]-'a']++
		left++
		count[s[k]-'a']--
		if countIsNull(count) {
			result = append(result, left)
		}
	}
	return result
}

func countIsNull(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func TestAl438(t *testing.T) {
	//fmt.Printf("%v", findAnagrams("abab", "ab"))
	fmt.Printf("%v", findAnagrams("cbaebabacd", "abc"))
}
