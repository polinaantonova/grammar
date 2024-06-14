package tree

import (
	"io"
	"strings"
)

type Node struct {
	Name     string
	Children []Node
}

func PrintNode(root Node, output io.Writer) {
	var printNode func(node Node, offset int)

	printNode = func(node Node, offset int) {
		prefix := strings.Builder{}
		prefix.Grow(offset * 4)
		for i := 0; i < offset; i++ {
			prefix.WriteString("....")
		}
		_, _ = io.WriteString(output, prefix.String())
		_, _ = io.WriteString(output, node.Name)
		_, _ = io.WriteString(output, "\n")
		for _, child := range node.Children {
			printNode(child, offset+1)
		}
	}

	printNode(root, 0)
}
