package tree

type Node struct {
	left *Node
	right *Node
	val interface{}
}

type BtreeI interface {
	PreorderTraversal(root *Node)
	InorderTraversal(root *Node)
	PostorderTraversal(root *Node)
}