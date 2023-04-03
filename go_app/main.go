// Made by Mohamad
// Made on Apr 3 2023
//
// This program calculates the volume of a square based pyramid

package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var length float64
	var width float64
	var height float64
	var volume float64

	// Get input from user
	fmt.Print("Enter the length of the base: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width of the base: ")
	fmt.Scan(&width)
	fmt.Print("Enter the height of the pyramid: ")
	fmt.Scan(&height)

	// Calculate volume
	volume = (length * width * height) / 3

	// Output volume
	fmt.Println("\nThe volume of the pyramid is:", volume, "cm³")

	fmt.Println("\nDone.")
}
