package main

// maxArea function calculates the maximum area of water that can be trapped
// between vertical lines represented by heights in the given slice.
func maxArea(height []int) int {
	// Initialize two pointers at the beginning and end of the slice.
	left, right := 0, len(height)-1
	// Initialize the variable to store the maximum area.
	result := 0
	// Iterate until the left pointer is less than the right pointer.
	for left < right {
		// Find the minimum height between the two pointers.
		minHeight := min(height[left], height[right])
		// Calculate the area using the minimum height and the distance between the pointers.
		result = max(result, minHeight*(right-left+1))
		// Move the pointer with the smaller height towards the center,
		// since a lower height will not increase the area, but can only reduce it.
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	// Return the maximum area found.
	return result
}
