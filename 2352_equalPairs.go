package main

import "strconv"

// equalPairs function takes a 2D grid of integers and returns the number of pairs of rows and columns that have equal values in the grid.
func equalPairs(grid [][]int) int {
	// Initialize the result variable to store the count of equal pairs.
	result := 0

	// Initialize two maps to store the counts of rows and columns as strings.
	rows := map[string]int{}
	columns := map[string]int{}

	// Iterate over each row of the grid.
	for i := 0; i < len(grid); i++ {
		// Initialize an empty string to represent the current row.
		row := ""
		// Initialize an empty string to represent the current column.
		column := ""

		// Iterate over each element in the row.
		for j := 0; j < len(grid); j++ {
			// Append the integer value of the element to the row string, separated by a space.
			row += strconv.Itoa(grid[i][j]) + " "
			// Append the integer value of the element in the corresponding column to the column string, separated by a space.
			column += strconv.Itoa(grid[j][i]) + " "
		}

		// Increment the count of the current row string in the rows map.
		rows[row]++
		// Increment the count of the current column string in the columns map.
		columns[column]++
	}

	// Iterate over the rows map.
	for key, value := range rows {
		// Increment the result by the product of the count of the current row string and the count of the corresponding column string.
		result += value * columns[key]
	}

	// Return the total count of equal pairs.
	return result
}
