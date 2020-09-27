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

// Поиск ноды, которая будет помещена на место удаленной
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
		// Если у nodeToRemove есть потомки справа, то ищем ноду с максимальным значением ключа справа
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

func (avl *AVL) String() string {
	return treeToString(avl.root, true)
}

func (avl *AVL) IsBalanced() bool {
	return isBalanced(avl.root)
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
