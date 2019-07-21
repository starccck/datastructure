package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []int{8, 2, 9, 7, 5, 4, -9, 57}
	var root = &Node{}
	root = root.Insert(nil, &Node{val: 10})

	for _, val := range tests {
		root = root.Insert(root, &Node{val: val})
	}

	root.PreorderTraversal(root)
}