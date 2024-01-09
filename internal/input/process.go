package input

import (
	"fmt"
	"io"
	"os"
)

func Process(action func(io.Reader)) {
	args := os.Args[1:]

	if len(args) == 0 {
		action(os.Stdin)
	} else {
		for _, path := range args {
			file, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}

			action(file)
			file.Close()
		}
	}
}
