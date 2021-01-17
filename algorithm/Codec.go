package algorithm

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

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
			tokens = append(tokens, "null")
		}
	}
	return "[" + strings.Join(tokens, ",") + "]"
}

func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 2 || (data[0] != '[' && data[1] != ']') {
		return nil
	}

	tokens := strings.Split(data[1:len(data)-1], ",")
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
		for j := 0; j < len(children); j++ {
			if tokens[i+j] == "null" {
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

