package main

import "fmt"

//Modify reverse to reverse the characters of a []byte slice that represents a
//UTF-8-encoded str ing , in place. Can you do it without allocating new memory?
//func reverse(s string) string {
func reverse(b []byte) []byte {
	r := []rune(string(b))
	n := len(r)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return []byte(string(r))
}

func main() {
	//s := "铭刻了记忆，老去了悲凉"
	s := "风流雾香迷，月薄霞淡雨。红幽树芳飞，雪落花艳舞"
	fmt.Println(string(reverse([]byte(s))))
}