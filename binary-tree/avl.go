package binary_tree

type AVL struct {
	root *node
	size int
}

func (avl *AVL) Insert(key int, value interface{}) {
	avl.size++
	if avl.root == nil {
		avl.root = newNode(key, value, nil)
	} else {
		insert(key, value, avl.root)
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
