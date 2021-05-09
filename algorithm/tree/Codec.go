package algorithm

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/serialize-and-deserialize-bst/
// 题目：297. 二叉树的序列化与反序列化
// 难度：困难
// 描述：序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
// 思路：BFS或DFS
type Codec struct {
}

// BFS
func (this *Codec) serialize(root *TreeNode) string {
	tokens := []string{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		if front != nil {
			tokens = append(tokens, strconv.FormatInt(int64(front.Val), 10))
			q = append(q, front.Left)
			q = append(q, front.Right)
		} else {
			tokens = append(tokens, "#")
		}
	}
	return strings.Join(tokens, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	if len(tokens) <= 0 {
		return nil
	}

	rootVal, _ := strconv.ParseInt(tokens[0], 10, 32)
	root := &TreeNode{int(rootVal), nil, nil}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		children := []**TreeNode{&front.Left, &front.Right}
		for j := 0; j < len(children) && (i+j) < len(tokens); j++ {
			if tokens[i+j] == "#" {
				continue
			}

			val, _ := strconv.ParseInt(tokens[i+j], 10, 32)
			*children[j] = &TreeNode{int(val), nil, nil}
			q = append(q, *children[j])
		}
		i += len(children)
	}
	return root
}

// DFS
func (this *Codec) serialize2(root *TreeNode) string {
	var ans []string
	this.serialize2Core(root, &ans)
	return strings.Join(ans, ",")
}

func (this *Codec) serialize2Core(node *TreeNode, ans *[]string) {
	if node == nil {
		*ans = append(*ans, "#")
		return
	}

	*ans = append(*ans, strconv.FormatInt(int64(node.Val), 10))
	this.serialize2Core(node.Left, ans)
	this.serialize2Core(node.Right, ans)
}

func (this *Codec) deserialize2(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	var pos int
	root := (*TreeNode)(nil)
	this.deserialize2Core(&pos, tokens, &root)
	return root
}

func (this *Codec) deserialize2Core(pos *int, tokens []string, node **TreeNode) {
	if *pos >= len(tokens) {
		return
	}

	if tokens[*pos] == "#" {
		*pos++
		return
	}

	val, _ := strconv.ParseInt(tokens[*pos], 10, 64)
	*node = &TreeNode{int(val), nil, nil}
	*pos++
	this.deserialize2Core(pos, tokens, &(*node).Left)
	this.deserialize2Core(pos, tokens, &(*node).Right)
}
