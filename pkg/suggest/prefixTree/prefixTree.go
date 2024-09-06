package prefixTree

type Node struct {
	children map[rune]*Node
	payload  []*QueryStat
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node, 1),
	}
}

func (n *Node) Add(query *QueryStat) { //w в худшем случае w*a
	currentNode := n
	currentNode.payload = append(currentNode.payload, query)

	for _, r := range query.name { //w
		if node, ok := currentNode.children[r]; ok { //1
			currentNode = node
		} else {
			node = NewNode()
			currentNode.children[r] = node
			currentNode = node
		}
		if len(currentNode.payload) <= 7 {
			currentNode.payload = append(currentNode.payload, query)
		} else {
			continue
		}
	}
}

func (n *Node) Search(prefix string) []*QueryStat { //w
	currentNode := n

	for _, r := range prefix { //w
		if node, ok := currentNode.children[r]; ok {
			currentNode = node
		} else {
			return nil
		}
	}
	return currentNode.payload
}
