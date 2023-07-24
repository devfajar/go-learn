package main

import "fmt"

func main() {
	messageFromDoris := []string{
		"You doing anything later??",
		"Did you get my last message",
		"Don't leave me hanging...",
		"Please respond i'm lonely!",
	}
	numMessages := float64(len(messageFromDoris))
	costPerMessage := .02

	// Dont touch above this line
	totalCost := costPerMessage * numMessages

	// Dont touch below this line

	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)
}
