package prefixTree

type QueryStat struct {
	name      string
	frequency int
}

func NewQuery(name string, freq int) *QueryStat {
	return &QueryStat{name: name, frequency: freq}
}

func (q *QueryStat) Frequency() int {
	return q.frequency
}

func (q *QueryStat) Name() string {
	return q.name
}
