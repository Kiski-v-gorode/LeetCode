package main

// productExceptSelf calculates the product of all elements in the array nums except self.
// It returns a new array where each element at index i contains the product of all elements in nums except nums[i].
func productExceptSelf(nums []int) []int {
	n := len(nums)

	// Create a prefix array to store the product of all elements before each index (inclusive).
	prefix := make([]int, n+1)
	prefix[0] = 1
	for i := 0; i < n; i++ {
		prefix[i+1] = nums[i] * prefix[i]
	}

	// Create a postfix array to store the product of all elements after each index (inclusive).
	postfix := make([]int, n+1)
	postfix[n] = 1
	for i := n - 1; i >= 0; i-- {
		postfix[i] = nums[i] * postfix[i+1]
	}

	// Calculate the product of elements except self by multiplying corresponding elements from prefix and postfix arrays.
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, prefix[i]*postfix[i+1])
	}

	return result
}
