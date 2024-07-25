package main

import (
	"fmt"
)

type Node struct {
	Name     string
	Children []Node
}

func Visit(node Node, path []string) {
	if len(node.Children) != 0 {
		path = append(path, node.Name)
	}

	if len(node.Children) == 0 {
		path = append(path, node.Name)
		fmt.Println(path[1:])

		//дописать if len(path)>0
		//path = path[:len(path)-1]

	}
	for _, child := range node.Children {
		Visit(child, path)
	}
}

func VisitTwoNodes(node Node, path []string) {
	path = append(path, node.Name)
	if len(path) == 3 {
		fmt.Println(path[1:])
	}
	if len(path) < 3 && len(node.Children) == 0 {
		fmt.Println(path[1:])
	}

	for _, child := range node.Children {
		VisitTwoNodes(child, path)
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
						Children: []Node{
							{
								Name: "Сумки",
								Children: []Node{
									{
										Name: "Кожаные",
									},
								},
							},
							{
								Name: "Платья",
							},
						},
					},
					{
						Name: "Мужская",
					},
				},
			},
			{
				Name: "Хобби",
				Children: []Node{
					{
						Name: "Мангалы",
					},
				},
			},
			{
				Name: "Авто",
			},
		},
	}
	var path []string
	Visit(node, path)
	VisitTwoNodes(node, path)

}
