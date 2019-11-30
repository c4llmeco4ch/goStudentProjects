package main

import (
	"crypto/rand"
	"testing"
)

func TestInvertedSort(t *testing.T) {
	arr := []byte{5, 4, 3, 2, 1}
	c := make(chan byte)
	go mergeSort(arr, c)
	last := byte(0)
	for val := range c {
		if val < last {
			t.Errorf("%d > %d", last, val)
		} else {
			last = val
		}
	}
}

func TestSortedSort(t *testing.T) {
	arr := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan byte)
	go mergeSort(arr, c)
	last := byte(0)
	for val := range c {
		if val < last {
			t.Errorf("%d > %d", last, val)
		} else {
			last = val
		}
	}
}

func TestRandomSlice(t *testing.T) {
	arr := make([]byte, 500)
	_, err := rand.Read(arr)
	if err != nil {
		panic(err)
	}
	c := make(chan byte)
	go mergeSort(arr, c)
	last := byte(0)
	for val := range c {
		if val < last {
			t.Errorf("%d > %d", last, val)
		} else {
			last = val
		}
	}
}
