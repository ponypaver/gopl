package main

import "fmt"

//none-in-place insert sort
//func nonInplaceInsertSort(a []int) []int {
//	n := len(a)
//	if len(a) <= 1 {
//		return a
//	}
//
//	b := make([]int, n)
//	b[0] = a[0]
//	i, j := 0, 0
//
//	for i = 1; i < n; i++ {
//		for j = 0; j < i; j++ {
//			if b[j] > a[i] {
//				break
//			}
//		}
//		copy(b[j+1:], b[j:])
//		b[j] = a[i]
//	}
//
//	return b
//}
//
//// in-place insert sort
//func insertSort(a []int) {
//	n := len(a)
//
//	for i := 1; i < n; i++ {
//		v := a[i]
//		for j := 0; j < i; j++ {
//			if a[j] > a[i] {
//				//gap := i - j
//				//copy(a[j+1:j+gap+1), a[j:j+gap])
//				copy(a[j+1:i+1], a[j:i])
//				a[j] = v
//			}
//		}
//	}
//}

func insertSort(a []int) {
	n := len(a)

	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
}

func main() {
	a := []int{6, 2, 9, 3, 5, 7, 8}
	//a := []int{1,2,3,4,5}
	insertSort(a)
	fmt.Println(a)
}
