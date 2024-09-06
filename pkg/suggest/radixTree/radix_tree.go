package radixTree

import (
	"github.com/polina/grammar/pkg/suggest/prefixTree"
)

type Node struct {
	value    string
	children []*Node
	payload  []*prefixTree.QueryStat
}

func NewNode() *Node {
	return &Node{}
}

type CommonPrefix struct {
	commonPrefix    string
	queryNameRemain string
	nodeValueRemain string
}

func Match(queryName string, nodeValue string) CommonPrefix {
	i := 0
	for {
		if i == len(queryName) || i == len(nodeValue) {
			break
		}
		if queryName[i] != nodeValue[i] {
			break
		}
		i++
	}
	res := CommonPrefix{commonPrefix: queryName[:i], queryNameRemain: queryName[i:], nodeValueRemain: nodeValue[i:]}
	return res
}

func (n *Node) Add(query *prefixTree.QueryStat) {
	match := Match(query.Name(), n.value)
	queryNameRemain := match.queryNameRemain
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
	var match CommonPrefix

	for _, child := range n.children {
		// [ab] a +
		// [a] ab +
		// [ab] ac +

		match = Match(queryNameRemain, child.value)

		if match.commonPrefix != "" {
			childToModify = child
			break
		}
	}

	if childToModify == nil {
		node := NewNode()
		node.value = queryNameRemain
		node.payload = append(node.payload, query)
		n.children = append(n.children, node)
		return
	}

	// [ab] ?a > [a[b]]
	if match.nodeValueRemain != "" {
		childToModify.value = match.commonPrefix

		node := NewNode()
		node.value = match.nodeValueRemain
		node.payload = childToModify.payload

		if len(childToModify.payload) < 7 {
			childToModify.payload = append(childToModify.payload, query)
		}
		childToModify.children = append(childToModify.children, node)
	}

	//[a[b]] ?abc > [a[b[c]]]
	//положи запрос в payload node, откуси от запроса общий префикс и ищи childToModify среди детей node
	if match.queryNameRemain != "" {
		childToModify.add(match.queryNameRemain, query)
	}
}

func (n *Node) Search(prefix string) []*prefixTree.QueryStat {
	var childToModify *Node
	var match CommonPrefix

	for _, child := range n.children {
		match = Match(prefix, child.value)

		if match.commonPrefix != "" {
			childToModify = child
			break
		}
	}

	if match.commonPrefix == "" {
		return nil
	}

	if len(match.queryNameRemain) != 0 {
		return childToModify.Search(match.queryNameRemain)
	}
	return childToModify.payload
}
