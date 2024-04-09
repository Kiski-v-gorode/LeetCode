package main

// Function longestSubarray finds the length of the longest contiguous subarray
// with at most one zero in the given array of integers nums.
// It returns the length of the longest subarray.
func longestSubarray(nums []int) int {
	// Initialize variables to track the result (length of the longest subarray),
	// the left pointer for the sliding window, and the count of zeros in the current window.
	result := 0
	left := 0
	countZeros := 0

	// Iterate over the elements of the array using the right pointer.
	for right := 0; right < len(nums); right++ {
		// If the current element is zero, increment the count of zeros.
		if nums[right] == 0 {
			countZeros++
		}

		// If the count of zeros exceeds 1, move the left pointer to maintain
		// at most one zero in the current window.
		if countZeros > 1 {
			for countZeros > 1 {
				// If the element at the left end of the window is zero, decrement the count of zeros.
				if nums[left] == 0 {
					countZeros--
				}
				// Move the left pointer to the right to slide the window.
				left++
			}
		}

		// Update the result with the length of the current subarray.
		result = max(result, right-left)
	}

	// Return the length of the longest subarray with at most one zero.
	return result
}
