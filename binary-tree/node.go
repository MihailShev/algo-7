package binary_tree

import (
	"fmt"
	"strings"
)

type NodesByDeep []*node

type node struct {
	key    int
	value  interface{}
	left   *node
	right  *node
	parent *node
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
		left:   nil,
		right:  nil,
		parent: parent,
		deep:   deep,
	}
}

func (n *node) String(t string) string {
	s := strings.Builder{}

	mes := fmt.Sprintf("deep %d ", n.deep)

	s.WriteString(mes)

	if n.parent != nil {
		mes = fmt.Sprintf("parent: %d ", n.parent.key)
		s.WriteString(mes)
	}

	mes = fmt.Sprintf("%s key: %d", t, n.key)

	s.WriteString(mes)

	return s.String()
}

func (n NodesByDeep) Len() int {
	return len(n)
}

func (n NodesByDeep) Less(i, j int) bool {
	return n[i].deep < n[j].deep
}

func (n NodesByDeep) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
