package binary_tree

type node struct {
	key    int
	value  interface{}
	left   *node
	right  *node
	parent *node
	height int
	deep   int
}

func newNode(key int, value interface{}, parent *node) *node {
	var deep int

	if parent != nil {
		deep = parent.deep + 1
	} else {
		deep = 1
	}

	return &node{
		key:    key,
		value:  value,
		height: 1,
		left:   nil,
		right:  nil,
		parent: parent,
		deep:   deep,
	}
}

func (n *node) String(t string, showHeight bool) string {
	return nodeToString(n, t, showHeight)
}

func (n *node) updateDeep() {
	if n.parent != nil {
		n.deep = n.parent.deep + 1
	} else {
		n.deep = 1
	}

	if n.left != nil {
		n.left.updateDeep()
	}

	if n.right != nil {
		n.right.updateDeep()
	}
}
