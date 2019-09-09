package binTreeFunc

import (
	"math"
	"reflect"
)

type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Count() int {
	if node != nil {
		return node.Left.Count()+node.Right.Count()+1
	}
	return 0
}

func  (node *TreeNode) Height() int {
	if node != nil {
		leftHeight := node.Left.Height()
		rightHeight := node.Right.Height()
		return int(math.Max(float64(leftHeight),float64(rightHeight))) +1
	}
	return 0
}

func (node *TreeNode) IsEqual(another *TreeNode) bool {
	if node == nil && another == nil {
		return true
	}
	if node != nil && another != nil && reflect.DeepEqual(node.Value, node.Value) {
		if node.Left.IsEqual(another.Left) {
			if node.Right.IsEqual(another.Right) {
				return true
			}
		}
	}
	return false
}