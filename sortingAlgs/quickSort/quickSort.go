package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	nums    []int
	numVals int = -1
)

func main() {
	if len(os.Args) >= 2 {
		numVals, _ = strconv.Atoi(os.Args[1])
	} else {
		numVals = 100
	}
	nums = initSlice()
	quickSort(nums, 0, len(nums)-1)
	for _, val := range nums {
		fmt.Printf("%d\n", val)
	}
}

func partition(arr []int, start int, end int) int {
	pos := start - 1
	part := arr[end]
	for i := start; i < end; i++ {
		if arr[i] < part {
			pos++
			arr[pos], arr[i] = arr[i], arr[pos]
		}
	}
	pos++
	arr[pos], arr[end] = arr[end], arr[pos]
	return pos
}

func quickSort(arr []int, start int, end int) {
	if start < end {
		pi := partition(arr, start, end)

		quickSort(arr, start, pi-1)
		quickSort(arr, pi+1, end)
	}
}

func initSlice() []int {
	vals := []int{}
	for i := 0; i < numVals; i++ {
		vals = append(vals, rand.Int())
	}
	return vals
}
