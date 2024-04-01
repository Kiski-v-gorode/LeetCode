package main

// increasingTriplet checks whether there exists an increasing triplet sequence in the array nums.
// The function returns true if such a sequence exists, otherwise it returns false.
func increasingTriplet(nums []int) bool {
	n := len(nums)

	// If the length of the array is less than 3, an increasing triplet sequence cannot exist,
	// so return false immediately.
	if n < 3 {
		return false
	}

	// Create the leftMin array, where for each element nums[i],
	// it stores the minimum value of elements up to index i inclusive.
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	// Create the rightMax array, where for each element nums[i],
	// it stores the maximum value of elements after index i inclusive.
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	// Iterate over the nums array and check if there exists a triplet where
	// the first number is less than the second, and the second is less than the third.
	for i, num := range nums {
		if leftMin[i] < num && num < rightMax[i] {
			return true
		}
	}

	// If no such triplet is found, return false
	return false
}
