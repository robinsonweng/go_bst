package main

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
	assert.Equal(t, []string{"right"}, *node.path)
}

func TestSearchShouldReturnFalseWhenValueDontExists(t *testing.T) {
	node := Node{Val: 2, path: &[]string{}}
	node.Left = &Node{Val: 1}
	node.Right = &Node{Val: 3}

	result := Search(&node, 4)

	assert.Equal(t, false, result)
	assert.Equal(t, []string{"right"}, *node.path)
}

func TestInsertShouldReturnTrueWhenInsertSuccess(t *testing.T) {

}
