package main

import (
	"fmt"
	"sort"
)

func getOperationInput() (string, error) {
	var operation string
	fmt.Print("Введите операцию которую вы хотите произвести (AVG - среднее, SUM - сумма, MED - медиана): ")
	fmt.Scanln(&operation)
	switch operation {
	case "AVG", "SUM", "MED":
		return operation, nil
	default:
		return "", fmt.Errorf("неверная операция")
	}
}

func getNumbersInput() ([]float64, error) {
	var numbers []float64
	var number float64
	fmt.Print("Введите числа через пробел: ")
	for {
		n, err := fmt.Scanf("%f", &number)
		if n == 0 || err != nil {
			break
		}
		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("не введено ни одного числа")
	}

	return numbers, nil
}

func calculate(operation string, numbers []float64) (float64, error) {
	switch operation {
	case "AVG":
		sum := 0.0
		for _, number := range numbers {
			sum += number
		}
		return sum / float64(len(numbers)), nil
	case "SUM":
		sum := 0.0
		for _, number := range numbers {
			sum += number
		}
		return sum, nil
	case "MED":
		sort.Float64s(numbers)
		length := len(numbers)
		if length%2 == 0 {
			return (numbers[length/2-1] + numbers[length/2]) / 2.0, nil
		}
		return numbers[length/2], nil
	default:
		return 0.0, fmt.Errorf("неверная операция")
	}
}

func main() {
	operation, err := getOperationInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	numbers, err := getNumbersInput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result, err := calculate(operation, numbers)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Результат операции %s: %f\n", operation, result)
}
