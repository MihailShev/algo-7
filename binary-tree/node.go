package binary_tree

import (
	"fmt"
	"strings"
)

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
	s := strings.Builder{}

	s.WriteString(fmt.Sprintf("deep %d\t ", n.deep))

	if showHeight {
		s.WriteString(fmt.Sprintf("height %d\t ", n.height))
	}

	s.WriteString(fmt.Sprintf("%s \t%d \t", t, n.key))

	if n.parent != nil {
		s.WriteString(fmt.Sprintf("parent: %d ", n.parent.key))
	} else {
		s.WriteString("null    ")
	}

	return s.String()
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
