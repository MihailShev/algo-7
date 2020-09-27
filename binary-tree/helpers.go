package binary_tree

import (
	"sort"
	"strings"
)

func treeToString(root *node, showHeight bool) string {
	if root == nil {
		return "Tree is empty"
	}
	arr := make([]*node, 0)
	nodeToArray(root, &arr)

	root.updateDeep()
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
		s.WriteString(v.String(t, showHeight))
		s.WriteString("\n")
	}

	return s.String()
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

func safeGetNodeHeight(n *node) int {
	if n != nil {
		return n.height
	}

	return 0
}
