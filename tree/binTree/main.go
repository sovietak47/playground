package main

import (
	"playground/tree/binTree/binTreeFunc"
	"fmt"
)


/*
				A
		B				C
	D		E		F		G
		  J			  M
*/
func initBinTree() *binTreeFunc.TreeNode {
	root := &binTreeFunc.TreeNode{Value:"A"}
	root.Left = &binTreeFunc.TreeNode{Value:"B"}
	root.Right = &binTreeFunc.TreeNode{Value:"C"}
	root.Left.Left =  &binTreeFunc.TreeNode{Value:"D"}
	root.Left.Right =  &binTreeFunc.TreeNode{Value:"E"}
	root.Right.Left =  &binTreeFunc.TreeNode{Value:"F"}
	root.Right.Right =  &binTreeFunc.TreeNode{Value:"G"}
	root.Left.Right.Left =  &binTreeFunc.TreeNode{Value:"J"}
	root.Right.Left.Right =  &binTreeFunc.TreeNode{Value:"M"}

	return root
}


var printFunc = func(node *binTreeFunc.TreeNode) {
	fmt.Printf("%v ", node.Value)
}

func main() {
	root := initBinTree()

	binTreeFunc.PreOrderWalk(root, printFunc)
	fmt.Println()
	binTreeFunc.MidOrderWalk(root, printFunc)
	fmt.Println()
	binTreeFunc.PostOrderWalk(root, printFunc)
	fmt.Println()
	binTreeFunc.LevelOrderWalk(root, printFunc)




}
