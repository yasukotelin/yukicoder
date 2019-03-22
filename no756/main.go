package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var d int
	fmt.Scanln(&d)

	champer := makeChamper(d)

	fmt.Println(champer[d+1 : d+2])
}

func makeChamper(digit int) string {
	var sb strings.Builder
	var count int
	sb.WriteString("0.")
	for sb.Len() < digit+3 {
		count++
		sb.WriteString(strconv.Itoa(count))
	}
	return sb.String()
}
