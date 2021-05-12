package main

import "testing"

func swapPlaces(array []int, j int) {
	array[j], array[j+1] = array[j+1], array[j]
}

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				swapPlaces(array, j)
			}
		}
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := []int{
			48, 96, 86, 42,
			57, 82, 63, 70,
			37, 34, 43, 27,
			19, 97, 9, 17,
			48, 96, 86, 68,
			48, 12, 86, 68,
			79, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 70,
			57, 82, 48, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}

		bubbleSort(x)
	}
}
