package binary_tree

type AVL struct {
	root *node
	size int
}

func (avl *AVL) Size() int {
	return avl.size
}

func (avl *AVL) Insert(key int, value interface{}) {
	avl.size++

	if avl.root == nil {
		avl.root = newNode(key, value, nil)
	} else {
		n := insert(key, value, avl.root)

		if n.parent != nil {
			balance(n.parent, avl)
		}
	}
}

func (avl *AVL) Search(key int) interface{} {
	n := search(key, avl.root)

	if n != nil {
		return n.value
	} else {
		return nil
	}
}

func (avl *AVL) Remove(key int) interface{} {
	n := search(key, avl.root)

	if n == nil {
		return nil
	}

	avl.size--

	if n.left != nil || n.right != nil {
		var tmp *node

		tmp = findMax(n.left)

		if tmp != n.left {
			tmp.parent.right = nil
		} else {
			tmp.parent.left = nil
		}

		tmp.parent = n.parent

		if tmp.parent == nil {
			avl.root = tmp
		} else {
			if tmp.parent.left == n {
				tmp.parent.left = tmp
			} else {
				tmp.parent.right = tmp
			}
		}

		tmp.left = n.left
		tmp.right = n.right

		if tmp.left != nil {
			tmp.left.parent = tmp
		}

		if tmp.right != nil {
			tmp.right.parent = tmp
		}

		balance(tmp, avl)

	} else {
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}

		balance(n.parent, avl)
	}

	return n
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

func (avl *AVL) String() string {
	return treeToString(avl.root, true)
}
