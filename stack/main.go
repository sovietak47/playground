package main

import (
	"playground/stack/modle"
	"fmt"
)

func main() {
	init_test()
	empty_test()
	push_pop_test()
	exist()
}

func init_test() {
	stack := new(modle.Stack)
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Length())
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println()
}

func empty_test() {
	stack := new(modle.Stack)
	stack.Push(1)
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Length())
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println()

	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Length())
	fmt.Println(stack.Top())
	fmt.Println()
}

func push_pop_test() {
	stack := new(modle.Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Length())
	fmt.Println(stack.Top())
	fmt.Println(stack)
	fmt.Println()

	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println()
}

func exist() {
	type subStru struct {
		num float64
		str string
		bol bool
		dict map[string]string
	}

	type stru struct {
		num float64
		str string
		bol bool
		dict map[string]stru
		stru *stru
		substru subStru
	}

	testStru := stru{
		num:1.2,
		str:"aaa",
		bol:true,
		dict:map[string]stru{"bbb":stru{num:3.4,str:"bbb"}},
		stru:&stru{num:5.6,str:"ccc"},
		substru:subStru{
			num:7.8,
			str:"ddd",
			bol:true,
			dict:map[string]string{"eee":"fff"},
		},
	}
	stack := new(modle.Stack)
	stack.Push(testStru)
	fmt.Println(stack.Exist(testStru))
	fmt.Println()
}