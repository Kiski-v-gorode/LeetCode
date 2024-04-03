package main

import (
	"cmp"
	"slices"
)

// maxOperations function takes a slice of integers `nums` and an integer `k` as input,
// and returns the maximum number of operations that can be performed on `nums`
// such that the sum of two elements equals `k`.
func maxOperations(nums []int, k int) int {
	// Check if the input slice `nums` is sorted.
	sortedOrNot := slices.IsSorted(nums)
	// If `nums` is not sorted, sort it using a custom comparison function.
	if !sortedOrNot {
		slices.SortFunc(nums, func(i, j int) int {
			return cmp.Compare(i, j)
		})
		// Call the helper function with the sorted `nums`.
		return helper(nums, k)
	} else {
		// If `nums` is already sorted, directly call the helper function.
		return helper(nums, k)
	}
}

// helper function takes a sorted slice of integers `nums` and an integer `k` as input,
// and returns the maximum number of operations that can be performed such that the sum of
// two elements equals `k`.

func helper(nums []int, k int) int {
	// Initialize two pointers at the start and end of the slice.
	left, right := 0, len(nums)-1
	// Initialize a variable to store the result.
	result := 0
	for left < right {
		// If the sum of elements at pointers `left` and `right` equals `k`,
		// increment `result` and move the pointers accordingly.
		if nums[left]+nums[right] == k {
			result++
			left++
			right--
			// If the sum is greater than `k`, decrement `right`.
		} else if nums[left]+nums[right] > k {
			right--
			// If the sum is less than `k`, increment `left`.
		} else {
			left++
		}
	}
	// Return the maximum number of operations performed.
	return result
}
