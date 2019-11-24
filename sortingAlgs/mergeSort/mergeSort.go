package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

var (
	nums    []byte //The slice of numbers we want to sort
	numVals int    = -1
)

//User can optionally add a parameter that determines how many random numbers will be sorted
//If none are provided, 100 will be used
func main() {
	if len(os.Args) >= 2 {
		numVals, _ = strconv.Atoi(os.Args[1])
	} else {
		numVals = 100
	}
	nums = initSlice()
	nums = mergeSort(nums)
	for _, value := range nums {
		fmt.Printf("%d\n", value)
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

func mergeSort(arr []byte) []byte {
	if len(arr) <= 1 { //base case
		return arr
	}

	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	sortArr := make([]byte, len(arr))
	lIndex, rIndex := 0, 0
	for lIndex < len(left) && rIndex < len(right) {
		leftLeast := left[lIndex] <= right[rIndex]
		if leftLeast {
			sortArr[lIndex+rIndex] = left[lIndex]
			lIndex++
		} else {
			sortArr[lIndex+rIndex] = right[rIndex]
			rIndex++
		}
	}
	if lIndex < len(left) {
		for ; lIndex < len(left); lIndex++ {
			sortArr[lIndex+rIndex] = left[lIndex]
		}
	}
	if rIndex < len(right) {
		for ; rIndex < len(right); rIndex++ {
			sortArr[lIndex+rIndex] = right[rIndex]
		}
	}
	return sortArr
}
