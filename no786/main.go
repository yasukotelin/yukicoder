package main

import "fmt"

var pattern int
var limit int

func main() {
	fmt.Scanln(&limit)

	stepCount(0)

	fmt.Println(pattern)
}

func stepCount(total int) {
	if limit == total {
		pattern++
		return
	}
	if limit < total {
		return
	}

	stepCount(total + 1)
	stepCount(total + 2)
}
