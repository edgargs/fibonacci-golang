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
	fmt.Println("Value of parameter is: ",arg)
	x, err := strconv.Atoi(arg)
	var result uint64 = 0
	if err == nil {
		for i := 0; i <= x; i++ {
			result = f()
		}
		fmt.Println("Result with Golang: ",result)
	}
}
