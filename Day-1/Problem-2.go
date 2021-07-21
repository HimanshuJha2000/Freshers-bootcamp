package main

import "fmt"

type treeNode struct {
	data  string
	left  *treeNode
	right *treeNode
}

func (root *treeNode) PreOrder() {
	if root != nil {
		fmt.Print(root.data)
		root.left.PreOrder()
		root.right.PreOrder()
	}
}

func (root *treeNode) PostOrder() {
	if root != nil {
		root.left.PostOrder()
		root.right.PostOrder()
		fmt.Print(root.data)
	}
}

func main() {
	root := treeNode{"+", nil, nil}
	root.left = &treeNode{"a", nil, nil}
	root.right = &treeNode{"-", nil, nil}
	root.right.left = &treeNode{"b", nil, nil}
	root.right.right = &treeNode{"c", nil, nil}

	fmt.Println("PreOrder traversal : ")
	root.PreOrder()
	fmt.Println("\nPostOrder traversal : ")
	root.PostOrder()
}