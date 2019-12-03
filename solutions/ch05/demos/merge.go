package main

import (
	"fmt"
	"sort"
)

// merge two sorted sub array
func merge(a, a1, a2 []int) {
	tmp := make([]int, len(a))
	i, j, k := 0, 0, 0

	for i < len(a1) && j < len(a2) {
		minor := 0
		if a1[i] < a2[j] {
			minor = a1[i]
			i++
		} else {
			minor = a2[j]
			j++
		}
		tmp[k] = minor
		k++
	}

	if i < len(a1) {
		copy(tmp[k:], a1[i:])
	} else {
		copy(tmp[k:], a2[j:])
	}

	copy(a, tmp)
}

func mergeSort(a []int, p, r int) {
	q := p + (r-p)>>1

	if p >= r {
		return
	}
	mergeSort(a, p, q)
	mergeSort(a, q+1, r)
	merge(a[p:r+1], a[p:q+1], a[q+1:r+1])
}

func main() {
	a := []int{6, 5, 3, 1, 8, 7, 2, 4, 9,10}
	mergeSort(a, 0, len(a)-1)
	fmt.Println(a)
	fmt.Println(sort.IntsAreSorted(a))
}
