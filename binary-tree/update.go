package binary_tree

const (
	parent = "parent"
	left   = "left"
	right  = "right"
)

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

func balance(a *node, tree *AVL) {
	updateHeight(a)

	l := safeGetRelative(a, left)
	b := safeGetRelative(a, right)

	if safeGetNodeHeight(b)-safeGetNodeHeight(l) == 2 {
		leftRotate(a, b)
	} else {
		b = safeGetRelative(a, left)
		r := safeGetRelative(a, right)

		if safeGetNodeHeight(b)-safeGetNodeHeight(r) == 2 {
			rightRotate(a, b)
		}
	}

	if a.parent == nil {
		tree.root = a
	} else {
		balance(a.parent, tree)
	}

}

func leftRotate(a, b *node) {
	c := safeGetRelative(b, left)
	r := safeGetRelative(b, right)

	// Малое левое вращение
	if safeGetNodeHeight(c) <= safeGetNodeHeight(r) {
		smallLeftRotate(a, b, c)
		a = b
	} else if safeGetNodeHeight(c) > safeGetNodeHeight(r) {
		// Большое левое вращение
		m := safeGetRelative(c, right)
		smallRightRotate(b, c, m)
		n := safeGetRelative(c, left)
		smallLeftRotate(a, c, n)
		a = c
	}
}

func rightRotate(a, b *node) {
	l := safeGetRelative(b, left)
	c := safeGetRelative(b, right)

	if safeGetNodeHeight(c) <= safeGetNodeHeight(l) {
		// Малое правое вращение
		smallRightRotate(a, b, c)
		a = b
	} else {
		// Большое правое вращение
		m := safeGetRelative(c, left)
		smallLeftRotate(b, c, m)
		n := safeGetRelative(c, right)
		smallRightRotate(a, c, n)
		a = c
	}
}

func smallLeftRotate(a, b, c *node) {
	a.right = c
	updateHeight(a)

	if c != nil {
		c.parent = a
	}

	b.left = a
	updateHeight(b)

	b.parent = a.parent
	a.parent = b

	if b.parent != nil {
		if b.parent.left == a {
			b.parent.left = b
		} else {
			b.parent.right = b
		}
		updateHeight(b.parent)
	}
}

func smallRightRotate(a, b, c *node) {
	a.left = c
	updateHeight(a)
	if c != nil {
		c.parent = a
	}

	b.right = a
	updateHeight(b)

	b.parent = a.parent
	a.parent = b

	if b.parent != nil {
		if b.parent.left == a {
			b.parent.left = b
		} else {
			b.parent.right = b
		}

		updateHeight(b.parent)
	}

}

func safeGetRelative(n *node, t string) *node {

	if n != nil {

		if t == left {
			return n.left
		}

		if t == right {
			return n.right
		}

		if t == parent {
			return n.parent
		}
	}

	return nil
}

func updateHeight(n *node) {
	leftHeight := safeGetNodeHeight(n.left)
	rightHeight := safeGetNodeHeight(n.right)

	if leftHeight >= rightHeight {
		n.height = leftHeight + 1
	} else {
		n.height = rightHeight + 1
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
