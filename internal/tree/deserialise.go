package tree

import (
	"regexp"
	"strings"
)

func Deserialize(data string) *TreeNode {
	re := regexp.MustCompile(`\r?\n`)
	data = re.ReplaceAllString(data, "")

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
