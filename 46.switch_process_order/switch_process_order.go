package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	p("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		p("Today")
	case today + 1:
		p("Tomorrow")
	case today + 2:
		p("In two days")
	default:
		p("Too far away.")
	}
}
