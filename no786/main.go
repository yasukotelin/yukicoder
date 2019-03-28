package main

import "fmt"

func main() {
	/**
	周期性が存在する。
	nが持ちうるパターンは n(通り) = (n-1通り) + (n-2通り)

	n = 1 {1} 1通り
	n = 2 {1, 1} {2} 2通り
	n = 3 {1, 1, 1}{1, 2}{2, 1} 3通り
	n = 4 {1, 1, 1, 1}{1, 1, 2}{1, 2, 1}{2, 1, 1}{2, 2} 5通り
	n = 5 {...}... 8通り

	*/
	var n int
	fmt.Scanln(&n)

	np, np1, np2 := 0, 1, 0
	for i := 0; i < n; i++ {
		np = np2 + np1
		np2 = np1
		np1 = np
	}

	fmt.Println(np)
}
