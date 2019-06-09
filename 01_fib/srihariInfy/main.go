package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//fibonacci function which gives fibonacci sequence for
//positive and negative numbers except 0
func fib(n int) {

	first, second, next, sign := 0, 1, 0, 1

	//to print positive or negative fibonacci sequence
	if n < 0 {
		sign = -1
	}

	//loop which iterates and prints fibonacci sequence
	for i := 1; i <= n*sign; i++ {
		if i <= 1 {
			next = i
		} else {
			next = first + second
			first = second
			second = next
		}
		//printing the number
		if i%2 == 0 {
			fmt.Fprintln(out, next*sign)
		} else {
			fmt.Fprintln(out, next)
		}

	}
}

func main() {
	fib(7)
}
