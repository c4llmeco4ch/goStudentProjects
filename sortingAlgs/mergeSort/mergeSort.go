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
	ms := make(chan byte)
	go mergeSort(nums, ms)
	pos := 0
	for val := range ms {
		nums[pos] = val
		pos++
	}
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

func mergeSort(arr []byte, ms chan byte) {
	if len(arr) <= 1 {
		if len(arr) == 1 { //base case
			ms <- arr[0]
		}
		close(ms)
		return
	}
	leftMS := make(chan byte)
	go mergeSort(arr[:len(arr)/2], leftMS)
	rightMS := make(chan byte)
	go mergeSort(arr[len(arr)/2:], rightMS)
	left, lOK := <-leftMS
	right, rOK := <-rightMS
	for lOK && rOK {
		leftLeast := left <= right
		if leftLeast {
			ms <- left
			left, lOK = <-leftMS
		} else {
			ms <- right
			right, rOK = <-rightMS
		}
	}
	if lOK {
		ms <- left
		for val := range leftMS {
			ms <- val
		}
	}
	if rOK {
		ms <- right
		for val := range rightMS {
			ms <- val
		}
	}
	close(ms)
}
