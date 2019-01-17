package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	p := fmt.Println
	switch {
	case t.Hour() < 12:
		p("Good Morning!")
	case t.Hour() < 17:
		p("Good Afternoon!")
	default:
		p("Good Evening!")
	}
}
