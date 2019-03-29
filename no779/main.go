package main

import (
	"fmt"
	"time"
)

var (
	startHeisei = time.Date(1989, time.January, 8, 0, 0, 0, 0, time.Local)
	endHeisei   = time.Date(2019, time.April, 30, 0, 0, 0, 0, time.Local)
)

func main() {
	var y int
	var m time.Month
	var d int
	fmt.Scanln(&y, &m, &d)

	t := time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	if startHeisei.Unix() <= t.Unix() && t.Unix() <= endHeisei.Unix() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
