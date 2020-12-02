package main

import "fmt"

func main() {
	sum := 1

	//for i := 0; i < 10; i++  the init and post are optional
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
