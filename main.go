package main

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
	path   *[]string
}

func Search(node *Node, value int) bool {
	if node.Parent == nil && node.Val == value {
		fmt.Println("find at root")
		return true
	}

	if node.Val == value {
		return true
	}

	if value < node.Val && node.Left != nil {
		path := *node.path
		path = append(path, "move left")
		node.path = &path
		node.Left.Parent = node
		return Search(node.Left, value)
	} else if value > node.Val && node.Right != nil {
		path := *node.path
		path = append(path, "move right")
		node.path = &path
		node.Right.Parent = node
		return Search(node.Right, value)
	}

	return false
}

func Insert(node *Node, value int) bool {
	if node.Parent == nil && node.Val == value {
		fmt.Println("find at root")
		return false
	}

	if node.Val == value {
		return false
	} else if value < node.Val {
		if node.Left != nil {
			// move left
			Record(node, "move left")
			return Insert(node.Left, value)
		} else {
			// if left is nil, insert left
			Record(node, "insert left")
			node.Left = &Node{Val: value, path: node.path}
			return true
		}
	} else if value > node.Val {
		if node.Right != nil {
			// move right
			Record(node, "move right")
			return Insert(node.Right, value)
		} else {
			// if right is nil, insert right
			Record(node, "insert right")
			node.Right = &Node{Val: value, path: node.path}
			return true
		}
	}
	return false
}

func Record(node *Node, operation string) {
	(*node.path) = append((*node.path), operation)
}

func (bst *Node) Delete() {}

type Dict interface {
	Search(value int) bool
	Insert(value int) bool
	Delete()
}

func main() {
}
