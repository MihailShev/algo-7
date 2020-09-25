package binary_tree

func insert(key int, value interface{}, n *node) {
	if key == n.key {
		n.value = value
	} else if key > n.key {
		if n.right != nil {
			insert(key, value, n.right)
		} else {
			n.right = newNode(key, value, n)
		}
	} else {
		if n.left != nil {
			insert(key, value, n.left)
		} else {
			n.left = newNode(key, value, n)
		}
	}
}

func insertTree(parent *node, child *node) {
	if child.key > parent.key {
		if parent.right != nil {
			insertTree(parent.right, child)
		} else {
			parent.right = child
			child.parent = parent
		}
	} else {
		if parent.left != nil {
			insertTree(parent.left, child)
		} else {
			parent.left = child
			child.parent = parent
		}
	}
}

//???
func remove(root *node, key int) interface{} {
	n := search(key, root)

	if n == nil {
		return n
	}

	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}

		if n.left != nil {
			insertTree(n.parent, n.left)
		}

		if n.right != nil {
			insertTree(n.parent, n.right)
		}

	} else {
		if n.left != nil {
			n.left.parent = nil
			root = n.left
		}

		if n.right != nil {
			insertTree(root, n.right)
		}
	}

	return n.value
}
