package sort

import (
	"bytes"

	"github.com/leeyzero/leetcode/algorithm/base"
)

// https://leetcode-cn.com/problems/sort-characters-by-frequency/
// 题目：451. 根据字符出现频率排序
// 难度：中等
// 思路：桶排序
func frequencySort(s string) string {
	// 先建字符到频次的映射关系
	hash := map[byte]int{}
	maxCount := 0
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
		maxCount = base.Max(maxCount, hash[s[i]])
	}

	// 以每个频次做为桶，将相同频次的字符放到桶中
	buckets := make([][]byte, maxCount+1)
	for k, v := range hash {
		buckets[v] = append(buckets[v], k)
	}

	// 从最大频次的桶开始遍历
	ans := []byte{}
	for i := len(buckets) - 1; i > 0; i-- {
		for _, v := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{v}, i)...)
		}
	}
	return string(ans)
}
