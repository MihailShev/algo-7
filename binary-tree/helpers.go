package binary_tree

import (
	"sort"
	"strings"
)

func treeToString(t *Tree) string {
	arr := make([]*node, 0)
	nodeToArray(t.root, &arr)
	sort.Sort(NodesByDeep(arr))
	s := &strings.Builder{}

	for _, v := range arr {
		t := "root"
		if v.parent != nil {

			if v.key < v.parent.key {
				t = "left"
			} else {
				t = "right"
			}
		}
		s.WriteString(v.String(t))
		s.WriteString("\n")
	}

	return s.String()
}

func nodeToArray(n *node, arr *[]*node) {
	if n != nil {
		*arr = append(*arr, n)

		if n.left != nil {
			nodeToArray(n.left, arr)
		}

		if n.right != nil {
			nodeToArray(n.right, arr)
		}
	}
}

func search(key int, n *node) *node {

	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	}

	if key > n.key {
		return search(key, n.right)
	} else {
		return search(key, n.left)
	}
}

func insertTree(parent *node, child *node) {
	if child.key > parent.key {
		if parent.right != nil {
			insertTree(parent.right, child)
		} else {
			parent.right = child
			child.parent = parent
			child.updateDeep()
		}
	} else {
		if parent.left != nil {
			insertTree(parent.left, child)
		} else {
			parent.left = child
			child.parent = parent
			child.updateDeep()
		}
	}
}
