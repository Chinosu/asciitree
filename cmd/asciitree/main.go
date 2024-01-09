package main

import (
	"asciitree/internal/tree"
	"fmt"
)

func main() {
	// fmt.Println("hello world")

	serializedTree := "1,2,#,#,3,#,#"
	root := tree.Deserialize(serializedTree)
	tree.PrintLevelOrder(root)

	fmt.Println()
	fmt.Println()

	fmt.Println(tree.Print(root))
}
