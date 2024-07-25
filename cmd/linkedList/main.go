package main

import (
	"fmt"

	"github.com/polina/grammar/pkg/containers/linkedList"
)

type Node struct {
	Value    int
	Children []Node
}

func main() {
	node := Node{
		Value: 1,
		Children: []Node{
			{
				Value: 2,
				Children: []Node{
					{
						Value: 3,
					},
				},
			},
			{
				Value: 4,
			},
		},
	}

	l := linkedList.LinkedList[Node]{}
	l.AddLast(node)

	for l.IsEmpty() == false {
		currentNode := l.PopFirst()
		fmt.Println(currentNode)
		for _, child := range currentNode.Children {
			l.AddLast(child)
		}
	}
}
