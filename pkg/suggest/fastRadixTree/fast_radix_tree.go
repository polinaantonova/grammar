package fastRadixTree

import (
	"github.com/polina/grammar/pkg/suggest/prefixTree"
	"sort"
)

type Node struct {
	value    string
	children []*Node
	payload  []*prefixTree.QueryStat
}

func NewNode() *Node {
	return &Node{}
}

func Match(queryName string, nodeValue string) int {

	var i, minLen int

	if len(queryName) < len(nodeValue) {
		minLen = len(queryName)
	} else {
		minLen = len(nodeValue)
	}

	for { // w
		if i == minLen {
			break
		}
		if queryName[i] != nodeValue[i] {
			break
		}
		i++
	}
	return i
}

func (n *Node) BinarySearch(node *Node) int {
	if n.children == nil {
		return 0
	}

	indexStart := 0
	indexEnd := len(n.children) - 1
	indexMiddle := 0

	value := node.value[0]
	min := n.children[indexStart].value[0]
	max := n.children[indexEnd].value[0]
	middle := n.children[indexMiddle].value[0]

	if value < min {
		return 0
	}
	if value > max {
		return indexEnd + 1
	}

	for indexStart <= indexEnd {
		indexMiddle = (indexStart + indexEnd) >> 1

		//этого случая быть не должно
		if value == middle {
			return indexMiddle
		}

		if value < middle {
			indexEnd = indexMiddle - 1
		}

		if value > middle {
			indexStart = indexMiddle + 1
		}
	}
	return indexStart
}

// build N * a * w * w
func (n *Node) Add(query *prefixTree.QueryStat) { // a * w * w
	match := Match(query.Name(), n.value) //w
	queryNameRemain := query.Name()[match:]

	n.add(queryNameRemain, query)
}

func (n *Node) add(queryNameRemain string, query *prefixTree.QueryStat) {
	var found bool
	for _, q := range n.payload {
		if query.Name() == q.Name() {
			found = true
			break
		}
	}
	if !found && len(n.payload) < 7 {
		n.payload = append(n.payload, query)
	}

	var childToModify *Node
	var match int
	var commonPrefix, nodeValueRemain string

	for _, child := range n.children { //a
		// [ab] a +
		// [a] ab +
		// [ab] ac +

		match = Match(queryNameRemain, child.value) //w
		commonPrefix = queryNameRemain[:match]
		queryNameRemain = queryNameRemain[match:]
		nodeValueRemain = child.value[match:]

		if commonPrefix != "" {
			childToModify = child
			break
		}
	}

	if childToModify == nil { //a
		node := NewNode()
		node.value = queryNameRemain
		node.payload = append(node.payload, query)

		//бинарным поиском ищем место для node
		nodeIndex := n.BinarySearch(node) //log a

		if n.children == nil || nodeIndex == len(n.children) {
			n.children = append(n.children, node) //a
		} else { //a
			n.children = append(n.children, nil)
			copy(n.children[(nodeIndex+1):], n.children[nodeIndex:])
			n.children[nodeIndex] = node
		}
		return
	}

	// [ab] ?a > [a[b]]
	if nodeValueRemain != "" {
		childToModify.value = commonPrefix

		node := NewNode()
		node.value = nodeValueRemain
		node.payload = childToModify.payload

		if len(childToModify.payload) < 7 {
			childToModify.payload = append(childToModify.payload, query)
		}
		childToModify.children = append(childToModify.children, node) //a
	}

	//[a[b]] ?abc > [a[b[c]]]
	//положи запрос в payload node, откуси от запроса общий префикс и ищи childToModify среди детей node
	if queryNameRemain != "" { //w
		childToModify.add(queryNameRemain, query)
	}
}

// a * w * w
func (n *Node) Search(prefix string) []*prefixTree.QueryStat {
	if prefix == "" {
		return n.payload
	}

	var childToModify *Node
	var match int

	if len(n.children) == 0 {
		return nil
	}
	//бинарным поиском сравниваем первую букву префикса и первую букву nodevalue
	i := sort.Search(len(n.children), func(i int) bool { //log a
		return n.children[i].value[0] >= prefix[0]
	})

	if i == len(n.children) {
		return nil
	}

	match = Match(prefix, n.children[i].value) //w

	if match == 0 {
		return nil
	}

	childToModify = n.children[i]

	if match == len(prefix) {
		return childToModify.payload
	}

	queryNameRemain := prefix[match:]
	return childToModify.Search(queryNameRemain)
}
