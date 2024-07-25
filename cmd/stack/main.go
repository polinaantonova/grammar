package main

import (
	"fmt"

	"github.com/polina/grammar/pkg/containers/stack"
)

type Node struct {
	Name     string
	Children []Node
}

func Visit(node Node) {
	x := stack.NewStack[Node](4)
	x.Push(node)

	for x.Size() > 0 {
		currentNode := x.Pop()
		fmt.Println(currentNode)
		for _, child := range currentNode.Children {
			x.Push(child)
		}
	}
}

func main() {
	node := Node{
		Name: "Root",
		Children: []Node{
			{
				Name: "Одежда",
				Children: []Node{
					{
						Name: "Женская",
					},
					{
						Name: "Мужская",
					},
				},
			},
			{
				Name: "Хобби",
			},
		},
	}
	Visit(node)
}
