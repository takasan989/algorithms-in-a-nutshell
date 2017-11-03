package red_black_tree

import (
	"fmt"
)

type color int

const (
	BLACK color = iota
	RED
)

type Tree struct {
	root *Node
}

type Node struct {
	tree   *Tree
	color  color
	parent *Node
	left   *Node
	right  *Node
	value  int
}

func Search(tree *Tree, value int) bool {
	if tree.root == nil {
		return false
	}
	return search(tree.root, value)
}

func search(node *Node, value int) bool {
    if value == node.value {
        return true
    }
    
    if value < node.value {
        if node.left == nil {
            return false
        }
        return search(node.left, value)
    }
    
    if node.right == nil {
        return false
    }
    return search(node.right, value)
}

func NewTree() *Tree {
	tree := Tree{}
	return &tree
}

func NewNode() *Node {
	node := Node{}
	return &node
}

func (node *Node) grand() *Node {
	parent := node.parent
	
	if parent == nil {
		return nil
	}
	
	return parent.parent
}

func (node *Node) uncle() *Node {
	grand := node.grand()
	
	if grand == nil {
		return nil
	}
	
	if grand.left == node.parent {
		return grand.right
	} else {
		return grand.left
	}
}

func (node *Node) balance() {
	parent := node.parent
	grand := node.grand()
	uncle := node.uncle()
	
	if node.tree != nil || parent.color == BLACK {
		node.balanceCase1()
		return
	}
	
	if uncle != nil && uncle.color == RED {
		node.balanceCase2()
		return
	}
	
	if (grand.left == parent && parent.right == node) || (grand.right == parent && parent.left == node) {
		node.balanceCase3()
		return
	}
	
	if (grand.left == parent && parent.left == node) || (grand.left == parent && parent.left == node) {
		node.balanceCase4()
		return
	}
	
	fmt.Errorf("invalid state")
}

func (node *Node) balanceCase4() {
	parent := node.parent
	grand := node.grand()
	
	if grand.left == parent {
		grand.rotateRight()
	} else {
		grand.rotateLeft()
	}
	
	grand.color, parent.color = parent.color, grand.color
}

func (node *Node) balanceCase3() {
	parent := node.parent
	grand := node.grand()
	
	if grand.left == parent {
		parent.rotateLeft()
		parent.balance()
	} else {
		parent.rotateRight()
		parent.balance()
	}
}

func (node *Node) balanceCase2() {
	parent := node.parent
	grand := node.grand()
	uncle := node.uncle()
	
	parent.color = BLACK
	uncle.color = BLACK
	
	grand.color = RED
	grand.balance()
}

func (node *Node) balanceCase1() {
	if node.tree != nil {
		node.color = BLACK
		return
	}
	
	node.color = RED
}

func (tree *Tree) Append(value int) {
	if tree.root == nil {
		tree.root = &Node{value: value, tree: tree}
		return
	}
	
	node := tree.root.append(&Node{value: value, color: RED})
	node.balance()
}

func (tree *Tree) append(node *Node) *Node {
	if tree.root == nil {
		tree.root = node
		node.tree = tree
		return node
	}
	
	return tree.root.append(node)
}

func (node *Node) Append(value int) {
	if value < node.value {
		if node.left != nil {
			node.left.Append(value)
			return
		}
		node.left = &Node{value: value, parent: node}
		return
	}
	
	if node.right != nil {
		node.right.Append(value)
		return
	}
	node.right = &Node{value: value, parent: node}
}

func (parent *Node) append(node *Node) *Node {
	if node.value < parent.value {
		if parent.left != nil {
			return parent.left.append(node)
		}
		parent.left = node
		node.parent = parent
		node.balance()
		return node
	}
	
	if parent.right != nil {
		return parent.right.append(node)
	}
	
	parent.right = node
	node.parent = parent
	node.balance()
	return node
}

func (node *Node) rotateLeft() {
	if node.tree != nil {
		right := node.right
		node.right = right.left
		if right.left != nil {
			right.left.parent = node
		}
		
		right.tree = node.tree
		right.tree.root = right
		right.parent = nil
		node.tree = nil
		
		right.left = node
		node.parent = right
		return
	}
	
	right := node.right
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	
	right.parent = node.parent
	if node.parent.left == node {
		node.parent.left = right
	} else {
		node.parent.right = right
	}
	
	right.left = node
	node.parent = right
}

func (node *Node) rotateRight() {
	if node.tree != nil {
		left := node.left
		node.left = left.right
		if left.right != nil {
			left.right.parent = node
		}
		
		left.tree = node.tree
		left.tree.root = left
		left.parent = nil
		node.tree = nil
		
		left.right = node
		node.parent = left
		return
	}
	
	left := node.left
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	
	left.parent = node.parent
	if node.parent.left == node {
		node.parent.left = left
	} else {
		node.parent.right = left
	}
	
	left.right = node
	node.parent = left
}

func (tree *Tree) String() string {
	return fmt.Sprintf("Tree{%s}", tree.root.String())
}

func (node *Node) String() string {
	value := node.value
	
	left := ""
	if node.left != nil {
		left = fmt.Sprintf(",left: %s", node.left)
	}
	
	right := ""
	if node.right != nil {
		right = fmt.Sprintf(",right: %s", node.right)
	}
	
	return fmt.Sprintf("Node{%d%s%s}", value, left, right)
}

func (tree *Tree) StringDetail() string {
	return fmt.Sprintf("Tree{%s}", tree.root.StringDetail())
}

func (node *Node) StringDetail() string {
	value := node.value
	
	color := "B"
	if node.color == RED {
		color = "R"
	}
	
	left := ""
	if node.left != nil {
		left = fmt.Sprintf(",left: %s", node.left.StringDetail())
	}
	
	right := ""
	if node.right != nil {
		right = fmt.Sprintf(",right: %s", node.right.StringDetail())
	}
	
	return fmt.Sprintf("Node{%s%d%s%s}", color, value, left, right)
}