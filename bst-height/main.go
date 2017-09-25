package main

import "fmt"

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

func NewBSTNode() *BSTNode {
	// set our own zero value of -1 for checking of root node
	b := BSTNode{}
	b.data = -1
	return &b
}

func findHeight(n *BSTNode) int {
	if n.left == nil && n.right == nil {
		return 0
	}
	lh := 0
	rh := 0
	if n.left != nil {
		lh = 1 + findHeight(n.left)
	}
	if n.right != nil {
		rh = 1 + findHeight(n.right)
	}
	max := lh
	if rh > lh {
		max = rh
	}
	return max
}

// insert a node into our BST based on the simplest definition
// https://en.wikipedia.org/wiki/Binary_search_tree
func insert(i int, root *BSTNode) {
	if root.data == -1 {
		root.data = i
		return
	}
	if i <= root.data {
		if root.left == nil {
			root.left = NewBSTNode()
			root.left.data = i
		} else {
			insert(i, root.left)
		}
	} else {
		if root.right == nil {
			root.right = NewBSTNode()
			root.right.data = i
		} else {
			insert(i, root.right)
		}
	}
}

func main() {
	list := []int{8, 6, 25, 14, 3, 7, 15, 24, 9}
	root := NewBSTNode()
	for _, i := range list {
		insert(i, root)
	}
	max := findHeight(root)
	fmt.Printf("max height: %d\n", max)
}
