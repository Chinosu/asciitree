package tree

import (
	"log"
	"strconv"
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

		val, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatalf("Error converting string to integer: %v", err)
		}

		node := NewTreeNode(val)
		node.Left = buildTree()
		node.Right = buildTree()

		return node
	}

	return buildTree()
}
