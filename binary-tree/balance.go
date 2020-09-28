package binary_tree

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

func updateHeight(n *node) {
	leftHeight := safeGetNodeHeight(n.left)
	rightHeight := safeGetNodeHeight(n.right)

	if leftHeight >= rightHeight {
		n.height = leftHeight + 1
	} else {
		n.height = rightHeight + 1
	}
}

func replaceNodeToRemove(nodeToRemove, nodeToReplace *node, avl *AVL) {
	// Заменяем ссылку на родителя у nodeToReplace
	nodeToReplace.parent = nodeToRemove.parent

	if nodeToReplace.parent == nil {
		avl.root = nodeToReplace
	} else {
		// Если удаляемая нода была у родителя слева, то вставляем nodeToReplace слева
		if nodeToReplace.parent.left == nodeToRemove {
			nodeToReplace.parent.left = nodeToReplace
		} else {
			// Если удаляемая нода была у родителя справа, то вставляем nodeToReplace справа
			nodeToReplace.parent.right = nodeToReplace
		}
	}

	nodeToReplace.left = nodeToRemove.left
	nodeToReplace.right = nodeToRemove.right

	if nodeToReplace.left != nil {
		nodeToReplace.left.parent = nodeToReplace
	}

	if nodeToReplace.right != nil {
		nodeToReplace.right.parent = nodeToReplace
	}
}

func findAndCutNodeToReplace(nodeToRemove *node, avl *AVL) *node {
	var nodeToReplace *node

	if nodeToRemove.left != nil {
		// Если у nodeToRemove есть потомки слева, то ищем ноду с максимальным значением ключа слева
		nodeToReplace = findMax(nodeToRemove.left)

		// Если nodeToReplace не прямой потомок nodeToRemove, то удаляем ссылку на nodeToReplace у родителя nodeToReplace
		if nodeToReplace != nodeToRemove.left {
			nodeToReplace.parent.right = nodeToReplace.left
			if nodeToReplace.parent.right != nil {
				nodeToReplace.parent.right.parent = nodeToReplace.parent
			}
			balance(nodeToReplace.parent, avl)
		} else {
			// Если nodeToReplace прямой потомок nodeToRemove, то удаляем у nodeToRemove ссылку на nodeToReplace
			nodeToReplace.parent.left = nodeToReplace.left
			if nodeToReplace.parent.left != nil {
				nodeToReplace.parent.left.parent = nodeToReplace.parent
			}
		}

	} else if nodeToRemove.right != nil {
		// Если у nodeToRemove есть потомки справа, то ищем ноду с минимальным значением ключа справа
		nodeToReplace = findMin(nodeToRemove.right)

		// Если nodeToReplace не прямой потомок nodeToRemove, то удаляем ссылку на nodeToReplace у родителя nodeToReplace
		if nodeToReplace != nodeToRemove.right {
			nodeToReplace.parent.left = nodeToReplace.right
			if nodeToReplace.parent.left != nil {
				nodeToReplace.parent.left.parent = nodeToReplace.parent
			}
			balance(nodeToReplace.parent, avl)
		} else {
			// Если nodeToReplace прямой потомок nodeToRemove, то удаляем ссылку у nodeToRemove на nodeToReplace
			nodeToReplace.parent.right = nodeToReplace.right
			if nodeToReplace.right != nil {
				nodeToReplace.parent.right.parent = nodeToReplace.parent
			}
		}
	}

	return nodeToReplace
}

func findMin(n *node) *node {
	if n.left != nil {
		return findMin(n.left)
	}

	return n
}

func findMax(n *node) *node {
	if n.right != nil {
		return findMax(n.right)
	}

	return n
}

func safeGetNodeHeight(n *node) int {
	if n != nil {
		return n.height
	}

	return 0
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

func isBalanced(n *node) bool {
	if n != nil {

		l := safeGetNodeHeight(n.left)
		r := safeGetNodeHeight(n.right)

		if l > r && l-r > 1 {
			return false
		}

		if l < r && r-l > 1 {
			return false
		}

		return !isBalanced(n.left) || isBalanced(n.right)
	}

	return true
}
