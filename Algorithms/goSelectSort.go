package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(seq []int) ([]int, time.Duration) {
	start := time.Now()
	n := len(seq) - 1
	var minIndex int
	for i := 0; i < (n); i++ {
		minIndex = i
		for j := i + 1; j < n; j++ {
			if seq[j] < seq[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			swap := seq[i]
			seq[i] = seq[minIndex]
			seq[minIndex] = swap
		}
	}
	elasped := time.Since(start)

	return seq, elasped
}
func main() {

	seq := []int{}
	N := 10000
	for i := 0; i < N; i++ {
		seq = append(seq, rand.Intn(1000))
	}

	s, t := selectionSort(seq)
	fmt.Printf("%v\n, time elaspad: %v, length: %d\n", s, t, len(s))
}
