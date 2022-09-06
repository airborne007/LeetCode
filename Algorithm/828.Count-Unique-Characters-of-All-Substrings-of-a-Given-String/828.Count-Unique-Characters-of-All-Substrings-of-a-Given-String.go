package main

import "fmt"

func uniqueLetterString(s string) (ans int) {
	last0, last1, total := [26]int{}, [26]int{}, 0
	for i := range last0 {
		last0[i] = -1
		last1[i] = -1
	}
	for i, c := range s {
		c -= 'A'
		total += (i - 2*last0[c] + last1[c])
		ans += total
		last0[c], last1[c] = i, last0[c]
	}
	return
}

func main() {
	fmt.Println(uniqueLetterString("ABC"))      // 10
	fmt.Println(uniqueLetterString("ABA"))      // 8
	fmt.Println(uniqueLetterString("LEETCODE")) // 92
}
