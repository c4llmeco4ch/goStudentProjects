package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var (
	nums   []int
	values int
)

func main() {
	if len(os.Args) >= 2 {
		values, _ = strconv.Atoi(os.Args[1])
	} else {
		values = 100
	}
	nums = initSlice()
	place := 1
	maxVal := maxInt(nums)
	for maxVal/place >= 1 {
		nums = radixSort(nums, place)
		place *= 10
	}
	printSlice(nums)
}

func maxInt(arr []int) int {
	biggest := -1 * (math.MinInt32 - 1)
	for _, num := range arr {
		if num > biggest {
			biggest = num
		}
	}
	return biggest
}

func radixSort(arr []int, place int) []int {
	digits := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, value := range arr {
		loc := (value / place) % 10
		digits[loc]++
	}
	for pos := 1; pos < len(digits); pos++ {
		digits[pos] += digits[pos-1]
	}
	return rearrange(arr, digits, place)
}

func rearrange(input []int, digits [10]int, place int) []int {
	output := make([]int, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		loc := (input[i] / place) % 10
		output[digits[loc]-1] = input[i]
		digits[loc]--
	}
	return output
}

func initSlice() []int {
	vals := []int{}
	for i := 0; i < values; i++ {
		vals = append(vals, rand.Int())
	}
	return vals
}

func printSlice(arr []int) {
	for key, val := range arr {
		fmt.Printf("| %d : %d |\n", key, val)
	}
}
