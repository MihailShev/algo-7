package binary_tree

type Tree struct {
	root *node
	size int
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *Tree) Search(key int) interface{} {
	n := search(key, t.root)

	if n != nil {
		return n.value
	} else {
		return nil
	}
}

func (t *Tree) Insert(key int, value interface{}) {
	t.size++
	if t.root == nil {
		t.root = newNode(key, value, nil)
	} else {
		insert(key, value, t.root)
	}
}

func (t *Tree) Remove(key int) interface{} {
	n := search(key, t.root)

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
			t.root = n.left
		}

		if n.right != nil {
			insertTree(t.root, n.right)
		}
	}

	t.size--

	return n.value
}

func (t *Tree) String() string {
	return treeToString(t.root, false)
}

func (t *Tree) newNode(key int, value interface{}, parent *node) *node {
	n := newNode(key, value, parent)

	return n
}
