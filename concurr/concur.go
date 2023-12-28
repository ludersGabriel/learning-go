package main

import (
	"fmt"
	"time"
)

func Sum(s []int, c chan int) {
	var sum int = 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

const N = 100000000

func main() {
	var s []int = make([]int, N)

	for i := range s {
		s[i] = i
	}

	start := time.Now()
	// sum of 1 to N sequential
	var sum int = 0
	for _, v := range s {
		sum += v
	}

	elapsed := time.Since(start)

	fmt.Printf("Time taken for sequential sum: %s\n", elapsed)

	c := make(chan int)
	start = time.Now()

	go Sum(s[:N/4], c)
	go Sum(s[N/4:s[N/4*2]], c)
	go Sum(s[N/4*2:s[N/4*3]], c)
	go Sum(s[N/4*3:], c)

	x, y, z, w := <-c, <-c, <-c, <-c

	elapsed = time.Since(start)

	fmt.Printf("Time taken for concurrent sum: %s\n", elapsed)

	fmt.Printf("Sums are: %d, %d\n", sum, x+y+z+w)
}
