package main

import "fmt"

func s(a []int, low, high, t int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2

	if t == a[mid] {
		return mid
	}

	if t < a[mid] {
		high = mid - 1
	} else {
		low = mid + 1
	}

	return s(a, low, high, t)
}

func bs(a []int, n, v int) int {
	var (
		low  = 0
		high = n - 1
	)

	for low <= high {
		mid := (low + high) / 2
		if v == a[mid] {
			return mid
		}
		if v < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range a {
		fmt.Printf("search %v, position: %v\n", i, s(a, 0, len(a)-1, i))
		fmt.Printf("search %v, position: %v\n", i, bs(a, len(a), i))
	}
}