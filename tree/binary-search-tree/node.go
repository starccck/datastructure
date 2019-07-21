package main

type Node struct {
	left *Node
	right *Node
	val int
}

type BStreeI interface {
	Search(root *Node, val int) *Node
	Insert(root *Node, val int) *Node
	Delete(root *Node, val int)
	PreorderTraversal(root *Node)
	InorderTraversal(root *Node)
	PostorderTraversal(root *Node)
}