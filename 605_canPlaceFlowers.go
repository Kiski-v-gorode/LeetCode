package main

// canPlaceFlowers determines if it is possible to place n flowers in a flowerbed
// The flowerbed is represented as an array of integers where 0 means empty and 1 means occupied
// The function returns true if it is possible, false otherwise
func canPlaceFlowers(flowerbed []int, n int) bool {
	// add two 0s at the beginning and end of the flowerbed
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)

	// iterate through the flowerbed with indices i, j, and k
	i, j, k := 0, 1, 2
	for k < len(flowerbed) {
		// if the current cell is empty and the next two cells are also empty, place a flower
		if flowerbed[i] == 0 && flowerbed[j] == 0 && flowerbed[k] == 0 {
			flowerbed[j] = 1
			n--
		}
		i++
		j++
		k++
	}

	// if n is 0 return true
	if n <= 0 {
		return n <= 0
	}

	// otherwise, return false
	return false
}

// The function canPlaceFlowersAlternative checks whether the specified number of flowers n can be placed in a flowerbed without violating the rules as follows:
func canPlaceFlowersAlternative(flowerbed []int, n int) bool {
	// - If n is equal to 0, it means no flowers need to be placed, so it returns true.
	if n == 0 {
		return true
	}

	// - If n is not equal to 0 and the length of the flowerbed is 1, then flowers can be placed only if the flowerbed is empty.
	if n != 0 && len(flowerbed) == 1 {
		return flowerbed[0] == 0
	}

	// - If n is not equal to 0 and the first two elements of the flowerbed are equal to 0, then a flower is planted in the first element,
	//   and the count of remaining flowers n decreases by 1.
	if n != 0 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
	}

	// - Then there is an iteration over the elements of the flowerbed starting from the second element up to the second-to-last element,
	//   and if the current element and its neighbors on the left and right are equal to 0, a flower is planted in the current element,
	//   and the count of remaining flowers n decreases by 1.
	for i := 1; i < len(flowerbed)-1 && n != 0; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}

	// - If there is still at least one flower remaining (n is not equal to 0) and the last two elements of the flowerbed are equal to 0,
	//   the count of remaining flowers n decreases by 1.
	if n != 0 && flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		n--
	}

	// - It returns true if the count of remaining flowers n becomes 0, otherwise it returns false.
	return n == 0
}
