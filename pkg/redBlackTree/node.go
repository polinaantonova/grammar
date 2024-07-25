/*

If the parent is red, the tree might violate the Red Property, requiring fixes.
*/

package redBlackTree

type Node struct {
	value      int
	color      string
	leftChild  *Node
	rightChild *Node
	parent     *Node
}

type Root struct {
	root *Node
}

func (r *Root) Insert(value int) {
	if r.root == nil {
		r.root = &Node{
			value: value,
			color: "black",
		}
		return
	}
	r.root.Insert(value)
}

func (n *Node) Insert(value int) {
	//Insert the new node like in a standard BST.
	if value < n.value {
		if n.leftChild == nil {
			n.leftChild = &Node{
				value:  value,
				color:  "red",
				parent: n,
			}
			FixViolations(n.leftChild)
			return
		}
		n.leftChild.Insert(value)
	}
	if value > n.value {
		if n.rightChild == nil {
			n.rightChild = &Node{
				value:  value,
				color:  "red",
				parent: n,
			}
			FixViolations(n.rightChild)
		}
		n.rightChild.Insert(value)
	}

}

func (n *Node) GetGrandparent() *Node {
	if n.parent == nil {
		return nil
	}
	return n.parent.parent
}

func (n *Node) GetUncle() *Node {
	grandparent := n.parent.parent

	if n.parent == nil || n.parent.parent == nil {
		return nil
	}
	if n.parent == grandparent.leftChild {
		return grandparent.rightChild
	}
	if n.parent == grandparent.rightChild {
		return grandparent.leftChild
	}
	return nil
}

func (n *Node) LeftRotation() {
	//
	//    /                   \
	//    x                   x
	//   /  \               /   \
	//  ?    y             ?     y
	//      / \                 /  \
	//     a   b               a    b
	//
	//After Left Rotation:
	//
	//      y
	//     / \
	//    x   b
	//     \
	//      a
	//Left Rotation Steps:
	//
	//   1 Set y to be the right child of n. +
	//   2 Move y’s left subtree to x’s right subtree. +
	//   3 Update the parent of x and y. +
	//   4 Update x’s parent to point to y instead of x. +
	//   5 Set y’s left child to x. +
	//   6 Update x	’s parent to y.
	//1
	y := n.rightChild
	//2
	n.rightChild = y.leftChild

	//5
	if y.leftChild != nil {
		y.leftChild.parent = n
	}

	//3
	y.parent = n.parent

	//4
	if n.parent == nil {
		n.parent = y
	} else if n == n.parent.leftChild {
		n.parent.leftChild = y
	} else {
		n.parent.rightChild = y
	}

	//6
	y.leftChild = n
	n.parent = y
}

func (n *Node) RightRotation() {
	//Before Right Rotation:
	//
	//      x
	//     /
	//    y
	//   / \
	//  a   b
	//
	//After Right Rotation:
	//
	//    y
	//   / \
	//  a   x
	//     /
	//    b

	//    Set y to be the left child of x.+
	//    Move y’s right subtree to x’s left subtree.+
	//    Update the parent of x and y.+
	//    Update x’s parent to point to y instead of x.+
	//    Set y’s right child to x.
	//    Update x’s parent to y.

	y := n.leftChild
	//2
	n.leftChild = y.rightChild

	//5
	if y.rightChild != nil {
		y.rightChild.parent = n
	}

	//3
	y.parent = n.parent
	//4
	if n.parent == nil {
		n.parent = y
	} else if n == n.parent.leftChild {
		n.parent.leftChild = y
	} else {
		n.parent.rightChild = y
	}

	//6
	y.rightChild = n
	n.parent = y
}

