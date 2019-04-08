package main

import "fmt"

func main() {
	var n string
	fmt.Scanln(&n)

	fmt.Println(check(n))
}

func check(n string) int {
	if len(n) == 1 || n[0:1] != "1" {
		return -1
	}
	for i := 1; i < len(n); i++ {
		if n[i:i+1] != "3" {
			return -1
		}
	}
	return len(n) - 1
}
