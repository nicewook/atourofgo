 package main

 import (
   "fmt"
 )

 func add(x, y int) int {  // x, y int == x int, y int
   return x + y
 }


 func main() {
   fmt.Println(add(3, 7))
 }

