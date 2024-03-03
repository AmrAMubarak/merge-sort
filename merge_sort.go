package main

import "fmt"

func merge(arr []int, start int, midpoint int, end int) {
	var left_length int = midpoint - start + 1
	var right_length int = end - midpoint

	var left_array []int = make([]int, left_length)
	var right_array []int = make([]int, right_length)

	for i := 0; i < left_length; i++ {
		left_array[i] = arr[start+i]
	}

	for j := 0; j < right_length; j++ {
		right_array[j] = arr[midpoint+1+j]
	}

	i, j, k := 0, 0, start

	for i < left_length && j < right_length {
		if left_array[i] <= right_array[j] {
			arr[k] = left_array[i]
			k++
			i++
		} else {
			arr[k] = right_array[j]
			k++
			j++
		}
	}
	for i < left_length {
		arr[k] = left_array[i]
		k++
		i++
	}
	for j < right_length {
		arr[k] = right_array[j]
		k++
		j++
	}
}

func merge_sort(arr []int, start int, end int) {
	if end <= start {
		return
	}
	var midpoint int = (start + end) / 2
	merge_sort(arr, start, midpoint)
	merge_sort(arr, midpoint+1, end)
	merge(arr, start, midpoint, end)
}

func main() {
	arr := [8]int{4, -9, 4, 0, -3, 1, 8, 2}
	merge_sort(arr[:], 0, 7)
	for i := 0; i < 8; i++ {
		fmt.Print(arr[i], " ")
	}
}
