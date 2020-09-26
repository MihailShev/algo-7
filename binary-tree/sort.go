package binary_tree

type NodesByDeep []*node

func (n NodesByDeep) Len() int {
	return len(n)
}

func (n NodesByDeep) Less(i, j int) bool {
	return n[i].deep < n[j].deep
}

func (n NodesByDeep) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
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
