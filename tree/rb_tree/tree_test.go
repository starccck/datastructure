package rb_tree

import (
	"fmt"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	nodes := []*Node{
		{
			value: 5,
		}, {
			value: 3,
		}, {
			value: 9,
		}, {
			value: 78,
		}, {
			value: 1,
		}, {
			value: 56,
		}, {
			value: 2,
		}, {
			value: 89,
		}, {
			value: -87,
		}, {
			value: -200,
		},
	}
	tree := &Tree{}
	for _, v := range nodes {
		tree.Insert(v)
	}

	tree.Delete(2)

	//tree.leftRotate(tree.root)

	seq := make([]interface{}, 0)
	seq = tree.MiddleTraverse(tree.root, seq)
	for _, v := range seq {
		fmt.Println(v.(*Node).value, v.(*Node).color)
	}
}
