package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("2")
	defer fmt.Println("1")

	fmt.Println("hello")
}
