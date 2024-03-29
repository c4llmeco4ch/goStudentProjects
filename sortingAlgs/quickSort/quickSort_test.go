package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"testing"
)

func TestInvertedSort(t *testing.T) {
	arr := []byte{5, 4, 3, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			for j := range arr {
				val := arr[j]
				fmt.Print(strconv.Itoa(int(val)))
			}
			t.Errorf("%d > %d", arr[i], arr[i+1])
		}
	}
}

func TestSortedSort(t *testing.T) {
	arr := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			for j := range arr {
				val := arr[j]
				fmt.Print(strconv.Itoa(int(val)))
			}
			t.Errorf("%d > %d", arr[i], arr[i+1])
		}
	}
}

func TestRandomSlice(t *testing.T) {
	arr := make([]byte, 10000)
	_, err := rand.Read(arr)
	if err != nil {
		panic(err)
	}
	quickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			t.Errorf("%d > %d", arr[i], arr[i+1])
		}
	}
}
