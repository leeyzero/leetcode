package algorithm

import (
	"sort"
)

// https://leetcode-cn.com/problems/queue-reconstruction-by-height/
// 题目：406. 根据身高重建队列
// 难度：中等
// 解题思路：贪心，先固定一个维度
// 有两个维度h和k，如果同时考虑两个维度，问题会变得很复杂，可以先固定一个维度，问题是先固定那个维度呢？
// 如果先确定k，按k排序后，会发现k的排列不符合条件，h也不符合条件，两个维度还是没有确定下来
// 如果先确定h，按h排序后，将身高从高到低排序，身高相同，k小的站前面，此时确定了身高，也就是排前面的都比后面的高
// 此时只需要按照k重新插入队列即可，每次插入节点时，如果节点的k大于等于当前位置，直接插入即可；如果小于当前位置，
// 插入到正确位置即可，因为前面的节点都比该节点高，插入并不影响前面节点的排序
// 整个插入过程如下：
// 排序完：
// [[7,0], [7,1], [6,1], [5,0], [5,2]，[4,4]]
// 插入的过程：
// 插入[7,0]：[[7,0]]
// 插入[7,1]：[[7,0],[7,1]]
// 插入[6,1]：[[7,0],[6,1],[7,1]]
// 插入[5,0]：[[5,0],[7,0],[6,1],[7,1]]
// 插入[5,2]：[[5,0],[7,0],[5,2],[6,1],[7,1]]
// 插入[4,4]：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for i := 1; i < len(people); i++ {
		curr := people[i]
		j := i
		for curr[1] < j {
			people[j] = people[j-1]
			j--
		}
		if i != j {
			people[j] = curr
		}
	}
	return people
}
