package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchOneNodeShouldReturnTrueWhenValueExists(t *testing.T) {
	node := Node{Val: 1, path: &[]string{}}

	result := Search(&node, 1)

	assert.Equal(t, true, result)
	assert.Equal(t, []string{}, *node.path)
}

func TestSearchShouldReturnTrueWhenValueExists(t *testing.T) {
	node := Node{Val: 2, path: &[]string{}}
	node.Left = &Node{Val: 1}
	node.Right = &Node{Val: 3}

	result := Search(&node, 3)

	assert.Equal(t, true, result)
	assert.Equal(t, []string{"move right"}, *node.path)
}

func TestSearchShouldReturnFalseWhenValueDontExists(t *testing.T) {
	node := Node{Val: 2, path: &[]string{}}
	node.Left = &Node{Val: 1}
	node.Right = &Node{Val: 3}

	result := Search(&node, 4)

	assert.Equal(t, false, result)
	assert.Equal(t, []string{"move right"}, *node.path)
}

func TestInsertShouldReturnTrueWhenInsertSuccess(t *testing.T) {
	path := &[]string{}
	node := Node{Val: 2, path: path}
	node.Left = &Node{Val: 1, path: path}
	node.Right = &Node{Val: 3, path: path}

	result := Insert(&node, 4)

	assert.Equal(t, true, result)
	assert.Equal(t, []string{"move right", "insert right"}, *node.path)
}

func TestInsertShouldReturnFalseWhenInsertValueAlreadyExists(t *testing.T) {
	path := &[]string{}
	node := Node{Val: 2, path: path}

	result := Insert(&node, 2)

	assert.Equal(t, false, result)
	assert.Equal(t, []string{"value exists"}, *node.path)
}

func TestASCTraversalShouldReturnASCOrderValueWhenGivenBalancedBST(t *testing.T) {
	node := Node{Val: 100, path: &[]string{}}
	node.Left = &Node{Val: 20}
	node.Right = &Node{Val: 200}

	leftTree := node.Left
	leftTree.Parent = &node
	leftTree.Left = &Node{Val: 10, Parent: leftTree}
	leftTree.Right = &Node{Val: 30, Parent: leftTree}

	rightTree := node.Right
	rightTree.Parent = &node
	rightTree.Left = &Node{Val: 150}
	rightTree.Right = &Node{Val: 400}

	path := &[]int{}
	ASCTraversal(&node, path)

	assert.Equal(t, []int{10, 20, 30, 100, 150, 200, 400}, *path)
}

func TestDESCTraversalShouldReturnDESCOrderValueWhenGivenBalancedBST(t *testing.T) {
	node := Node{Val: 100, path: &[]string{}}
	node.Left = &Node{Val: 20}
	node.Right = &Node{Val: 200}

	leftTree := node.Left
	leftTree.Parent = &node
	leftTree.Left = &Node{Val: 10, Parent: leftTree}
	leftTree.Right = &Node{Val: 30, Parent: leftTree}

	rightTree := node.Right
	rightTree.Parent = &node
	rightTree.Left = &Node{Val: 150}
	rightTree.Right = &Node{Val: 400}

	path := &[]int{}
	DESCTraversal(&node, path)

	assert.Equal(t, []int{400, 200, 150, 100, 30, 20, 10}, *path)
}

func TestDeleteShouldReturnCorrectTreeWhenGivenBalancedBSTWithOneDepth(t *testing.T) {
	node := Node{Val: 100, path: &[]string{}}
	node.Left = &Node{Val: 20}
	node.Right = &Node{Val: 200}

	Delete(&node, 100)

	path := &[]int{}
	ASCTraversal(&node, path)

	assert.Equal(t, []int{20, 200}, *path)
}
