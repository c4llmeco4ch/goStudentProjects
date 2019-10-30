package main

import(
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var(
	nums []int  //The slice of numbers we want to sort
	numVals int = -1
)

//User can optionally add a parameter that determines how many random numbers will be sorted
//If none are provided, 100 will be used
func main(){
	if len(os.Args) >= 2 {
		numVals, _ = strconv.Atoi(os.Args[1])
	} else{
		numVals = 100
	}
	nums = initSlice()
	nums = mergeSort(nums)
	for _, value := range nums {
		fmt.Printf("%d\n", value)
	}
}

func initSlice() []int {
	vals := []int{}
	for i := 0; i < numVals; i++ {
		vals = append(vals, rand.Int())
	}
	return vals
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 { //base case
		return arr
	}

	left := mergeSort(arr[:len(arr) / 2])
	right := mergeSort(arr[len(arr) / 2:])
	sortArr := []int{}
	lIndex, rIndex := 0, 0
	for lIndex < len(left) && rIndex < len(right) {
		leftLeast := left[lIndex] <= right[rIndex]
		if leftLeast {
			sortArr = append(sortArr, left[lIndex])
			lIndex++
		} else{
			sortArr = append(sortArr, right[rIndex])
			rIndex++
		}
	}
	if lIndex < len(left) {
		for ; lIndex < len(left); lIndex++ {
			sortArr = append(sortArr, left[lIndex])
		}
	}
	if rIndex < len(right) {
		for ; rIndex < len(right); rIndex++ {
			sortArr = append(sortArr, right[rIndex])
		}
	}
	return sortArr
}