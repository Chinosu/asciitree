package tree

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val string) *TreeNode {
	return &TreeNode{Val: val}
}
