package main

import "fmt"

//split ontvang 16 en stuurt x en y terug
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
