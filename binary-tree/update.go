package binary_tree

const (
	parent = "parent"
	left   = "left"
	right  = "right"
)

func insert(key int, value interface{}, parent *node) *node {
	if key == parent.key {
		parent.value = value
		return parent
	} else if key > parent.key {
		if parent.right != nil {
			return insert(key, value, parent.right)
		} else {
			parent.right = newNode(key, value, parent)
			return parent.right
		}
	} else {
		if parent.left != nil {
			return insert(key, value, parent.left)
		} else {
			parent.left = newNode(key, value, parent)
			return parent.left
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

func balance(a *node, tree *AVL) {
	updateHeight(a)

	l := safeGetRelative(a, left)
	b := safeGetRelative(a, right)

	if safeGetNodeHeight(b)-safeGetNodeHeight(l) == 2 {
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

	b = safeGetRelative(a, left)
	r := safeGetRelative(a, right)

	if safeGetNodeHeight(b)-safeGetNodeHeight(r) == 2 {
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

	if a.parent == nil {
		tree.root = a
	} else {
		balance(a.parent, tree)
	}

}

func balance2(a *node, tree *AVL) {
	updateHeight(a)

	b := safeGetRelative(a, left)
	r := safeGetRelative(a, right)
	l := safeGetRelative(b, left)
	c := safeGetRelative(b, right)

	heightDifference := safeGetNodeHeight(b) - safeGetNodeHeight(r)

	if heightDifference == 2 {
		// Малое правое вращение
		if safeGetNodeHeight(c) <= safeGetNodeHeight(l) {
			smallRightRotate(a, b, c)
			a = b
		} else if safeGetNodeHeight(c) > safeGetNodeHeight(l) {
			// Большое правое вращение
			m := safeGetRelative(c, left)
			smallLeftRotate(b, c, m)
			n := safeGetRelative(c, right)
			smallRightRotate(a, c, n)
			a = c
		}
	}

	if heightDifference == -2 {
		// Малое левое вращение
		if safeGetNodeHeight(c) <= safeGetNodeHeight(l) {
			m := safeGetRelative(r, left)
			smallLeftRotate(a, r, m)
			a = r
		} else if safeGetNodeHeight(c) > safeGetNodeHeight(l) {
			// Большое левое вращение
			n := safeGetRelative(c, right)
			smallRightRotate(r, c, n)
			m := safeGetRelative(c, left)
			smallLeftRotate(a, c, m)
		}
	}

	if a.parent == nil {
		tree.root = a
	} else {
		balance(a.parent, tree)
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
