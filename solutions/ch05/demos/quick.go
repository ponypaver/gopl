package main

import "fmt"

func partition(a []int, low, high int) int {
	//pivot := a[high]
	i, j := low, low
	for ; j < high; j++ {
		if a[j] < a[high] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[i], a[high] = a[high], a[i]

	return i
}

func quickSort(a []int, low, high int) {
	if low < high {
		pi := partition(a, low, high)
		quickSort(a, low, pi -1 )
		quickSort(a, pi+1, high)
	}
}

func main() {
	//a := []int{1, 3, 5, 7, 8, 9, 4, 10, 0, 2}
	//a := []int {3,7,8,5,2,1,9,5,4}
	a := []int{1, 0}
	pi := partition(a, 0, len(a)-1)
	fmt.Println(pi, a)
	//quickSort(a, 0, len(a) -1 )
	//fmt.Println(a)
}