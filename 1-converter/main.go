package main

import "fmt"

const (
	USDoEUR = 0.87
	USDoRUB = 90.0
)

func main() {
	EURoRUB := USDoRUB / USDoEUR

	fmt.Printf("1 EUR = %.2f RUB\n", EURoRUB)
}
