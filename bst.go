package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(value int) {
	newNode := &Node{value, nil, nil}
	if bst.root == nil {
		bst.root = newNode
	} else {

	}
}

func insertNode(root *Node, newNode *Node) {
	if newNode.value < root.value {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.value > root.value {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}

func main() {
	fmt.Println("bst")
}
