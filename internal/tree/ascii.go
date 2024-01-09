package tree

import (
	"asciitree/internal/minmax"
	"fmt"
	"regexp"
	"strings"
)

const (
	padding = 2
)

func indent(lines []string, margin int) int {
	if margin >= 0 {
		return margin
	}

	spaces := strings.Repeat(" ", -margin)

	for i, line := range lines {
		lines[i] = spaces + line
	}

	return 0
}

func charCount(s string) int {
	count := 0
	for range s {
		count++
	}

	return count
}

func merge(left []string, right []string) []string {
	minSize := minmax.Min(len(left), len(right))
	offset := 0
	re := regexp.MustCompile(`\S.*`)
	for i := 0; i < minSize; i++ {
		replaced := re.ReplaceAllString(right[i], "")
		offset = minmax.Max(offset, charCount(left[i])+padding-charCount(replaced))
	}

	indent(right, -indent(left, offset))

	for i := 0; i < minSize; i++ {
		left[i] += right[i][charCount(left[i]):]
	}

	if len(right) > minSize {
		left = append(left, right[minSize:]...)
	}

	return left
}

func buildLines(node *TreeNode) []string {
	if node == nil {
		return nil
	}

	lines := merge(buildLines(node.Left), buildLines(node.Right))

	half := int(len(node.Val) / 2)
	i := half

	if len(lines) > 0 {
		var line string
		i = strings.Index(lines[0], "*")

		if node.Right == nil {
			line = strings.Repeat(" ", i) + "┌─┘"
			i += 2
		} else if node.Left == nil {
			i = indent(lines, i-2)
			line = strings.Repeat(" ", i) + "└─┐"
		} else {
			dist := len(lines[0]) - 1 - i

			repeatSpace := strings.Repeat(" ", i)
			repeatDashHalf1 := strings.Repeat("─", dist/2-1)
			repeatDashHalf2 := strings.Repeat("─", (dist-1)/2)

			line = fmt.Sprintf("%s┌%s┴%s┐", repeatSpace, repeatDashHalf1, repeatDashHalf2)

			i += dist / 2
		}

		lines[0] = line
	}

	lines = append([]string{strings.Repeat(" ", indent(lines, i-half)) + node.Val}, lines...)
	lines = append([]string{strings.Repeat(" ", i+minmax.Max(0, half-i)) + "*"}, lines...)
	return lines
}

func Ascii(root *TreeNode) string {
	return strings.Join(buildLines(root)[1:], "\n")
}
