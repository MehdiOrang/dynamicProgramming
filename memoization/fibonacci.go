package main

import (
	"fmt"
	"time"
)

func fib(n uint64) uint64 {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibo(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fibo(n-1, memo) + fibo(n-2, memo)
	return memo[n]

}

//1,1,2,3,5
func main() {
	start := time.Now()
	fmt.Println(fibo(45, map[int]int{}))
	end := time.Now()
	fmt.Println(end.Sub(start), "with memoization")
	start = time.Now()
	fmt.Println(fib(45))
	end = time.Now()
	fmt.Println(end.Sub(start), "without memoization")

}
