package main

import (
    "fmt"
	"os"
	"strconv"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() uint64 {
	var a, b uint64 = 0, 1
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	//fmt.Println(f(), f(), f(), f(), f())
	arg := os.Args[1]
	x, err := strconv.Atoi(arg)
	if err == nil {
		for i := 0; i < x; i++ {
			fmt.Println(f())
		}
	}
}