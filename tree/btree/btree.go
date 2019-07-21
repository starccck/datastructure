package main

import (
	"fmt"
)

func (tree *Node) PreorderTraversal(node *Node) {
	if node != nil {
		fmt.Printf("%v ", node.val)
	}
	
	tree.PreorderTraversal(node.left)
	tree.PreorderTraversal(node.right)
}

func (tree *Node) InorderTraversal(node *Node) {
	if node == nil {
		return
	}

	tree.InorderTraversal(node.left)
	fmt.Printf("%v ", node.val)
	tree.InorderTraversal(node.right)
}

func (tree *Node) PostorderTraversal(node *Node) {
	if node == nil {
		return
	}

	tree.PostorderTraversal(node.left)
	tree.PostorderTraversal(node.right)
	fmt.Printf("%v ", node.val)
}