package main

import (
	"fmt"
)

type Node struct {
	key   int
	left  *Node
	right *Node
	size  int
}

func newNode(key int) *Node {
	return &Node{
		key:   key,
		left:  nil,
		right: nil,
		size:  1,
	}
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return newNode(key)
	}
	root.size++

	if key < root.key {
		root.left = insert(root.left, key)
	} else {
		root.right = insert(root.right, key)
	}

	return root
}

func remove(root *Node, key int) *Node {
	if root == nil {
		return root
	}

	root.size--

	if key < root.key {
		root.left = remove(root.left, key)
	} else if key > root.key {
		root.right = remove(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		root.key = minValueNode(root.right)
		root.right = remove(root.right, root.key)
	}

	return root
}

func minValueNode(node *Node) int {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.key
}

func kthMaximum(root *Node, k int) int {
	if root == nil || k > root.size {
		return -1 // or any other appropriate value
	}

	rightSize := 1
	if root.right != nil {
		rightSize += root.right.size
	}

	if k == rightSize {
		return root.key
	} else if k < rightSize {
		return kthMaximum(root.right, k)
	} else {
		return kthMaximum(root.left, k-rightSize)
	}
}

func processCommands(commands [][]int) []int {
	var result []int
	var root *Node

	for _, command := range commands {
		cmd, key := command[0], command[1]
		if cmd == 1 {
			root = insert(root, key)
		} else if cmd == 0 {
			result = append(result, kthMaximum(root, key))
		} else if cmd == -1 {
			root = remove(root, key)
		}
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	var commands [][]int
	for i := 0; i < n; i++ {
		var cmd, key int
		fmt.Scan(&cmd, &key)
		commands = append(commands, []int{cmd, key})
	}

	output := processCommands(commands)

	for _, value := range output {
		fmt.Println(value)
	}
}
