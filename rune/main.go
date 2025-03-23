package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type textAnalyzer struct {
	wordCount int
	lineCount int
	characterCount int
	upperCount int
	lowerCount int
}

func NewTextAnalyzer() *textAnalyzer {
	return &textAnalyzer{}
}

func (t *textAnalyzer) Analyze(text string) {
	//  count the number of runes in the string
	t.lineCount = utf8.RuneCountInString(text) - utf8.RuneCountInString(strings.ReplaceAll(text, "\n", "")) //count the lines in our text
	t.characterCount = utf8.RuneCountInString(text)

	t.wordCount = len(strings.Fields(text)) // count the words in our text
	for _, r := range text {
		if unicode.IsUpper(r) {
			t.upperCount++
		} else if unicode.IsLower(r) {
			t.lowerCount++
		}
	}
}

func (t textAnalyzer) PrintsStats() {
	fmt.Printf("Word count: %d\n", t.wordCount)
	fmt.Printf("Line count: %d\n", t.lineCount)
	fmt.Printf("Character count: %d\n", t.characterCount)
	fmt.Printf("Upper count: %d\n", t.upperCount)
	fmt.Printf("Lower count: %d\n", t.lowerCount)
}

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
		fmt.Printf("Index: %d, Rune: '%c' - %d, Unicode: U+%04X\n", i, r, r, r)
	}

	// Length of string vs length of runes
	fmt.Printf("\nString length (bytes): %d\n", len(str))
	fmt.Printf("Number of runes: %d\n", len(runes))
	fmt.Println()
	{
		text := `Hello 世
						my friend ã`
		a := NewTextAnalyzer()
		a.Analyze(text)
		a.PrintsStats()
}
}
