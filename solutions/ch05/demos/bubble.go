package main

import "fmt"

// select sort
func selects(a []int) {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := i+1; j<n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func bubbles(a []int) {
	n := len(a)

	for i := 0; i < n; i++ {
		sorted := true
		for j := 0; j<n-1-i; j++ {
			if a[j] > a[j+1] {
				sorted = false
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		if sorted {
			break
		}
	}
}

func main() {
	//a := []int{6, 2, 9, 3, 5, 7, 8}
	a := []int{6,1,2,3,4,5}
	//a := []int{5,4,3,2,1}
	bubbles(a)
	fmt.Println(a)
}