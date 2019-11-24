package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

var (
	nums    []byte
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

func partition(arr []byte, start int, end int) int {
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

func quickSort(arr []byte, start int, end int) {
	if start < end {
		pi := partition(arr, start, end)

		quickSort(arr, start, pi-1)
		quickSort(arr, pi+1, end)
	}
}

func initSlice() []byte {
	vals := make([]byte, numVals)
	_, err := rand.Read(vals)
	if err != nil {
		panic(err)
	}
	return vals
}
