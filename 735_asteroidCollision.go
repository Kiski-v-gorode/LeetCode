package main

// Function asteroidCollision takes a slice of integers asteroids as input and returns a slice of integers.
func asteroidCollision(asteroids []int) []int {
	// Initialize an empty stack to handle asteroid collisions.
	stack := []int{}

	// Iterate through all asteroids in the input slice.
	for _, asteroid := range asteroids {
		// Handling asteroid collisions:
		// If the stack is not empty, the current asteroid is negative (moving left), and the top asteroid in the stack is positive (moving right),
		// then collision checking occurs.
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			// Compute the difference between the current asteroid and the top asteroid in the stack.
			diff := asteroid + stack[len(stack)-1]
			// If the difference is negative, the top asteroid in the stack is destroyed,
			// so it is removed from the stack.
			if diff < 0 {
				stack = stack[:len(stack)-1]
				// If the difference is positive, the current asteroid is destroyed,
				// so it is not added to the stack, but it is zeroed out.
			} else if diff > 0 {
				asteroid = 0
				// If the difference is zero, both the current asteroid and the top asteroid in the stack are destroyed,
				// so both are removed from the stack, and the current asteroid is zeroed out.
			} else {
				asteroid = 0
				stack = stack[:len(stack)-1]
			}
		}

		// If the current asteroid remains undestroyed (not zeroed out), add it to the stack.
		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}

	// Return the stack after processing all asteroids.
	return stack
}
