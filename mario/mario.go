package mario

import "fmt"

func Mario(height int) {
	if height <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer for the pyramid height.")
		return
	}

	for i := 1; i <= height; i++ {
		// Print spaces to right-align the hashes
		for j := 1; j <= height-i; j++ {
			fmt.Print(" ")
		}

		// Print hashes
		for k := 1; k <= i; k++ {
			fmt.Print("#")
		}

		fmt.Println() // Move to the next line for the next row
	}
}
