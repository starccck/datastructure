package main

import (
	"fmt"
	"testing"
)

func TestSearchLeftTreeMaxNode(t *testing.T) {
	root := &Node{
		val:10,
	}

	root.left = &Node{val:7,}
	root.right = &Node{val:15}
	root.left.left = &Node{val:2}
	root.left.right = &Node{val:8}
	root.right.right = &Node{val:20}

	node := searchLeftTreeMaxNode(root)
	fmt.Println(node)

	node = searchLeftTreeMaxNode(root.left)
	fmt.Println(node)

	node = searchLeftTreeMaxNode(root.left.left)
	fmt.Println(node)
}
