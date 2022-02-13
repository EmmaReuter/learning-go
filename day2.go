package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) float64 {
	z := 1.0
	for (z*z-x)/(2*z) > 0.00001 || (z*z-x)/(2*z) < -0.00001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	sum := 1

	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)

	if sum < 10 {
		fmt.Println("less than 10")
	}
	fmt.Println(sqrt(9))
	defer fmt.Println("Done with program")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Unknown")
	}

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
