package main

import (
	"fmt"
	"time"
)

func howSum(N int, Numbers []int) []int {
	if N == 0 {
		return nil
	}
	if N < 0 {
		return []int{}
	}

	for _, number := range Numbers {
		remainder := N - number
		//	fmt.Println(remainder)
		remainderSum := howSum(remainder, Numbers)
		//fmt.Println(remainderSum)
		if remainderSum != nil {
			remainderSum = append(remainderSum, number)
			return remainderSum
		}

	}

	return nil
}

func main() {
	start := time.Now()
	result := howSum(7, []int{2, 3}) //[3,2,2]
	end := time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = howSum(7, []int{5, 3, 4, 7}) // [4,3]
	end = time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = howSum(8, []int{2, 3, 5}) //[2,2,2,2]
	end = time.Now()
	fmt.Println(result, end.Sub(start))

	start = time.Now()
	result = howSum(300, []int{7, 14}) //nil
	end = time.Now()
	fmt.Println(result, end.Sub(start))
}
