package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	total := 1

	for i := 0; i < 3; i++ {
		hexNum := 16
		line := readLine(sc)
		if line == "NONE" {
			total *= hexNum * hexNum // total * F * F
			continue
		}
		rmNum := len(strings.Split(line, ","))
		hexNum -= rmNum

		total *= hexNum * hexNum
	}

	fmt.Println(total)
}

func readLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
