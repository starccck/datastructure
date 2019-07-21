package main

import (
	"fmt"
)

func (tree *Node)Search(root *Node, val int) *Node {
	for root != nil {
		if val < root.val {
			root = root.left
		} else if val > root.val {
			root = root.right
		} else {
			return root
		}
	}

	return nil
}

func (tree *Node) Insert(node *Node, toInsert *Node) *Node {
	if node == nil {
		return toInsert
	}


	root := node
	for node != nil {
		if toInsert.val < node.val {
			if node.left == nil {
				node.left = toInsert
				return root
			}
			node = node.left
		} else if node.val < toInsert.val {
			if node.right == nil {
				node.right = toInsert
				return root
			}
			node = node.right
		} else {
			return root
		}
	}

	return root
}

func (tree *Node) Delete(node *Node, toDelete int) {
	var lastNode *Node
	for node != nil {
		if toDelete < node.val {
			lastNode = node
			node = lastNode.left
		} else if toDelete > node.val {
			lastNode = node
			node = lastNode.right
		} else {	
			// 左边子树为空
			if node.left == nil {
				if node == lastNode.left {
					lastNode.left = node.right
				} else {
					lastNode.right = node.right
				}
				return
			}

			// 查找左子树的最大节点, 找到后删除
			//pNode := lastNode
			toDelete := node
			node = node.left
			for node.right != nil {
				lastNode = node
				node = lastNode.right
			}
			toDelete.val = node.val
			lastNode.right = nil
		}
	}
}

func (tree *Node) PreorderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.val)
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