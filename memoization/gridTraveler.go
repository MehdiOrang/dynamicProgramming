package main

import (
	"fmt"
	"strconv"
	"time"
)

func traveler(m int, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = traveler(m-1, n, memo) + traveler(m, n-1, memo)
	return memo[key]
}

func main() {
	start := time.Now()
	fmt.Println(traveler(1,1 , map[string]int{}))
	end := time.Now()
	fmt.Println(end.Sub(start))
	start = time.Now()
	fmt.Println(traveler(2, 3, map[string]int{}))
	end = time.Now()
	fmt.Println(end.Sub(start))
	start = time.Now()
	fmt.Println(traveler(3, 3, map[string]int{}))
	end = time.Now()
	fmt.Println(end.Sub(start))
	start = time.Now()
	fmt.Println(traveler(18,18, map[string]int{}))
	end = time.Now()
	
}
