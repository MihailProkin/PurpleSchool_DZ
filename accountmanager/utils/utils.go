package utils

import "fmt"

func promptData(prompt string) string {
	var input string
	fmt.Print(prompt + ": ")
	fmt.Scanln(&input)
	return input
}
