package radixTree

import (
	"github.com/polina/grammar/pkg/suggest/prefixTree"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

//func TestMatch(t *testing.T) {
//	t.Run("mum, mum", func(t *testing.T) {
//		res := Match("mum", "mum")
//
//		assert.Equal(t, res, []string{"mum", "", ""})
//	})
//
//	t.Run("mum, mumum", func(t *testing.T) {
//		res := Match("mum", "mumum")
//
//		assert.Equal(t, res, []string{"mum", "", "um"})
//	})
//
//	t.Run("mum, mumum", func(t *testing.T) {
//		res := Match("mum", "mumum")
//
//		assert.Equal(t, res, []string{"mum", "", "um"})
//	})
//
//	t.Run("mumum, mum", func(t *testing.T) {
//		res := Match("mumum", "mum")
//
//		assert.Equal(t, res, []string{"mum", "um", ""})
//	})
//
//	t.Run("mum, apple", func(t *testing.T) {
//		res := Match("mum", "apple")
//
//		assert.Equal(t, res, []string{"", "mum", "apple"})
//	})
//
//	t.Run("empty query", func(t *testing.T) {
//		res := Match((""), "mum")
//
//		assert.Equal(t, res, []string{"", "", "mum"})
//	})
//}

func TestNode_Add(t *testing.T) {

	t.Run("payload overfill", func(t *testing.T) {
		node := NewNode()
		queries := []*prefixTree.QueryStat{
			prefixTree.NewQuery("a", 10),
			prefixTree.NewQuery("b", 9),
			prefixTree.NewQuery("c", 8),
			prefixTree.NewQuery("d", 7),
			prefixTree.NewQuery("e", 6),
			prefixTree.NewQuery("f", 5),
			prefixTree.NewQuery("g", 4),
			prefixTree.NewQuery("h", 3),
			prefixTree.NewQuery("i", 2),
		}
		for _, query := range queries {
			node.Add(query)
		}

		require.Equal(t, node.payload, queries[:7])
	})

	t.Run("a, b, a", func(t *testing.T) {
		query1 := prefixTree.NewQuery("a", 100)
		query2 := prefixTree.NewQuery("b", 100)
		node := NewNode()
		node.Add(query1)
		node.Add(query2)
		node.Add(query1)

		t.Run("[a][b]", func(t *testing.T) {
			res := &Node{value: "",
				children: []*Node{
					{
						value: "a",
						payload: []*prefixTree.QueryStat{
							prefixTree.NewQuery("a", 100),
						},
					},
					{
						value: "b",
						payload: []*prefixTree.QueryStat{
							prefixTree.NewQuery("b", 100),
						},
					},
				},
				payload: []*prefixTree.QueryStat{
					prefixTree.NewQuery("a", 100),
					prefixTree.NewQuery("b", 100),
				},
			}
			require.True(t, reflect.DeepEqual(node, res))
		})

		t.Run("[a][b] search", func(t *testing.T) {
			t.Run("a", func(t *testing.T) {
				res := node.Search("a")
				require.Equal(t, []*prefixTree.QueryStat{query1}, res)
			})

			t.Run("b", func(t *testing.T) {
				res := node.Search("b")
				require.Equal(t, []*prefixTree.QueryStat{query2}, res)
			})

			t.Run("ab", func(t *testing.T) {
				res := node.Search("ab")
				require.Empty(t, res)
			})
		})
	})

	t.Run("[a[b]], abc", func(t *testing.T) {
		query1 := prefixTree.NewQuery("a", 100)
		query2 := prefixTree.NewQuery("ab", 100)
		query3 := prefixTree.NewQuery("abc", 100)

		node := NewNode()
		node.Add(query1)
		node.Add(query2)
		node.Add(query3)

		t.Run("[a[b[c]]", func(t *testing.T) {
			res := &Node{
				value: "",
				children: []*Node{
					{
						value: "a",
						children: []*Node{
							{
								value: "b",
								children: []*Node{
									{
										value: "c",
										payload: []*prefixTree.QueryStat{
											prefixTree.NewQuery("abc", 100),
										},
									},
								},
								payload: []*prefixTree.QueryStat{
									prefixTree.NewQuery("ab", 100),
									prefixTree.NewQuery("abc", 100),
								},
							},
						},
						payload: []*prefixTree.QueryStat{
							prefixTree.NewQuery("a", 100),
							prefixTree.NewQuery("ab", 100),
							prefixTree.NewQuery("abc", 100),
						},
					},
				},
				payload: []*prefixTree.QueryStat{
					prefixTree.NewQuery("a", 100),
					prefixTree.NewQuery("ab", 100),
					prefixTree.NewQuery("abc", 100),
				},
			}
			require.True(t, reflect.DeepEqual(node, res))
		})

		t.Run("a[b[c]] search", func(t *testing.T) {
			t.Run("a", func(t *testing.T) {
				res := node.Search("a")
				require.Equal(t, []*prefixTree.QueryStat{query1, query2, query3}, res)
			})

			t.Run("abc", func(t *testing.T) {
				res := node.Search("abc")
				require.Equal(t, []*prefixTree.QueryStat{query3}, res)
			})

			t.Run("abcd", func(t *testing.T) {
				res := node.Search("abcd")
				require.Empty(t, res)
			})
		})
	})

	t.Run("abcd, abc", func(t *testing.T) {
		query1 := prefixTree.NewQuery("abcd", 100)
		query2 := prefixTree.NewQuery("abc", 100)
		node := NewNode()
		node.Add(query1)
		node.Add(query2)

		t.Run("[abc[d]]", func(t *testing.T) {
			res := &Node{
				value: "",
				children: []*Node{
					{
						value: "abc",
						children: []*Node{
							{
								value: "d",
								payload: []*prefixTree.QueryStat{
									prefixTree.NewQuery("abcd", 100),
								},
							},
						},
						payload: []*prefixTree.QueryStat{
							prefixTree.NewQuery("abcd", 100),
							prefixTree.NewQuery("abc", 100),
						},
					},
				},
				payload: []*prefixTree.QueryStat{
					prefixTree.NewQuery("abcd", 100),
					prefixTree.NewQuery("abc", 100),
				},
			}
			require.True(t, reflect.DeepEqual(node, res))
		})

		t.Run("abc[d] search", func(t *testing.T) {
			t.Run("abc", func(t *testing.T) {
				res := node.Search("abc")
				require.True(t, reflect.DeepEqual([]*prefixTree.QueryStat{query1, query2}, res))
			})

			t.Run("d", func(t *testing.T) {
				res := node.Search("d")
				require.Empty(t, res)
			})

			t.Run("abcd", func(t *testing.T) {
				res := node.Search("abcd")
				require.True(t, reflect.DeepEqual([]*prefixTree.QueryStat{query1}, res))
			})
		})
	})

	t.Run("ab, ac", func(t *testing.T) {
		query1 := prefixTree.NewQuery("ab", 100)
		query2 := prefixTree.NewQuery("ac", 100)
		node := NewNode()
		node.Add(query1)
		node.Add(query2)

		t.Run("a[b, c] search", func(t *testing.T) {
			t.Run("ab", func(t *testing.T) {
				res := node.Search("ab")
				require.Equal(t, res, []*prefixTree.QueryStat{query1})
			})

			t.Run("a", func(t *testing.T) {
				res := node.Search("a")
				require.Equal(t, res, []*prefixTree.QueryStat{query1, query2})
			})

		})

		t.Run("a[b, c]", func(t *testing.T) {
			res := &Node{
				value: "",
				children: []*Node{
					{
						value: "a",
						children: []*Node{
							{
								value: "b",
								payload: []*prefixTree.QueryStat{
									prefixTree.NewQuery("ab", 100),
								},
							},
							{
								value: "c",
								payload: []*prefixTree.QueryStat{
									prefixTree.NewQuery("ac", 100),
								},
							},
						},
						payload: []*prefixTree.QueryStat{
							prefixTree.NewQuery("ab", 100),
							prefixTree.NewQuery("ac", 100),
						},
					},
				},
				payload: []*prefixTree.QueryStat{
					prefixTree.NewQuery("ab", 100),
					prefixTree.NewQuery("ac", 100),
				},
			}
			require.True(t, reflect.DeepEqual(node, res))
		})
	})
}
