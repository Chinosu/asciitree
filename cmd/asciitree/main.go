package main

import (
	"asciitree/internal/input"
	"asciitree/internal/tree"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input.Process(func(r io.Reader) {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}

		root := tree.Deserialize(buf.String())
		fmt.Println(tree.Ascii(root))
	})
}
