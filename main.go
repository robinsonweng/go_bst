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
		path = append(path, "left")
		node.path = &path
		node.Left.Parent = node
		return Search(node.Left, value)
	} else if value > node.Val && node.Right != nil {
		path := *node.path
		path = append(path, "right")
		node.path = &path
		node.Right.Parent = node
		return Search(node.Right, value)
	}

	return false
}

func (bst *Node) Insert(node *Node, value int) bool {
	return false
}

func (bst *Node) Delete() {}

type Dict interface {
	Search(value int) bool
	Insert(value int) bool
	Delete()
}

func main() {
}
