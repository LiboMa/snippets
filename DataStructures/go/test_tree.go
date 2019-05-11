package main

import (
	"fmt"
	//	"github.com/disiqueira/gotree"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

//func addnode(Node) (n Node, data int) {
//
//	if n.data == nil {
//		n.data = data
//	} else if n.data < data {
//
//		n.left = n
//		n.left.data = data
//
//	} else {
//		n.right = n
//		n.right.data = data
//	}
//
//	return n
//}

func NewNode(data int) *Node {

	root := &Node{data, nil, nil}

	return root
}

type Tree interface {
	NewNode(data int) Node
	AddNode(n Node, data int) Node
}

func AddNode(n *Node, data int) *Node {

	if n.data != 0 {
		if n.data > data {
			if n.left == nil {
				n.left = NewNode(data)
			} else {
				n.left = AddNode(n.left, data)
			}
		} else if n.data < data {
			if n.right == nil {
				n.right = NewNode(data)
			} else {
				n.right = AddNode(n.right, data)
			}
		}

	} else {
		n.data = data
	}
	return n

}
func InorderRecursive(root *Node) {

	if root == nil {
		return
	}

	InorderRecursive(root.left)
	fmt.Printf("%d ", root.data)
	InorderRecursive(root.right)
}

func PreorderRecursive(root *Node) {

	if root == nil {
		return
	}

	PreorderRecursive(root.right)
	fmt.Printf("%d ", root.data)
	PreorderRecursive(root.left)
}

func main() {
	tree := NewNode(0)
	tree = AddNode(tree, 1)
	tree = AddNode(tree, 2)
	tree = AddNode(tree, 4)
	tree = AddNode(tree, 10)
	tree = AddNode(tree, 6)
	tree = AddNode(tree, 80)
	tree = AddNode(tree, 8)
	tree = AddNode(tree, 9)
	fmt.Println("tree: ", tree)
	//fmt.Println("tree: ", tree.right)
	InorderRecursive(tree)
	//PreorderRecursive(tree)

}
