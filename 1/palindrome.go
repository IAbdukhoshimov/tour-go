package main

import "fmt"

func main() {
	fmt.Println(palindrome("dada"))
}

func palindrome(txt string) bool {
	l := len([]rune(txt))
	for i := 0; i <= l/2; i++ {
		if txt[i] != txt[l-1-i] {
			return false
		}
	}
	return true
}
