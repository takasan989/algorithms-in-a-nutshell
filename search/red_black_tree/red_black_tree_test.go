package red_black_tree

import (
	"testing"
)

func TestBalanceCase2(t *testing.T) {
	tree := &Tree{}
	g := tree.append(&Node{value:100, color: BLACK})
	p := tree.append(&Node{value:50, color: RED})
	u := tree.append(&Node{value:150, color: RED})
	n := tree.append(&Node{value:25, color: RED})
	
	n.balanceCase2()
	
	if p.color != BLACK || u.color != BLACK || g.color != BLACK {
		t.Errorf("balanceCase2 %v", p.color)
		t.Errorf("balanceCase2 %v", u.color)
		t.Errorf("balanceCase2 %v", g.color)
	}
}

func TestBalanceCase1(t* testing.T) {
	tree := &Tree{}
	node := tree.append(&Node{value:10, color: RED})
	
	node.balanceCase1()
	
	if node.color != BLACK {
		t.Errorf("balanceCase1 %v", node)
	}
	
	node = tree.append(&Node{value:5, color: RED})
	
	node.balanceCase1()
	
	if node.color != RED {
		t.Errorf("balanceCase1 %v", node)
	}
}

func TestRotateLeft(t *testing.T) {
	tree := &Tree{}
	tree.Append(3)
	tree.Append(2)
	tree.Append(5)
	tree.Append(4)
	tree.Append(6)
	tree.root.rotateLeft()
	
	actual := tree.String()
	expected := "Tree{Node{5,left: Node{3,left: Node{2},right: Node{4}},right: Node{6}}}"
	
	if actual != expected {
		t.Errorf("rotateLeft %v", actual)
	}
}

func TestRotateRight(t *testing.T) {
	tree := NewTree()
	tree.Append(5)
	tree.Append(6)
	tree.Append(3)
	tree.Append(2)
	tree.Append(4)
	tree.root.rotateRight()
	
	actual := tree.String()
	expected := "Tree{Node{3,left: Node{2},right: Node{5,left: Node{4},right: Node{6}}}}"
	
	if actual != expected {
		t.Errorf("rotateRight %v", actual)
	}
}

func TestAppendToNewTree(t *testing.T) {
	tree := NewTree()
	tree.Append(1)
	
	if tree.root.value != 1 {
		t.Errorf("Append %v", tree)
	}
}

func TestAppendToTree(t *testing.T) {
	tree := NewTree()
	tree.Append(2)
	tree.Append(1)
	tree.Append(3)
	
	if tree.root.left.value != 1 || tree.root.right.value != 3 {
		t.Errorf("Append %v", tree)
	}
}

func TestNewTree(t *testing.T) {
	tree := NewTree()
	
	if tree.root != nil {
		t.Errorf("NewTree %v", tree)
	}
}

func TestNewNode(t *testing.T) {
	root := NewNode()

	if root.color != BLACK || root.left != nil || root.right != nil {
		t.Errorf("NewNode error %v", root)
	}
}
