package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	total int
)

func QuickSort(seq *[]int, l int, r int) {
	total++
	// return nothing if the sub-array has only two members
	if l >= r {
		return
	}

	piovt := r
	counter := l

	for i := l; i <= piovt; i++ {
		// find the elements which less than piovt, and put it the left of piovt
		if (*seq)[i] <= (*seq)[piovt] {
			//swap((&seq)[i], (&seq)[piovt])
			(*seq)[counter], (*seq)[i] = (*seq)[i], (*seq)[counter]
			counter++
		}
	}
	// Sort left side of first piovt
	QuickSort(seq, l, (counter - 2)) // counter -2, farback from privious 2 indexing, excullte piovt position
	// Sort from side of first piovt
	QuickSort(seq, counter, r)

}

func main() {

	//arr := []int{10, 2, 1, 7, 9, 5}
	arr := []int{}
	N := 10000
	for i := 0; i < N; i++ {
		arr = append(arr, rand.Intn(N))
	}
	start := time.Now()
	n := len(arr)
	QuickSort(&arr, 0, n-1)
	end := time.Since(start)
	fmt.Println(arr)
	fmt.Println("CPU Time Used: ", end, "total compute", total)

}
