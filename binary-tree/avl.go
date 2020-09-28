package binary_tree

type AVL struct {
	root *node
	size int
}

func (avl *AVL) Size() int {
	return avl.size
}

func (avl *AVL) Clear() {
	avl.root = nil
	avl.size = 0
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
	nodeToRemove := search(key, avl.root)

	if nodeToRemove == nil {
		return nil
	}

	avl.size--

	// Находим и вырезаем из текущей позиции ноду, которая должна заменить удаляемую
	nodeToReplace := findAndCutNodeToReplace(nodeToRemove, avl)

	// Если нода для замены не найдена, значит у удаляемой ноды нет потомков
	// и можно просто стереть ссылку на нее у родителя
	if nodeToReplace == nil {
		// Если удаляемая нода не имеет потомков, то удаляем ссылки на нее у ее родителя
		if nodeToRemove.parent != nil {
			if nodeToRemove.parent.left == nodeToRemove {
				nodeToRemove.parent.left = nil
			} else {
				nodeToRemove.parent.right = nil
			}
			balance(nodeToRemove.parent, avl)
		} else {
			avl.root = nil
		}
	} else {
		replaceNodeToRemove(nodeToRemove, nodeToReplace, avl)
		balance(nodeToReplace, avl)
	}

	return nodeToRemove
}

func (avl *AVL) String() string {
	return treeToString(avl.root, true)
}

func (avl *AVL) IsBalanced() bool {
	return isBalanced(avl.root)
}
