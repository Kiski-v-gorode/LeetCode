package main

// moveZeroes moves all zeroes to the end of the nums slice while preserving the order of non-zero elements.
func moveZeroes(nums []int) {
	// Initialize two pointers: left and right
	left, right := 0, 1
	// While the right pointer is within the bounds of the nums slice
	for right < len(nums) {
		// If the element pointed to by left is zero and the element pointed to by right is not zero
		if nums[left] == 0 && nums[right] != 0 {
			// Swap the zero element with the non-zero element
			nums[left], nums[right] = nums[right], nums[left]
			// Move left forward
			left++
			// Move right forward
			right++
			// If the element pointed to by left is not zero
		} else if nums[left] != 0 {
			// Move left forward
			left++
			// Move right forward
			right++
			// If both elements are zero
		} else {
			// Only move right forward
			right++
		}
	}
}
