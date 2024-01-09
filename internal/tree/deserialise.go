package tree

import (
	"strings"
)

func Deserialize(data string) *TreeNode {
	elements := strings.Split(data, ",")
	index := 0

	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		if index >= len(elements) {
			return nil
		}

		elem := elements[index]
		index++

		if elem == "#" {
			return nil
		}

		node := NewTreeNode(elem)
		node.Left = buildTree()
		node.Right = buildTree()

		return node
	}

	return buildTree()
}
