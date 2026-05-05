package main

import (
	"calculator/math"
	"fmt"

	"github.com/labstack/echo/v5"
)

func main() {

	// Echo instance
	e := echo.New()

	sum := math.Sum(5, 2)
	fmt.Println(sum)

	sub := math.Sub(5, 2)
	fmt.Println(sub)
}
