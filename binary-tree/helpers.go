package binary_tree

import (
	"fmt"
	"sort"
	"strings"
)

const (
	root   = "root"
	parent = "parent"
	left   = "left"
	right  = "right"
)

func treeToString(n *node, showHeight bool) string {
	if n == nil {
		return "Tree is empty"
	}
	arr := make([]*node, 0)
	nodeToArray(n, &arr)

	n.updateDeep()
	sort.Sort(NodesByDeep(arr))

	s := &strings.Builder{}

	for _, v := range arr {
		t := root
		if v.parent != nil {

			if v.key < v.parent.key {
				t = left
			} else {
				t = right
			}
		}
		s.WriteString(v.String(t, showHeight))
		s.WriteString("\n")
	}

	return s.String()
}

func nodeToString(n *node, t string, showHeight bool) string {
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
