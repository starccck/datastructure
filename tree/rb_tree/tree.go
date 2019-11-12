package rb_tree

type COLOR int

const (
	RED COLOR = iota
	BLACK
)

type Node struct {
	value int
	left *Node
	right *Node
	parent *Node
	color COLOR
}

type Tree struct {
	root *Node
	size int
}

func (t *Tree)Insert(i *Node) {
	if t.size == 0 {
		i.color = BLACK
		t.root = i
		t.size++
		return
	}

	var n, last *Node = t.root, nil
	for n != nil {
		last = n
		if i.value < n.value {
			n = n.left
		} else if i.value > n.value {
			n = n.right
		} else {
			return
		}
	}

	if i.value < last.value {
		last.left = i
	} else {
		last.right = i
	}
	i.parent = last

	t.InsertFix(i)
	t.size++
}


// 插入修复函数
func (t *Tree) InsertFix(i *Node) {
	p := i.parent

	if p != nil && p.color == BLACK {
		return
	}

	for i.parent != nil && i.parent.color == RED {
		p = i.parent
		g := p.parent
		if p == g.left {
			u := g.right
			if u != nil && u.color == RED {
				p.color = BLACK
				u.color = BLACK
				g.color = RED
				i = g
			} else {
				if i == p.right {
					t.leftRotate(p)
					i, p = p, i
				}
				g.color = RED
				p.color = BLACK
				t.rightRotate(g)
			}
		} else {
			u := g.left
			if u != nil && u.color == RED {
				p.color = BLACK
				u.color = BLACK
				g.color = RED
				i = g
			} else {
				if i == p.left {
					t.rightRotate(p)
					i, p = p, i
				}
				g.color = RED
				p.color = BLACK
				t.leftRotate(g)
				i = p
			}
		}
		//p = i.parent
	}
	t.root.color = BLACK
}

func (t *Tree) Delete(v int) {
	n := t.root
	for n != nil && n.value != v {
		if v < n.value {
			n = n.left
		} else if v > n.value {
			n = n.right
		}
	}

	if n == nil {
		return
	}

	t.DeleteFix(n)
}

func (t *Tree) DeleteFix(n *Node) {
	for n != t.root {
		p := n.parent
		if n.left == nil && n.right == nil {
			if n.color == RED {
				deleteNode(n)
				return
			} else {
				b := getBrother(n)
				if b.left == nil && b.right == nil {
					b.color = RED
					p.color = BLACK
					deleteNode(n)
					return
				} else if b.left != nil && b.right != nil {
					if n == p.left {
						if b.color == BLACK {
							b.color = p.color
							b.right.color = BLACK
						} else {
							b.color = BLACK
							b.left.color = RED
						}
						p.color = BLACK
						deleteNode(n)

						t.leftRotate(p)
						return
					} else {
						if b.color == BLACK {
							b.color = p.color
							b.left.color = BLACK
						} else {
							b.color = BLACK
							b.right.color = RED
						}
						p.color = BLACK
						deleteNode(n)

						t.rightRotate(p)
						return
					}
				} else if b.left == nil {
					if b == p.left {
						// 左旋
						t.leftRotate(b)
						b = b.parent
					}
					// 右旋
					t.rightRotate(b.parent)
					b.color = p.color
					p.color = BLACK
					deleteNode(n)
				} else if b.right == nil {
					if b == p.right {
						// 右旋
						t.rightRotate(b)
						b = b.parent
					}
					// 左旋
					t.leftRotate(b.parent)
					b.color = p.color
					p.color = BLACK
					deleteNode(n)
				}
				return
			}
		} else if n.left != nil && n.right != nil {
			// 寻找后继节点， 交换位置
			// TODO
			rp := findNextNode(n.right)
			n.value, rp.value = rp.value, n.value
			n = rp
		} else {
			if n.left != nil {
				n.value = n.left.value
				n.left = nil
			} else {
				n.value = n.right.value
				n.right = nil
			}
			return
		}
	}
}

func deleteNode(n *Node) {
	if n == n.parent.left {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
}

func findNextNode(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}

	return n
}

func getBrother(n *Node) *Node {
	if n.parent == nil {
		return nil
	}

	if n == n.parent.left {
		return n.parent.right
	}
	return n.parent.left
}

// 右旋
func (t *Tree) rightRotate(n *Node) *Node {
	lc := n.left

	if lc == nil {
		return n
	}

	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = lc
		} else {
			n.parent.right = lc
		}
	}
	lc.parent = n.parent

	n.left = lc.right
	lc.right = n
	n.parent = lc
	if n.left != nil {
		n.left.parent = n
	}

	if lc.parent == nil {
		t.root = lc
	}

	return lc
}

// 左旋
func (t *Tree)leftRotate(n *Node) *Node {
	rc := n.right

	if rc == nil {
		return rc
	}

	if n.parent != nil {
		if n == n.parent.left {
			n.parent.left = rc
		} else {
			n.parent.right = rc
		}
	}
	rc.parent = n.parent

	n.right = rc.left
	rc.left = n
	n.parent = rc
	if n.right != nil {
		n.right.parent = n
	}

	if rc.parent == nil {
		t.root = rc
	}

	return rc
}

func (t *Tree) MiddleTraverse(n *Node, seq []interface{}) []interface{} {
	if n == nil {
		return seq
	}

	if n.left != nil {
		seq = t.MiddleTraverse(n.left, seq)
	}

	seq = append(seq, n)

	if n.right != nil {
		seq = t.MiddleTraverse(n.right, seq)
	}
	return seq
}