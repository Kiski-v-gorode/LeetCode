package main

// Function findDifference finds the difference between two slices of integers, nums1 and nums2.
// It returns a slice of slices, where each inner slice contains the elements that are present
// in one slice but not in the other.
func findDifference(nums1 []int, nums2 []int) [][]int {

	// Function createMap creates maps from the elements of two slices of integers, nums1 and nums2.
	// It returns two maps where the keys represent the unique elements of nums1 and nums2, respectively.
	createMap := func(nums1 []int, nums2 []int) (map[int]bool, map[int]bool) {
		// Initialize variables to track the current index in both slices.
		first := 0
		second := 0
		// Get the lengths of nums1 and nums2.
		firstLen := len(nums1)
		secondLen := len(nums2)
		// Create maps to store unique elements from nums1 and nums2.
		firstMap := make(map[int]bool)
		secondMap := make(map[int]bool)

		// Iterate over the elements of both slices.
		for first < firstLen || second < secondLen {
			// If there are elements remaining in nums1, add them to firstMap.
			if first < firstLen {
				if !firstMap[nums1[first]] {
					firstMap[nums1[first]] = true
				}
				first++
			}

			// If there are elements remaining in nums2, add them to secondMap.
			if second < secondLen {
				if !secondMap[nums2[second]] {
					secondMap[nums2[second]] = true
				}
				second++
			}
		}

		// Return the maps containing unique elements from nums1 and nums2.
		return firstMap, secondMap
	}

	// Create maps to store the elements of nums1 and nums2 as keys for efficient lookup.
	nums1Map, nums2Map := createMap(nums1, nums2)

	// Initialize the result slice.
	result := [][]int{}
	subResult := []int{}

	// Iterate over the elements in nums1Map.
	for key := range nums1Map {
		// If the key is present in nums1Map but not in nums2Map, add it to the subResult slice.
		if !nums2Map[key] {
			subResult = append(subResult, key)
		}
	}
	// Add the subResult slice to the result slice.
	result = append(result, subResult)

	// Reset subResult slice.
	subResult = []int{}

	// Iterate over the elements in nums2Map.
	for key := range nums2Map {
		// If the key is present in nums2Map but not in nums1Map, add it to the subResult slice.
		if !nums1Map[key] {
			subResult = append(subResult, key)
		}
	}
	// Add the subResult slice to the result slice.
	result = append(result, subResult)

	// Return the result slice containing the differences between nums1 and nums2.
	return result
}
