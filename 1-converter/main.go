package main

import "fmt"

const (
	USDoEUR = 0.87
	USDoRUB = 90.0
)

func readInput() {
	var amount float64
	var sourceCurrency string
	var targetCurrency string

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&sourceCurrency)

	fmt.Print("Введите целевую валюту (USD, EUR, RUB): ")
	fmt.Scan(&targetCurrency)

	convertCurrency(amount, sourceCurrency, targetCurrency)
}

func convertCurrency(amount float64, sourceCurrency, targetCurrency string) {
}

func main() {
	EURoRUB := USDoRUB / USDoEUR

	fmt.Printf("1 EUR = %.2f RUB\n", EURoRUB)
}
