package findclosestvalueinbst

import (
	"testing"
)

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func TestCase1(t *testing.T) {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)
	actual := root.FindClosestValue(12)

	if actual != 13 {
		t.Errorf("got %v want %v", actual, 13)
	}
}
