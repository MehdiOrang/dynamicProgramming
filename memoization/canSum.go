package main

import (
	"fmt"
	"time"
)

func canSum(N int, Numbers []int, memo map[int]bool) bool {
	if _, ok := memo[N]; ok {
		return memo[N]
	}
	if N == 0 {
		return true
	}
	if N < 0 {
		return false
	}

	for _, number := range Numbers {
		remainder := N - number

		if canSum(remainder, Numbers, memo) {
			memo[N] = true
			return true
		}

	}
	memo[N] = false
	return false
}

func main() {
	start := time.Now()
	result := canSum(7, []int{2, 3}, map[int]bool{}) //true
	end := time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = canSum(7, []int{5, 3, 4, 7}, map[int]bool{}) //true
	end = time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = canSum(8, []int{2, 3, 5}, map[int]bool{}) //true
	end = time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = canSum(300, []int{7, 14}, map[int]bool{}) //false
	end = time.Now()
	fmt.Println(result, end.Sub(start))
}
