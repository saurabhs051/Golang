package main

import "fmt"

func add(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func main() {
	m, n := 7, 5
	fmt.Print("Result : ") //can't use Println("Result : ",add(m,n)) here
	fmt.Println(add(m, n))
}
