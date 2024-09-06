package prefixTree

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNode_Add(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		query1 := &QueryStat{
			name: "ab",
		}

		query2 := &QueryStat{
			name: "ac",
		}

		node := NewNode()
		node.Add(query1)
		node.Add(query2)

		expectedNode := &Node{
			children: map[rune]*Node{
				'a': {
					children: map[rune]*Node{
						'b': {
							payload: []*QueryStat{
								query1,
							},
						},
						'c': {
							payload: []*QueryStat{
								query2,
							},
						},
					},
					payload: []*QueryStat{
						query1,
						query2,
					},
				},
			},
			payload: []*QueryStat{
				query1,
				query2,
			},
		}

		reflect.DeepEqual(node, expectedNode)
	})
	t.Run("search", func(t *testing.T) {
		query1 := &QueryStat{
			name: "ab",
		}

		query2 := &QueryStat{
			name: "ac",
		}

		query3 := &QueryStat{
			name: "abc",
		}

		node := NewNode()
		node.Add(query1)
		node.Add(query2)
		node.Add(query3)

		t.Run("ab", func(t *testing.T) {
			res := node.Search("ab")
			require.True(t, reflect.DeepEqual([]*QueryStat{query1, query3}, res))
		})

		t.Run("abc", func(t *testing.T) {
			res := node.Search("abc")
			require.True(t, reflect.DeepEqual([]*QueryStat{query3}, res))
		})

		t.Run("abcd", func(t *testing.T) {
			res := node.Search("abcd")
			require.Nil(t, nil, res)
		})
	})
}
