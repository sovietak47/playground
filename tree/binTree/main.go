package main

import (
	"fmt"
	"playground/tree/binTree/binTreeFunc"
)

/*
				A
		B				C
	D		E		F		G
		  J			  M
*/
func initBinTree() *binTreeFunc.TreeNode {
	root := &binTreeFunc.TreeNode{Value: "A"}
	root.Left = &binTreeFunc.TreeNode{Value: "B"}
	root.Right = &binTreeFunc.TreeNode{Value: "C"}
	root.Left.Left = &binTreeFunc.TreeNode{Value: "D"}
	root.Left.Right = &binTreeFunc.TreeNode{Value: "E"}
	root.Right.Left = &binTreeFunc.TreeNode{Value: "F"}
	root.Right.Right = &binTreeFunc.TreeNode{Value: "G"}
	root.Left.Right.Left = &binTreeFunc.TreeNode{Value: "J"}
	root.Right.Left.Right = &binTreeFunc.TreeNode{Value: "M"}

	return root
}

var printFunc = func(node *binTreeFunc.TreeNode) {
	fmt.Printf("%v ", node.Value)
}

func main() {
	root := initBinTree()

	root.PreOrderWalk(printFunc)
	fmt.Println()
	root.MidOrderWalk(printFunc)
	fmt.Println()
	root.PostOrderWalk(printFunc)
	fmt.Println()
	root.LevelOrderWalk(printFunc)
	fmt.Println()
	fmt.Println()

	binTreeFunc.PreOrderWalk2(root, printFunc)
	fmt.Println()
	binTreeFunc.MidOrderWalk2(root, printFunc)
	fmt.Println()
	binTreeFunc.PostOrderWalk2(root, printFunc)
	fmt.Println()
	fmt.Println()

	fmt.Println(root.Count())
	fmt.Println(root.Height())
	fmt.Println(root.IsEqual(root))
}
