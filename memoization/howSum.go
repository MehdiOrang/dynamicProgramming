package main

import (
	"fmt"
	"time"
)

func howSum(N int, Numbers []int, memo map[int][]int) []int {
	if _, ok := memo[N]; ok {
		return memo[N]
	}

	if N == 0 {
		return []int{}
	}
	if N < 0 {
		return nil
	}
	for _, number := range Numbers {
		remainder := N - number
		remainderSum := howSum(remainder, Numbers, memo)
		if remainderSum != nil {
			memo[N] = append(remainderSum, number)
			return memo[N]
		}

	}
	memo[N] = nil
	return nil
}

func main() {
	start := time.Now()
	result := howSum(7, []int{2, 3}, map[int][]int{}) //[3,2,2]
	end := time.Now()
	fmt.Println(result, end.Sub(start), map[int][]int{})

	start = time.Now()
	result = howSum(7, []int{5, 3, 4, 7}, map[int][]int{}) // [4,3]
	end = time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = howSum(8, []int{2, 3, 5}, map[int][]int{}) //[2,2,2,2]
	end = time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = howSum(300, []int{7, 14}, map[int][]int{}) //nil
	end = time.Now()
	fmt.Println(result, end.Sub(start))
}
