package binary_tree

type Tree struct {
	root *node
	size int
}

type node struct {
	key int
	value interface{}
	left *node
	right *node
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Add(key int, value interface{}) {
	t.size++
	if t.root == nil {
		t.root = newNode(key, value)
	} else {
		t.insert(key, value, t.root)
	}
}

func (t *Tree) Search(key int) interface{} {
	if t.root != nil {
		return search(key, t.root)
	}

	return nil
}

func search(key int, n *node) interface{} {

	if key == n.key {
		return n.value
	}

	if key > n.key {
		return search(key, n.left)
	} else {
		return search(key, n.right)
	}
}

func (t * Tree) insert(key int, value interface{}, n *node) {
	if key == n.key {
		n.value = value
	} else if key > n.key {
		if n.left != nil {
			t.insert(key, value, n.left)
		} else {
			n.left = newNode(key, value)
		}
	} else {
		if n.right != nil {
			t.insert(key, value, n.right)
		} else {
			n.right = newNode(key, value)
		}
	}
}

func newNode(key int, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
}