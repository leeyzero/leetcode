package skiplist

import (
	"math/rand"
	"time"
)

// https://leetcode.cn/problems/design-skiplist/
// 1206. 设计跳表
// 难度：困难
/*
参考：https://en.wikipedia.org/wiki/Skip_list

示例 1:
输入
["Skiplist", "add", "add", "add", "search", "add", "search", "erase", "erase", "search"]
[[], [1], [2], [3], [0], [4], [1], [0], [1], [1]]
输出
[null, null, null, null, false, null, true, false, true, false]

解释
Skiplist skiplist = new Skiplist();
skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // 返回 false
skiplist.add(4);
skiplist.search(1);   // 返回 true
skiplist.erase(0);    // 返回 false，0 不在跳表中
skiplist.erase(1);    // 返回 true
skiplist.search(1);   // 返回 false，1 已被擦除


提示:
0 <= num, target <= 2 * 104
调用search, add,  erase操作次数不大于 5 * 104
*/

const (
	FACTOR_P  = 0.25
	MAX_LEVEL = 32
)

type SkiplistNode struct {
	Val     int
	Forward []*SkiplistNode
}

type Skiplist struct {
	head  *SkiplistNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head:  &SkiplistNode{-1, make([]*SkiplistNode, MAX_LEVEL)},
		level: 0,
	}
}

func (this *Skiplist) Search(target int) bool {
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 找到第i层小于且最接入target的元素
		for curr.Forward[i] != nil && curr.Forward[i].Val < target {
			curr = curr.Forward[i]
		}
	}
	curr = curr.Forward[0]
	return curr != nil && curr.Val == target
}

func (this *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, MAX_LEVEL)
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && curr.Forward[i].Val < num {
			curr = curr.Forward[i]
		}
		update[i] = curr
	}

	lvl := this.randomLevel()
	if lvl > this.level {
		for i := this.level; i < lvl; i++ {
			update[i] = this.head
		}
		this.level = lvl
	}

	newNode := &SkiplistNode{num, make([]*SkiplistNode, lvl)}
	for i := 0; i < lvl; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, MAX_LEVEL)
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		for curr.Forward[i] != nil && curr.Forward[i].Val < num {
			curr = curr.Forward[i]
		}
		update[i] = curr
	}

	curr = curr.Forward[0]
	if curr == nil || curr.Val != num {
		return false
	}

	for i := 0; i < this.level && update[i].Forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].Forward[i] = curr.Forward[i]
	}

	// 更新level
	for this.level > 0 && this.head.Forward[this.level-1] == nil {
		this.level--
	}

	return true
}

func (this *Skiplist) randomLevel() int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	level := 1
	for level < MAX_LEVEL && rnd.Float64() < FACTOR_P {
		level++
	}
	return level
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
