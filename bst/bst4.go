package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	value int
	right *Node
	left  *Node
}

func newNode(value int) *Node {
	return &Node{
		value: value,
		right: nil,
		left:  nil,
	}
}

type Tree struct {
	root *Node
}

func newTree() *Tree {
	return &Tree{
		root: nil,
	}
}

func (t *Tree) insert(x int) {
	if t.root == nil {
		t.root = newNode(x)
	} else {
		t._insertRecursive(t.root, x)
	}
}

func (t *Tree) _insertRecursive(node *Node, x int) {
	if x < node.value {
		if node.left == nil {
			node.left = newNode(x)
		} else {
			t._insertRecursive(node.left, x)
		}
	} else if x > node.value {
		if node.right == nil {
			node.right = newNode(x)
		} else {
			t._insertRecursive(node.right, x)
		}
	}
}

func (t *Tree) getRightmostValues(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			element := queue.Front()
			queue.Remove(element)
			node := element.Value.(*Node)

			if i == levelSize-1 {
				result = append(result, node.value)
			}

			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
	}

	return result
}

func main() {
	tree := newTree()

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		tree.insert(value)
	}

	result := tree.getRightmostValues(tree.root)

	for i := 0; i < len(result); i++ {
		fmt.Print(result[i], " ")
	}
}
