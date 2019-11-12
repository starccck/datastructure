package main

func (tree *Node) LeftLeftRotation(root *Node) *Node {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	newRoot.height = max(newRoot.left, newRoot.right)
	return newRoot
}

func (tree *Node) RightRightRotation(root *Node) *Node {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	newRoot.height = max(newRoot.left, newRoot.right)
	return newRoot
}

func (tree *Node) LeftRightRotation(root *Node) *Node {
	newNode := root.left.left.right
	newNode.left = root.left.left
	newNode.left.height = 1
	newNode.height = 2
	root.left.left = newNode

	newNode = root
	root = root.left
	newNode.left = root.right
	root.right = newNode
	newNode.height = max(newNode.left, newNode.right)
	root.height = max(root.left, root.right)

	return root
}

func (tree *Node) rightLeftRotation(root *Node) *Node {
	newNode := root.right.right.left
	newNode.right = root.right.right
	root.right.right = newNode
	newNode.height = 2
	newNode.right.height = 1

	newNode = root
	root = root.right
	newNode.right = root.left
	root.left = newNode
	newNode.height = max(newNode.left, newNode.right)
	root.height = max(root.left, root.right)

	return root
}

func adjustHeight(root *Node) int {
	if root.left == nil && root.right == nil {
		root.height = 0
		return 0
	} else if root.left == nil {
		root.height = root.right.height + 1
	} else if root.right == nil {
		root.height = root.left.height + 1
	} else if root.left.height > root.right.height {
		root.height = root.left.height + 1
	} else {
		root.height = root.right.height + 1
	}

	return root.height
}

func (tree *Node) Insert(root, node *Node) *Node {
	if root == nil {
		return node
	}

	var newRoot *Node

	if node.val.(int) < root.left.val.(int) {
		root.left = tree.Insert(root.left, node)
		root.height = max(root.left, root.right)
	} else if node.val.(int) > root.right.val.(int) {
		root.right = tree.Insert(root.right, node)
		root.height = max(root.left, root.right)
	}
	newRoot = root

	if newRoot.left.height - newRoot.right.height == 2 {
		child := newRoot.left
		if child.left.height < child.right.height {
			newRoot = newRoot.LeftRightRotation(root)
		} else {
			newRoot = newRoot.LeftLeftRotation(root)
		}
	} else if newRoot.right.height - newRoot.left.height == 2 {
		child := newRoot.right
		if child.left.height > child.right.height {
			newRoot = newRoot.rightLeftRotation(newRoot)
		} else {
			newRoot = newRoot.RightRightRotation(newRoot)
		}
	}

	//newRoot.height = max(newRoot.left, newRoot.right)
	return newRoot
}

func (tree *Node) deletedNode(root, deleteNode *Node) *Node {
	if deleteNode.left == nil && deleteNode.right == nil {
		if root.left == deleteNode {
			root.left = nil
		} else if root.right == deleteNode {
			root.right = nil
		}
		return root
	} else if deleteNode.left == nil {
		if root.left == deleteNode {
			root.left = nil
		} else if root.right == deleteNode {
			root.right = nil
		}
		return root
	} else if deleteNode.right == nil {
		if root.left == deleteNode {
			root.left = nil
		} else if root.right == deleteNode {
			root.right = nil
		}
		return root
	} else {
		newNode := searchLeftTreeMaxNode(root)
		newNode.left = root.left
		newNode.right = root.right
		return newNode
	}
	return nil
}

func searchLeftTreeMaxNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.right == nil {
		return root
	}

	node := searchRightNode(root)
	root.height = max(root.left, root.right)
	return node
}

func searchRightNode(root *Node) *Node {
	if root.right.right != nil {
		newNode := searchRightNode(root.right)
		root.height = max(root.left, root.right)
		return newNode
	}
	newNode := root.right
	root.right = newNode.left
	root.height = max(root.left, root.right)
	return newNode
}

func (tree *Node) Delete(root *Node, val interface{}) *Node {
	if root == nil {
		return nil
	}

	var newRoot *Node
	if val.(int) == root.val.(int) {

	} else if val.(int) < root.val.(int) {
		newRoot = tree.Delete(root.left, val)
	} else if val.(int) > root.val.(int) {
		newRoot = tree.Delete(root.right, val)
	}

	return newRoot
}

func max(l, r *Node) int {
	if l == nil && r == nil {
		return 0
	} else if l == nil {
		return r.height + 1
	} else if r == nil {
		return l.height + 1
	}

	if l.height > r.height {
		return l.height + 1
	}

	return r.height + 1
}