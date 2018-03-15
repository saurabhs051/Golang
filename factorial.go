package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 || n == 1 {            // '||' and '&&', not 'or' or 'and'  
					//if can have brackets 
		return 1
	} else {			//else in same line
		return (n * fact(n-1))
	}
}

func main() {
	n := 5
	fmt.Println(fact(n))
}
