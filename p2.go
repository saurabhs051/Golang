package main

import (
	"fmt"
	"math"
	"math/rand" //necessary to import sub-package //can't skip
)

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println("The square root of 4 is : ", math.Sqrt(4))
	fmt.Println("A number from 1-100 : ", rand.Intn(100))
	num1, num2 := 5.6, 6.5
	var a int = 62
	fmt.Println(float64(a))
	w1, w2 := "hello", "there"

	fmt.Println(add(num1, num2))
	fmt.Println(multiple(w1, w2))

}
