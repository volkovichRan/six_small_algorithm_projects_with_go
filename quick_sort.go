package main

import (
	"fmt"
	"math/rand"
	"time"
)

func make_random_array(num_items, max int) []int {
	array := make([]int, num_items)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		var random = rand.Intn(max)
		array[i] = random
	}
	return array
}

func print_array(arr []int, num_items int) {
	var print_size = num_items
	var array_length = len(arr)

	if print_size > array_length {
		print_size = array_length
	}

	fmt.Println(arr[0:print_size])
}

func check_sorted(arr []int) {
	var sorted = true
	var arr_len = len(arr)
	var loop_until = arr_len - 1
	for i := 0; i < loop_until; i++ {
		if arr[i] > arr[i+1] {
			sorted = false
		}
	}
	if sorted {
		fmt.Println("The array is sorted")
	} else {
		fmt.Println("The array is NOT sorted!")
	}

}

func partition(slice []int) int {
	hi := len(slice) - 1
	pivot := slice[hi]
	var lo = 0
	tmp_pivot := lo - 1
	for lo := 0; lo < hi; lo++ {
		if slice[lo] <= pivot {
			tmp_pivot++
			tmp := slice[lo]
			slice[lo] = slice[tmp_pivot]
			slice[tmp_pivot] = tmp
		}
	}
	tmp_pivot++
	tmp := slice[tmp_pivot]
	slice[tmp_pivot] = slice[hi]
	slice[hi] = tmp
	return tmp_pivot
}

func quicksort(slice []int) {
	if len(slice) < 2 {
		return
	}
	pivot := partition(slice)
	quicksort(slice[0:pivot])
	quicksort(slice[pivot+1 : len(slice)])
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := make_random_array(num_items, max)
	print_array(arr, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}
