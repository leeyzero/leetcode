package search

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/restore-ip-addresses/
// 题目：93. 复原 IP 地址
// 难度：中等
// 描述：给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
const SEG_COUNT = 4

func restoreIpAddresses(s string) []string {
	var segments []int
	var ans []string
	dfsRestoreIpAddresses(s, 0, segments, &ans)
	return ans
}

func dfsRestoreIpAddresses(s string, segStart int, segments []int, ans *[]string) {
	// 如果找到了4段IP地址并且遍历完了字符串，那么就是一种答案
	if len(segments) == SEG_COUNT {
		if segStart >= len(s) {
			var curr []string
			for i := 0; i < SEG_COUNT; i++ {
				curr = append(curr, strconv.Itoa(segments[i]))
			}
			*ans = append(*ans, strings.Join(curr, "."))
		}
		return
	}

	// 如果还没有找到4段IP地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}

	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		segments = append(segments, 0)
		dfsRestoreIpAddresses(s, segStart+1, segments, ans)
		segments = segments[:len(segments)-1]
	} else {
		// 一般情况，枚举每一种可能性并递归
		addr := 0
		for segEnd := segStart; segEnd < len(s); segEnd++ {
			addr = addr*10 + int(s[segEnd]-'0')
			if addr > 0 && addr <= 0xFF {
				segments = append(segments, addr)
				dfsRestoreIpAddresses(s, segEnd+1, segments, ans)
				segments = segments[:len(segments)-1]
			} else {
				break
			}
		}
	}
}
