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

//func (node *Node) NewNode(data int) *Node {
func NewNode(data int) *Node {

	root := &Node{data, nil, nil}

	return root
}

//func (node *Node) AddNode(n *Node, data int) *Node {

// Method 1 Using func instead OO
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

// Method Using OO to add value
func (n *Node) Add(data int) *Node {

	if n.data != 0 {

		if n.data > data {

			if n.left == nil {
				n.left = NewNode(data)
			} else {
				n.left = n.left.Add(data)
			}

		} else if n.data < data {
			if n.right == nil {
				n.right = NewNode(data)
			} else {
				n.right = n.right.Add(data)
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

func PostorderRecursive(root *Node) {

	if root == nil {
		return
	}
	//PostorderRecursive(root.right)
	//PostorderRecursive(root.left)
	fmt.Printf("%d ", root.data)
}

func main() {
	//var tree I = Node{"0", nil, nil}

	//tree.NewNode(1)
	//fmt.Println("tree", tree)
	//var i I

	tree := NewNode(1)
	//i = AddNode(tree, 13)
	//	tree = AddNode(tree, 10)
	//	tree = AddNode(tree, 12)
	//	tree = AddNode(tree, 20)
	//	tree = AddNode(tree, 40)
	//	tree = AddNode(tree, 11)
	//	tree = AddNode(tree, 6)
	//	tree = AddNode(tree, 80)
	//	tree = AddNode(tree, 30)
	//	tree = AddNode(tree, 90)
	tree.Add(10)
	tree.Add(12)
	tree.Add(0)
	tree.Add(30)
	tree.Add(99)
	fmt.Println("tree: ", tree)
	//fmt.Println("tree: ", tree.right)
	InorderRecursive(tree)
	//PreorderRecursive(tree)
	//PostorderRecursive(tree)

}
