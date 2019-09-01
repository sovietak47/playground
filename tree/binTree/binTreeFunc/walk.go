package binTreeFunc

import (
	"playground/stack/modle"
)

func PreOrderWalk(node *TreeNode, walkFunc func(node *TreeNode)) {
	if node == nil {
		return
	}

	if walkFunc != nil {
		walkFunc(node)
	}

	PreOrderWalk(node.Left, walkFunc)
	PreOrderWalk(node.Right, walkFunc)
}

func MidOrderWalk(node *TreeNode, walkFunc func(node *TreeNode)) {
	if node == nil {
		return
	}

	MidOrderWalk(node.Left, walkFunc)

	if walkFunc != nil {
		walkFunc(node)
	}

	MidOrderWalk(node.Right, walkFunc)
}

func PostOrderWalk(node *TreeNode, walkFunc func(node *TreeNode)) {
	if node == nil {
		return
	}

	PostOrderWalk(node.Left, walkFunc)
	PostOrderWalk(node.Right, walkFunc)

	if walkFunc != nil {
		walkFunc(node)
	}
}

func LevelOrderWalk(root *TreeNode, walkFunc func(node *TreeNode)) {
	levelWalk([]*TreeNode{root}, walkFunc)
}

func levelWalk(nodes []*TreeNode, walkFunc func(node *TreeNode)) {
	if len(nodes) == 0 {
		return
	}

	for _, v := range nodes {
		if walkFunc != nil {
			walkFunc(v)
		}
	}

	nextLevel := getNextLevelNode(nodes)
	levelWalk(nextLevel, walkFunc)
}

func getNextLevelNode(nodes []*TreeNode) []*TreeNode {
	ret := []*TreeNode{}

	for _, v := range nodes {
		if v.Left != nil {
			ret = append(ret, v.Left)
		}
		if v.Right != nil {
			ret = append(ret, v.Right)
		}
	}

	return ret
}


func PreOrderWalk2(node *TreeNode, walkFunc func(node *TreeNode)) {
	if node == nil {
		return
	}

	stack := new(modle.Stack)
	t := node

	for t != nil || !stack.IsEmpty() {
		for t != nil {
			walkFunc(t)
			stack.Push(t)
			t = t.Left
		}

		if !stack.IsEmpty() {
			v, _ := stack.Pop()
			t = v.(*TreeNode)
			t = t.Right
		}
	}
	return
}


func MidOrderWalk2(node *TreeNode, walkFunc func(node *TreeNode)) {
	if node == nil {
		return
	}

	stack := new(modle.Stack)
	t := node

	for t != nil || !stack.IsEmpty() {
		for t != nil {
			stack.Push(t)
			t = t.Left
		}

		if !stack.IsEmpty() {
			v, _ := stack.Pop()
			t = v.(*TreeNode)
			walkFunc(t)
			t = t.Right
		}
	}
	return
}

func PostOrderWalk2(node *TreeNode, walkFunc func(node *TreeNode)) {
	if node == nil {
		return
	}

	stack:= new(modle.Stack)
	walkedRightNodes := new(modle.Stack)
	t := node

	for t != nil || !stack.IsEmpty() {
		for t != nil {
			stack.Push(t)
			t = t.Left
		}

		if !stack.IsEmpty() {
			pre := stack.Top().(*TreeNode)
			if pre.Right != nil && !walkedRightNodes.Exist(pre){
				t = pre.Right
				walkedRightNodes.Push(pre)
			} else {
				p, _ := stack.Pop()
				t = p.(*TreeNode)
				walkFunc(t)
				t = nil
			}
		}
	}
}