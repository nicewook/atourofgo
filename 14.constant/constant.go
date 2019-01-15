package main

import "fmt"

const Pi = 3.14

// constant can be char, string, boolean, number

func main() {
  const World = "안녕"
  p := fmt.Println
  p("Hello",  World)
  p("Happy", Pi, "Day")

  const Truth = true
  p("Go rules?", Truth)
}

