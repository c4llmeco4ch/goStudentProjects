package main

import(
	"fmt"
	"rand"
	"os"
)

var(
	nums []int  //The slice of numbers we want to sort
	numVals int //The number of numbers we want to sort
)

//User can optionally add a parameter that determines how many random numbers will be sorted
//If none are provided, 100 will be used
func main(){
	if len(os.Args) >= 2 {
		numVals = os.Args[1]
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
	vals := [numVals]int{}
	for i := 0; i < numVals; i++ {
		vals = append(vals, rand.Int())
	}
	return vals
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 { //base case
		return arr
	}

	left = mergeSort(arr[:len(arr) / 2])
	right = mergeSort(arr[len(arr) / 2:])

	sortArr := [len(arr)]int{}
	lIndex, rIndex := 0, 0
	for lIndex < len(left) && rIndex < right {
		leftLeast := left[lIndex] <= right[rIndex]
		if leftLeast {
			sortArr = append(sortArr, left[lIndex++])
		} else{
			sortArr = append(sortArr, right[rIndex++])
		}
	}
	if lIndex < len(left) {
		sortArr = append(sortArr, left[lIndex:])
	}
	if rIndex < len(right) {
		sortArr = append(sortArr, right[rIndex:])
	}
	return sortArr
}