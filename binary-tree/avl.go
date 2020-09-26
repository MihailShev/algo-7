package binary_tree

import "fmt"

type AVL struct {
	root *node
	size int
}

func (avl *AVL) Insert(key int, value interface{}) {
	avl.size++
	if avl.root == nil {
		avl.root = newNode(key, value, nil)
	} else {
		n := insert(key, value, avl.root)
		fmt.Println(n.key)
		balance(n.parent, avl)
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

func (avl *AVL) String() string {
	return treeToString(avl.root, true)
}
