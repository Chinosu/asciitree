package tree

import "fmt"

func PrintLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currentLevelSize := len(queue)

		for i := 0; i < currentLevelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			fmt.Print(currentNode.Val, " ")

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		fmt.Println()
	}
}