func FixViolations(n *Node) {
	//Случай 1: Текущий узел N в корне дерева. В этом случае, он перекрашивается в чёрный цвет, чтобы оставить верным Свойство 2 (Корень — чёрный). Так как это действие добавляет один чёрный узел в каждый путь, Свойство 5 (Все пути от любого данного узла до листовых узлов содержат одинаковое число чёрных узлов) не нарушается.
	if n.parent == nil {
		n.color = "black"
		return
	}

	//Случай 2: Предок P текущего узла чёрный, то есть Свойство 4 (Оба потомка каждого красного узла — чёрные) не нарушается. В этом случае дерево остаётся корректным. Свойство 5 (Все пути от любого данного узла до листовых узлов содержат одинаковое число чёрных узлов) не нарушается, потому что текущий узел N имеет двух чёрных листовых потомков, но так как N является красным, путь до каждого из этих потомков содержит такое же число чёрных узлов, что и путь до чёрного листа, который был заменен текущим узлом, так что свойство остается верным.
	if n.parent.color == "black" {
		return
	}

	//Случай 3: Если и родитель P, и дядя U — красные, то они оба могут быть перекрашены в чёрный, и дедушка G станет красным (для сохранения свойства 5 (Все пути от любого данного узла до листовых узлов содержат одинаковое число чёрных узлов)). Теперь у текущего красного узла N чёрный родитель. Так как любой путь через родителя или дядю должен проходить через дедушку, число чёрных узлов в этих путях не изменится. Однако, дедушка G теперь может нарушить свойства 2 (Корень — чёрный) или 4 (Оба потомка каждого красного узла — чёрные) (свойство 4 может быть нарушено, так как родитель G может быть красным). Чтобы это исправить, вся процедура рекурсивно выполняется на G из случая 1. 	uncle := n.GetUncle()
	uncle := n.GetUncle()
	grandparent := n.GetGrandparent()

	if uncle != nil && uncle.color == "red" {
		n.parent.color = "black"
		uncle.color = "black"
		grandparent.color = "red"
		FixViolations(grandparent)
		return
	}

	//Случай 4: Родитель P является красным, но дядя U — чёрный. Также, текущий узел N — правый потомок P, а P в свою очередь — левый потомок своего предка G. В этом случае может быть произведен поворот дерева, который меняет роли текущего узла N и его предка P. Тогда, для бывшего родительского узла P в обновленной структуре используем случай 5, потому что Свойство 4 (Оба потомка любого красного узла — чёрные) все ещё нарушено. Вращение приводит к тому, что некоторые пути (в поддереве, обозначенном «1» на схеме) проходят через узел N, чего не было до этого. Это также приводит к тому, что некоторые пути (в поддереве, обозначенном «3») не проходят через узел P. Однако, оба эти узла являются красными, так что Свойство 5 (Все пути от любого данного узла до листовых узлов содержат одинаковое число чёрных узлов) не нарушается при вращении. Однако Свойство 4 всё ещё нарушается, но теперь задача сводится к Случаю 5.

	if n == n.parent.rightChild && n.parent == grandparent.leftChild {
		grandparent.LeftRotation()
		n = n.leftChild

	} else if n == n.parent.leftChild && n.parent == grandparent.rightChild {
		grandparent.RightRotation()
		n = n.rightChild

	}

	//Случай 5: Родитель P является красным, но дядя U — чёрный, текущий узел N — левый потомок P и P — левый потомок G. В этом случае выполняется поворот дерева на G. В результате получается дерево, в котором бывший родитель P теперь является родителем и текущего узла N и бывшего дедушки G. Известно, что G — чёрный, так как его бывший потомок P не мог бы в противном случае быть красным (без нарушения Свойства 4). Тогда цвета P и G меняются и в результате дерево удовлетворяет Свойству 4 (Оба потомка любого красного узла — чёрные). Свойство 5 (Все пути от любого данного узла до листовых узлов содержат одинаковое число чёрных узлов) также остается верным, так как все пути, которые проходят через любой из этих трех узлов, ранее проходили через G, поэтому теперь они все проходят через P. В каждом случае, из этих трёх узлов только один окрашен в чёрный.
	n.parent.color = "black"
	grandparent.color = "red"
	if n == n.parent.leftChild {
		grandparent.RightRotation()
	} else {
		grandparent.LeftRotation()
	}
}

func TreeDepthTraversal(node *Node) []int {
	var res []int
	var f func(node *Node)

	f = func(node *Node) {
		if node == nil {
			return
		}
		f(node.leftChild)
		res = append(res, node.value)
		f(node.rightChild)
	}
	RootCheck(node)
	f(node)
	return res
}

func RootCheck(node *Node) {
	if node.parent == nil {
		return
	}
	RootCheck(node.parent)

}
