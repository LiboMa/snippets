package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(seq []int) ([]int, time.Duration, error) {
	start := time.Now()
	n := len(seq)
	var j int
	for i := 0; i < (n - 1); i++ {
		for j = 0; j < (n - 1 - i); j++ {

			// Use ascent sort order
			if seq[j] > seq[j+1] {
				swap := seq[j]
				seq[j] = seq[j+1]
				seq[j+1] = swap
			}
		}
	}
	elaspad := time.Since(start)
	fmt.Println("time elaspad: ", elaspad)
	return seq, elaspad, nil

}

func main() {

	seq := []int{}
	N := 100000
	for i := 0; i < N; i++ {
		seq = append(seq, rand.Intn(N))
	}

	s, t, _ := bubbleSort(seq)
	fmt.Printf("%v\n, time elaspad: %v, length: %d\n", s, t, len(s))

}
