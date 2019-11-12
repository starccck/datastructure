package main

import (
	"github.com/starccck/datastructure/tree"
)

type Node struct {
	left *Node
	right *Node
	val interface{}
	height int
}

type AVLTreeI interface {
	tree.BtreeI
	Insert(root, node *Node) *Node
	Delete(root *Node, val interface{}) *Node
	leftLeftRotation(root *Node) *Node
	rightRightRotation(root *Node) *Node
	leftRightRotation(root *Node) *Node
	rightLeftRotation(root *Node) *Node
}