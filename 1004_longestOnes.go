package main

// Function longestOnes takes a slice of numbers nums and an integer k, and returns
// the length of the longest substring containing at most k zeros.
func longestOnes(nums []int, k int) int {
	// Variable countZeros is used to track the number of zeros in the current window.
	countZeros := 0
	// Variable result stores the length of the longest substring with at most k zeros.
	result := 0
	// Variable left indicates the left end of the current window.
	left := 0

	// Iterate over the right end of the window.
	for right := 0; right < len(nums); right++ {
		// If the current element is equal to 0, increment the zero count.
		if nums[right] == 0 {
			countZeros++
		}

		// If the number of zeros in the current window exceeds k, move the left
		// end of the window to the right until the number of zeros becomes less than or equal to k.
		if countZeros > k {
			for countZeros > k {
				if nums[left] == 0 {
					countZeros--
				}
				left++
			}
		}

		// Update the result with the length of the longest substring.
		result = max(result, right-left+1)
	}
	// Return the length of the longest substring with at most k zeros.
	return result
}
