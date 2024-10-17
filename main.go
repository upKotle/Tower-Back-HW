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

func (bst *BST) Add(value int) {
	bst.root = addRec(bst.root, value)
}

func addRec(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}
	if value < node.value {
		node.left = addRec(node.left, value)
	} else if value > node.value {
		node.right = addRec(node.right, value)
	}
	return node
}

func isExist(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	if value < node.value {
		return isExist(node.left, value)
	}
	return isExist(node.right, value)
}

func (bst *BST) Search(value int) bool {
	return searchRec(bst.root, value)
}

func searchRec(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	if value < node.value {
		return searchRec(node.left, value)
	}
	return searchRec(node.right, value)
}

func (bst *BST) Delete(value int) {
	bst.root = deleteRec(bst.root, value)
}

func deleteRec(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = deleteRec(node.left, value)
	} else if value > node.value {
		node.right = deleteRec(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		minNode := findMin(node.right)
		node.value = minNode.value
		node.right = deleteRec(node.right, minNode.value)
	}
	return node
}

func findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func main() {
	bst := &BST{}
	bst.Add(10)
	bst.Add(20)
	if isExist(bst.root, 20) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	if isExist(bst.root, 30) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	bst.Delete(20)
	if isExist(bst.root, 20) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
