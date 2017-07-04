package main

import "fmt"

func Fibonacci(n int) {

	memo := [10]int{0, 1}
	for i := 2; i < n; i++ {
		memo[i] = memo[i - 1] + memo[i - 2]
	}

	fmt.Println(memo)
}

func main() {
	Fibonacci(10)
}
