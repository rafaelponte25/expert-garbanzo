package main

import (
	"fmt"
	"time"
)

func main() {
	timeTest()
}

func timeTest() {
	var defaultStart, end time.Time
	end = time.Now()
	defaultStart = end.Add(-15 * time.Minute)

	// 2021-12-17 17:07:54.959394 -0300 -03 m=+0.000122633
	fmt.Println(end)

	// 2021-12-17 16:52:54.959394 -0300 -03 m=-899.999877367
	fmt.Println(defaultStart)
}
