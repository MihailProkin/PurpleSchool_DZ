package main

import (
	"errors"
	"fmt"
)

// Коэффициенты для конвертации валют
var conversionRates = map[string]float64{
	"USD": 1,
	"EUR": 0.85,
	"RUB": 90.0,
}

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
func convertCurrency(amount float64, fromCurrency, toCurrency string, rates *map[string]float64) float64 {
	if fromCurrency == toCurrency {
		return amount
	}

	// Конвертация в USD
	amountInUSD := amount / (*rates)[fromCurrency]

	// Конвертация из USD в целевую валюту
	return amountInUSD * (*rates)[toCurrency]
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
	convertedAmount := convertCurrency(amount, fromCurrency, toCurrency, &conversionRates)

	// Вывод результата
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
