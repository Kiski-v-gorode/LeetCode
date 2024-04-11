package main

// kidsWithCandies returns a slice of booleans indicating whether each child has at least the given number of candies plus any additional candies.
func kidsWithCandies(candies []int, extraCandies int) []bool {

	// myMax returns the largest value in the provided slice of integers.
	myMax := func(arr []int) int {
		// result will hold the largest value
		result := 0

		// loop through each value in the array
		for _, value := range arr {
			// check if the current value is greater than the current largest value
			if value > result {
				// if the current value is greater, update the largest value
				result = value
			}
		}

		// return the largest value
		return result
	}

	// result is a slice of booleans that will be returned
	result := make([]bool, len(candies))

	// maxCandies is the maximum number of candies among all the children
	maxCandies := myMax(candies)

	// loop through each child and check if they have enough candies
	for i := 0; i < len(candies); i++ {
		// add the extra candies to the current child's count
		candies[i] += extraCandies

		// check if the current child has enough candies
		if candies[i] >= maxCandies {
			// the current child has enough candies, so set the result to true
			result[i] = true
		} else {
			// the current child does not have enough candies, so set the result to false
			result[i] = false
		}

		// subtract the extra candies from the current child's count
		candies[i] -= extraCandies
	}

	// return the result slice
	return result
}
