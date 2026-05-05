package main

import (
	"calculator/math"
	"fmt"
)

func main() {
	sum := math.Sum(5, 2)
	fmt.Println(sum)

	sub := math.Sub(5, 2)
	fmt.Println(sub)
}
