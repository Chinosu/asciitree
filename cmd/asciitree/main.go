package main

import (
	"asciitree/internal/tree"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	serializedTree := "1,uwu,#,#,3,4,#,#,5,#,#"
	root := tree.Deserialize(serializedTree)
	tree.PrintLevelOrder(root)
}
