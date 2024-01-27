package main

import (
	"fmt"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{
		key:   value,
		left:  nil,
		right: nil,
	}
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{
		root: nil,
	}
}

func (bst *BST) insert(node *Node, value int) *Node {
	if node == nil {
		return NewNode(value)
	}
	if value < node.key {
		node.left = bst.insert(node.left, value)
	} else if value > node.key {
		node.right = bst.insert(node.right, value)
	}
	return node
}

func (bst *BST) InOrderTraversal(node *Node) {
	if node != nil {
		bst.InOrderTraversal(node.left)
		fmt.Printf("%d ", node.key)
		bst.InOrderTraversal(node.right)
	}
}

func (bst *BST) remove(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.key {
		node.left = bst.remove(node.left, value)
	} else if value > node.key {
		node.right = bst.remove(node.right, value)
	} else {
		if node.left == nil {
			temp := node.right
			node = nil
			return temp
		} else if node.right == nil {
			temp := node.left
			node = nil
			return temp
		}
		temp := bst.findMin(node.right)
		node.key = temp.key
		node.right = bst.remove(node.right, temp.key)
	}
	return node
}

func (bst *BST) findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (bst *BST) next(root *Node, key int) *Node {
	cur := root
	successor := (*Node)(nil)
	for cur != nil {
		if cur.key > key {
			successor = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return successor
}

func (bst *BST) prev(root *Node, key int) *Node {
	cur := root
	successor := (*Node)(nil)
	for cur != nil {
		if cur.key < key {
			successor = cur
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return successor
}

func (bst *BST) search(x *Node, key int) *Node {
	if x == nil || key == x.key {
		return x
	}
	if key < x.key {
		return bst.search(x.left, key)
	} else {
		return bst.search(x.right, key)
	}
}

func (bst *BST) Insert(value int) {
	bst.root = bst.insert(bst.root, value)
}

func (bst *BST) Remove(value int) {
	bst.root = bst.remove(bst.root, value)
}

func (bst *BST) Exist(value int) {
	res := bst.search(bst.root, value)
	if res == nil {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}

func (bst *BST) Next(value int) {
	res := bst.next(bst.root, value)
	if res == nil {
		fmt.Println("none")
	} else {
		fmt.Println(res.key)
	}
}

func (bst *BST) Prev(value int) {
	res := bst.prev(bst.root, value)
	if res == nil {
		fmt.Println("none")
	} else {
		fmt.Println(res.key)
	}
}

func main() {
	bst := NewBST()
	var operation string
	for {
		_, err := fmt.Scan(&operation)
		if err != nil {
			break
		}
		if operation == "insert" {
			var x int
			fmt.Scan(&x)
			bst.Insert(x)
		} else if operation == "delete" {
			var x int
			fmt.Scan(&x)
			bst.Remove(x)
		} else if operation == "exists" {
			var x int
			fmt.Scan(&x)
			bst.Exist(x)
		} else if operation == "next" {
			var x int
			fmt.Scan(&x)
			bst.Next(x)
		} else if operation == "prev" {
			var x int
			fmt.Scan(&x)
			bst.Prev(x)
		}
	}
}
