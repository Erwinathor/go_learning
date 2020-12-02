package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 14
	for i := 2; i < 14; i++ {
		sum2 += i
	}
	fmt.Println(sum2)
}
