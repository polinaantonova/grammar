package tree

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PrintNode(t *testing.T) {
	node := Node{
		Name: "S",
		Children: []Node{
			{
				Name: "NP",
				Children: []Node{
					{Name: "N"},
				},
			},
			{
				Name: "VP",
				Children: []Node{
					{Name: "V"},
					{Name: "NP"},
				},
			},
		},
	}
	expected := `S
....NP
........N
....VP
........V
........NP
`
	buf := bytes.Buffer{}
	PrintNode(node, &buf)

	require.Equal(t, expected, buf.String())
}

func TestNode_Height(t *testing.T) {
	t.Run("several nodes", func(t *testing.T) {
		node := Node{
			Name: "S",
			Children: []Node{
				{
					Name: "NP",
					Children: []Node{
						{Name: "N"},
					},
				},
				{
					Name: "VP",
					Children: []Node{
						{Name: "V"},
						{Name: "NP"},
					},
				},
			},
		}
		require.Equal(t, 6, node.Height())
	})

	t.Run("one", func(t *testing.T) {
		node := Node{
			Name: "S",
		}

		require.Equal(t, 1, node.Height())
	})

}
