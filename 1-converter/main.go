package main

import (
	"errors"
	"fmt"
)

// Константы для конвертации валют
const (
	usdToEur = 0.85
	usdToRub = 90.0
)

// Функция для считывания и проверки ввода валюты
func getCurrencyInput(prompt string) (string, error) {
	var currency string
	fmt.Print(prompt)
	fmt.Scanln(&currency)
	switch currency {
	case "USD", "EUR", "RUB":
		return currency, nil
	default:
		return "", errors.New("неверная валюта")
	}
}

// Функция для считывания и проверки ввода числа
func getAmountInput(prompt string) (float64, error) {
	var amount float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&amount)
	if err != nil || amount <= 0 {
		return 0, errors.New("неверное значение")
	}
	return amount, nil
}

// Функция для расчета конвертации
func convertCurrency(amount float64, fromCurrency, toCurrency string) float64 {
	if fromCurrency == toCurrency {
		return amount
	}

	var amountInUSD float64

	// Конвертация в USD
	switch fromCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdToEur
	case "RUB":
		amountInUSD = amount / usdToRub
	}

	// Конвертация из USD в целевую валюту
	switch toCurrency {
	case "USD":
		return amountInUSD
	case "EUR":
		return amountInUSD * usdToEur
	case "RUB":
		return amountInUSD * usdToRub
	}

	return 0.0
}

func main() {
	fmt.Println("Добро пожаловать в калькулятор валют!")

	var fromCurrency, toCurrency string
	var amount float64
	var err error

	// Ввод исходной валюты
	for {
		fromCurrency, err = getCurrencyInput("Введите исходную валюту (USD, EUR, RUB): ")
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	// Ввод суммы
	for {
		amount, err = getAmountInput("Введите сумму: ")
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	// Ввод целевой валюты
	for {
		toCurrency, err = getCurrencyInput("Введите целевую валюту (USD, EUR, RUB): ")
		if err == nil {
			break
		}
		fmt.Println(err)
	}

	// Конвертация валюты
	convertedAmount := convertCurrency(amount, fromCurrency, toCurrency)

	// Вывод результата
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
