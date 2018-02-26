// No different from that in C
package main

import "fmt"

func main() {
	x := 15
	a := &x //pointing to address of x
	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println(*a) //value of a changed
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a) //same as x
}
