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

func bubblesort(arr []int) {
	var arr_len = len(arr)
	var loop_until = arr_len - 1
	is_sorted := false
	for !is_sorted {
		is_sorted = true
		for i := 0; i < loop_until; i++ {
			if arr[i] > arr[i+1] {
				var tmp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				is_sorted = false
			}
		}
	}

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
	bubblesort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}
