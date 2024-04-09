package main

// Function pivotIndex takes a slice of numbers nums and returns the index of an element
// that acts as a "pivot", i.e., the sum of elements to its left is equal to the sum of
// elements to its right. If there is no such index, it returns -1.
func pivotIndex(nums []int) int {
	// Create prefix and postfix arrays to compute sums of elements to the left and right
	// of each element in nums, respectively.
	prefix := createPrefixArray(nums)
	postfix := createPostfixArray(nums)

	// Iterate over the indices of nums.
	for i := 0; i < len(nums); i++ {
		// If the sum of elements to the left of nums[i] is equal to the sum of elements
		// to the right, then this index i is the "pivot", return it.
		if prefix[i] == postfix[i+1] {
			return i
		}
	}
	// If no "pivot" index is found, return -1.
	return -1
}

// Function createPrefixArray creates a prefix array for the given slice of numbers nums,
// where each element of the prefix array is the sum of elements of nums up to the corresponding
// element (inclusive).
func createPrefixArray(nums []int) []int {
	// Create a result array of length one more than nums for storing prefix sums.
	result := make([]int, len(nums)+1)
	// Compute prefix sums.
	for i := 1; i < len(result); i++ {
		result[i] = nums[i-1] + result[i-1]
	}
	return result
}

// Function createPostfixArray creates a postfix array for the given slice of numbers nums,
// where each element of the postfix array is the sum of elements of nums after the corresponding
// element (inclusive).
func createPostfixArray(nums []int) []int {
	// Create a result array of length one more than nums for storing postfix sums.
	result := make([]int, len(nums)+1)
	// Compute postfix sums.
	for i := len(result) - 2; i >= 0; i-- {
		result[i] = result[i+1] + nums[i]
	}
	return result
}
