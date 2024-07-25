package prefixTree

type Query struct {
	name      string
	frequency int
}

func NewQuery(name string, freq int) *Query {
	return &Query{name: name, frequency: freq}
}

func Frequency(q *Query) int {
	return q.frequency
}
