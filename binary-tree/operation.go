package binary_tree

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

func insert(key int, value interface{}, parent *node) *node {
	var res *node

	if key == parent.key {
		parent.value = value
		res = parent
	} else if key > parent.key {
		if parent.right != nil {
			res = insert(key, value, parent.right)
		} else {
			parent.right = newNode(key, value, parent)
			res = parent.right
		}
	} else {
		if parent.left != nil {
			res = insert(key, value, parent.left)
		} else {
			parent.left = newNode(key, value, parent)
			res = parent.left
		}
	}

	return res
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
