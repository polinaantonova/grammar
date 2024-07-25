package prefixTree

type Node struct {
	children map[rune]*Node
	payload  []*Query
}

func (n *Node) Add(q *Query) {
	if n == nil {
		n.payload = append(n.payload, q)

	}
}
