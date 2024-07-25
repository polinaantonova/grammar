package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Value      int   `json:"value"`
	LeftChild  *Node `json:"left,omitempty"`
	RightChild *Node `json:"right,omitempty"`
}

func ReverseTree2(node *Node) {
	if node == nil {
		return
	}
	tmp := node.RightChild
	node.RightChild = node.LeftChild
	node.LeftChild = tmp

	ReverseTree2(node.LeftChild)
	ReverseTree2(node.RightChild)
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	}
	if value < n.Value {
		return n.LeftChild.Search(value)
	}
	return n.RightChild.Search(value)
}

func (n *Node) Add(value int) {
	if value < n.Value {
		if n.LeftChild == nil {
			n.LeftChild = &Node{
				Value: value,
			}
			return
		}
		n.LeftChild.Add(value)
	}
	if value > n.Value {
		if n.RightChild == nil {
			n.RightChild = &Node{
				Value: value,
			}
			return
		}
		n.RightChild.Add(value)
	}
}

func BuildTree(slice []int) *Node {
	if len(slice) == 0 {
		return nil
	}
	tree := Node{
		Value: slice[0],
	}
	for _, value := range slice[1:] {
		tree.Add(value)
	}
	serialized, _ := json.MarshalIndent(tree, "", "    ")
	fmt.Println(string(serialized))
	return &tree
}
func main() {
	mySlice := []int{5, 10, 3, 6, 5, 2, 4, 7, 1, 8, 9}
	tree := BuildTree(mySlice)
	fmt.Println(tree.Search(100))
}
