package main

import (
	"fmt"
)

func main() {
	// Define a string
	str := "Hello, 世界"

	// Iterate over the string as runes
	fmt.Println("Iterating over string as runes:")
	for index, r := range str {
		fmt.Printf("Index: %d, Rune: '%c', Unicode: U+%04X\n", index, r, r)
	}

	// Convert a rune to a string
	r := '世' // Rune for the character 世
	runeToString := string(r)
	fmt.Printf("\nRune to String: Rune '%c' -> String \"%s\"\n", r, runeToString)

	// Convert a string to runes
	fmt.Println("\nString to Runes:")
	runes := []rune(str)
	for i, r := range runes {
		fmt.Printf("Index: %d, Rune: '%c', Unicode: U+%04X\n", i, r, r)
	}

	// Length of string vs length of runes
	fmt.Printf("\nString length (bytes): %d\n", len(str))
	fmt.Printf("Number of runes: %d\n", len(runes))
}
