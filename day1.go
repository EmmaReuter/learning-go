package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func main() {
	fmt.Printf("You now you have %g problems\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(23, 32))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(i, j)

	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
}
