package redBlackTree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_LeftRotation(t *testing.T) {

	t.Run("left rotation", func(t *testing.T) {
		//
		//
		//    x                   1
		//     \                   \
		//      y                   3
		//     / \                 /  \
		//    a   b               2    4
		//
		//After Left Rotation:
		//
		//      3
		//     / \
		//    1   4
		//     \
		//      2
		tree := Root{}
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(4)
		tree.Insert(2)
		res := TreeDepthTraversal(tree.root)
		assert.Equal(t, []int{1, 2, 3, 4}, res)
	})

	t.Run("right rotation", func(t *testing.T) {
		//Before Right Rotation:
		//
		//      x              4
		//     /              /
		//    y              2
		//   / \            / \
		//  a   b          1   3
		//
		//After Right Rotation:
		//
		//    2
		//   / \
		//  1   4
		//     /
		//    3

		//    Set y to be the left child of x.+
		//    Move y’s right subtree to x’s left subtree.+
		//    Update the parent of x and y.+
		//    Update x’s parent to point to y instead of x.+
		//    Set y’s right child to x.
		//    Update x’s parent to y.

		root := Root{}
		root.Insert(4)
		root.Insert(2)
		root.Insert(1)
		root.Insert(3)
		root.Insert(5)

		res := TreeDepthTraversal(root.root)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
	})

	t.Run("many nodes", func(t *testing.T) {
		tree := Root{}

		tree.Insert(7)
		tree.Insert(3)
		tree.Insert(18)
		tree.Insert(10)
		tree.Insert(22)
		tree.Insert(8)
		tree.Insert(11)
		tree.Insert(26)
		tree.Insert(2)
		tree.Insert(6)
		tree.Insert(13)
		res := TreeDepthTraversal(tree.root)
		fmt.Println(res)
	})

}
