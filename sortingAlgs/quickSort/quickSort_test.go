package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInvertedSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			for j := range arr {
				val := arr[j]
				fmt.Print(strconv.Itoa(val))
			}
			t.Errorf("%d > %d", arr[i], arr[i+1])
		}
	}
}

func TestSortedSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			for j := range arr {
				val := arr[j]
				fmt.Print(strconv.Itoa(val))
			}
			t.Errorf("%d > %d", arr[i], arr[i+1])
		}
	}
}
