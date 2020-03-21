package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(seq *[]int) ([]int, time.Duration) {
	start := time.Now()
	n := len(*seq)
	var minIndex int
	// Outer Loop for looping the whole list by index
	for i := 0; i < n; i++ {
		minIndex = i
		// Inner Loop for looping the dynamic sub list
		for j := i + 1; j < n; j++ {
			if (*seq)[j] < (*seq)[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			//swap := seq[i]
			//seq[i] = seq[minIndex]
			//seq[minIndex] = swap
			(*seq)[i], (*seq)[minIndex] = (*seq)[minIndex], (*seq)[i]
		}
	}
	elasped := time.Since(start)

	return *seq, elasped
}

func main() {

	seq := []int{}
	N := 300000
	for i := 0; i < N; i++ {
		seq = append(seq, rand.Intn(N))
	}

	s, t := selectionSort(&seq)
	fmt.Printf("%v\n, time elaspad: %v, length: %d\n", s, t, len(s))
}
