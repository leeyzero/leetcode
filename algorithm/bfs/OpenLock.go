package bfs

/*
https://leetcode.cn/problems/open-the-lock/
题目：752. 打开转盘锁
难度：中等

你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

示例 1:
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。

示例 2:
输入: deadends = ["8888"], target = "0009"
输出：1
解释：把最后一位反向旋转一次即可 "0000" -> "0009"。

示例 3:
输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：无法旋转到目标数字且不被锁定。


提示：
1 <= deadends.length <= 500
deadends[i].length == 4
target.length == 4
target 不在 deadends 之中
target 和 deadends[i] 仅由若干位数字组成
*/

func openLock(deadends []string, target string) int {
	// 将deadends存入字典
	dict := map[string]bool{}
	for _, v := range deadends {
		dict[v] = true
	}

	start := "0000"
	queue := []string{start}
	visited := map[string]bool{start: true}
	steps := 0
	for len(queue) > 0 {
		for i, n := 0, len(queue); i < n; i++ {
			// 取出队头节点
			curr := queue[0]
			queue = queue[1:]

			// 在死亡字典中，忽略
			if _, ok := dict[curr]; ok {
				continue
			}
			// 找到目标数字
			if curr == target {
				return steps
			}

			// 找出curr的所有邻节点
			for _, adj := range getAdajcent(curr) {
				if _, ok := visited[adj]; ok {
					// 忽略访问过的节点
					continue
				}
				queue = append(queue, adj)
				visited[adj] = true
			}
		}
		steps += 1
	}

	// 没找到
	return -1
}

func getAdajcent(nums string) []string {
	var ans []string
	for i := 0; i < 4; i++ {
		ans = append(ans, turnUp(nums, i))
		ans = append(ans, turnDown(nums, i))
	}
	return ans
}

func turnUp(nums string, pos int) string {
	byteNums := []byte(nums)
	if byteNums[pos] == '9' {
		byteNums[pos] = '0'
	} else {
		byteNums[pos] += 1
	}
	return string(byteNums)
}

func turnDown(nums string, pos int) string {
	byteNums := []byte(nums)
	if byteNums[pos] == '0' {
		byteNums[pos] = '9'
	} else {
		byteNums[pos] -= 1
	}
	return string(byteNums)
}
