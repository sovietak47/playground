package sortFunc

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func BubbleSort(li []int) {
	for i := 0; i < len(li); i++ {
		for j := 0; j < len(li)-1-i; j++ {
			if li[j] > li[j+1] {
				li[j], li[j+1] = li[j+1], li[j]
			}
		}
	}
}

func SelectSort(li []int) {
	for i := 0; i < len(li); i++ {
		pos := i
		for j := i; j < len(li); j++ {
			if li[j] < li[pos] {
				pos = j
			}
		}
		li[i], li[pos] = li[pos], li[i]
	}
}

func InsertSort(li []int) {
	for i := 1; i < len(li); i++ {
		insVal := li[i]
		j := i - 1
		for j >= 0 && li[j] > insVal {
			li[j+1] = li[j]
			j--
		}
		li[j+1] = insVal
	}
}

func ShellSort(li []int) {
	for gap := len(li) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(li); i = i + gap {
			insVal := li[i]
			j := i - gap
			for j >= 0 && li[j] > insVal {
				li[j+gap] = li[j]
				j = j - gap
			}
			li[j+gap] = insVal
		}
	}
}

func QuickSort(li []int, left, right int) {
	if left >= right { //等于号主要是为了right-left 一定大于0
		return
	}

	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left

	li[i], li[r] = li[r], li[i]

	for i < j {
		for i < j && li[j] >= li[i] { //这个 i < j  不可被外层的i < j 替代，这个检测更频繁，而且更为重要
			j--
		}
		li[i], li[j] = li[j], li[i]
		for i < j && li[i] <= li[j] {
			i++
		}
		li[i], li[j] = li[j], li[i]
	}

	QuickSort(li, left, i-1)
	QuickSort(li, i+1, right)
}

func MergeSort(li []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	MergeSort(li, left, mid)
	MergeSort(li, mid+1, right)
	merge(li, left, mid, right)
}

func merge(li []int, left, mid, right int) {
	i := left
	j := mid + 1

	tmp := []int{}

	for i <= mid && j <= right {
		if li[i] <= li[j] {
			tmp = append(tmp, li[i])
			i++
		} else {
			tmp = append(tmp, li[j])
			j++
		}
	}

	if i <= mid {
		tmp = append(tmp, li[i:mid+1]...)
	}

	if j <= right {
		tmp = append(tmp, li[j:right+1]...)
	}

	for k := 0; k < len(tmp); k++ {
		li[left+k] = tmp[k]
	}
}

func Heapsort(li []int) {
	for i := len(li)/2 - 1; i >= 0; i-- {
		heapShift(li, i, len(li)-1)
	}

	for j := len(li) - 1; j > 0; j-- {
		li[j], li[0] = li[0], li[j]
		heapShift(li, 0, j-1)
	}
}

func heapShift(li []int, begin int, end int) {
	i := begin
	j := 2*i + 1

	for j <= end {
		if li[j+1] >= li[j] && j+1 <= end {
			j = j + 1
		}

		if li[j] >= li[i] {
			li[i], li[j] = li[j], li[i]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
}

func CountSort(li []int) {
	max := li[0]
	for i := 0; i < len(li); i++ {
		if li[i] > max {
			max = li[i]
		}
	}

	counter := make([]int, max+1)
	for j := 0; j < len(li); j++ {
		counter[li[j]]++
	}

	n := 0
	for k, v := range counter {
		for m := 0; m < v; m++ {
			li[n] = k
			n++
		}
	}
}

func BucketSort(li []int, bucketNum int) {
	max := li[0]
	min := li[0]
	for i := 0; i < len(li); i++ {
		if li[i] > max {
			max = li[i]
		}
		if li[i] < min {
			min = li[i]
		}
	}

	buckets := make([][]int, bucketNum)
	for j := 0; j < len(li); j++ {
		index := (li[j] - min) / ((max - min) / (bucketNum - 1))
		buckets[index] = append(buckets[index], li[j])
		// 以插入排序为例
		k := len(buckets[index]) - 1 - 1
		for k >= 0 && buckets[index][k] > li[j] {
			buckets[index][k+1] = buckets[index][k]
			k--
		}
		buckets[index][k+1] = li[j]
	}

	n := 0
	for _, v1 := range buckets {
		for _, v2 := range v1 {
			li[n] = v2
			n++
		}
	}

}

func RadixSort(li []int) {
	max := li[0]
	for i := 0; i < len(li); i++ {
		if li[i] > max {
			max = li[i]
		}
	}

	for j := 0; j < len(strconv.Itoa(max)); j++ {
		buckets := make([][]int, 10)
		for k := 0; k < len(li); k++ {
			index := li[k] / int(math.Pow10(j)) % 10
			buckets[index] = append(buckets[index], li[k])
		}

		m := 0
		for _, v := range buckets {
			for _, v2 := range v {
				li[m] = v2
				m++
			}
		}
	}
}
