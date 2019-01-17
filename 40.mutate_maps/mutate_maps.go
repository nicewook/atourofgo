package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	p := fmt.Println
	p("The value", m["Answer"])

	m["Answer"] = 48
	p("The value", m["Answer"])

	delete(m, "Answer")
	p("The value", m["Answer"])

	v, ok := m["Answer"]
	p("The value", v, "Present?", ok)
}
