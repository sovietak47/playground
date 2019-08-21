package main

import (
	"fmt"
	"playground/sort/sortFunc"
)

func main() {

	li := []int{0, 3, 2, 20, 1, 10, 5, 7, 111, 10, 14, 9, 35}

	//sortFunc.BubbleSort(li)
	//sortFunc.SelectSort(li)
	//sortFunc.InsertSort(li)
	//sortFunc.ShellSort(li)
	//sortFunc.QuickSort(li,0,len(li)-1)
	//sortFunc.MergeSort(li, 0, len(li)-1)
	//sortFunc.Heapsort(li)
	//sortFunc.CountSort(li)
	//sortFunc.BucketSort(li,10)
	sortFunc.RadixSort(li)
	fmt.Println(li)

}
