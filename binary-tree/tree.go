package binary_tree

import (
	"fmt"
	"sort"
	"strings"
)

type Tree struct {
	root    *node
	size    int
	maxDeep int
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) MaxDeep() int {
	return t.maxDeep
}

func (t *Tree) Search(key int) interface{} {
	n := search(key, t.root)

	if n != nil {
		return n.value
	} else {
		return nil
	}
}

func (t *Tree) Add(key int, value interface{}) {
	t.size++
	if t.root == nil {
		t.root = t.newNode(key, value, nil)
	} else {
		t.insert(key, value, t.root)
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
			t.root = n.left
		}

		if n.right != nil {
			insertTree(t.root.right, n.right)
		}
	}

	t.size--

	return n.value
}

func (t *Tree) String() string {
	arr := make([]*node, 0)
	some(t.root, &arr)
	sort.Sort(NodesByDeep(arr))
	s := &strings.Builder{}

	for _, v := range arr {
		t := ""
		if v.parent != nil {

			if v.key < v.parent.key {
				t = "left"
			} else {
				t = "right"
			}
		}
		s.WriteString(v.String(t))
		s.WriteString("\n")
	}

	return s.String()
}

func some(n *node, arr *[]*node) {
	if n != nil {
		*arr = append(*arr, n)

		if n.left != nil {
			some(n.left, arr)
		}

		if n.right != nil {
			some(n.right, arr)
		}
	}
}

func toString(n *node, t string, s *strings.Builder) {
	if n != nil {

		mes := fmt.Sprintf("deep %d ", n.deep)

		s.WriteString(mes)

		if n.parent != nil {
			mes = fmt.Sprintf("parent: %d ", n.parent.key)
			s.WriteString(mes)
		}

		mes = fmt.Sprintf("%s key: %d\n", t, n.key)

		s.WriteString(mes)

		if n.left != nil {
			toString(n.left, "left", s)
		}

		if n.right != nil {
			toString(n.right, "right", s)
		}

	}

	return
}

func search(key int, n *node) *node {

	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	}

	if key > n.key {
		return search(key, n.right)
	} else {
		return search(key, n.left)
	}
}

func (t *Tree) insert(key int, value interface{}, n *node) {
	if key == n.key {
		n.value = value
	} else if key > n.key {
		if n.right != nil {
			t.insert(key, value, n.right)
		} else {
			n.right = t.newNode(key, value, n)
		}
	} else {
		if n.left != nil {
			t.insert(key, value, n.left)
		} else {
			n.left = t.newNode(key, value, n)
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
			child.deep = parent.deep + 1
		}
	} else {
		if parent.left != nil {
			insertTree(parent.left, child)
		} else {
			parent.left = child
			child.parent = parent
			child.deep = parent.deep + 1
		}
	}
}

func (t *Tree) newNode(key int, value interface{}, parent *node) *node {
	n := newNode(key, value, parent)
	if n.deep > t.maxDeep {
		t.maxDeep = n.deep
	}

	return n
}
