package main

import "fmt"

//declaration comes after parameter!
//func add(x, y int) int { works as well
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
