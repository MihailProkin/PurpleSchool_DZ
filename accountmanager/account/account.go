package account

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

// Метод для вывода информации о пароле
func (acc account) outputPassword() {
	fmt.Println("Login:", acc.login, "| Password:", acc.password, "| URL:", acc.url)
}

// Метод для генерации пароля
func (acc *account) generatePassword(n int) {
	if n < 4 {
		fmt.Println("Длина пароля должна быть не менее 4 символов!")
		return
	}

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	res := make([]rune, n)

	// Добавляем по одному символу каждого типа
	res[0] = lowerRunes[rand.Intn(len(lowerRunes))]
	res[1] = upperRunes[rand.Intn(len(upperRunes))]
	res[2] = digitRunes[rand.Intn(len(digitRunes))]
	res[3] = specialRunes[rand.Intn(len(specialRunes))]

	// Заполняем оставшиеся позиции случайными символами
	for i := 4; i < n; i++ {
		res[i] = allRunes[rand.Intn(len(allRunes))]
	}

	// Перемешиваем результат, чтобы порядок символов был случайным
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })

	// Преобразуем массив в строку и присваиваем его полю password
	acc.password = string(res)
}

// Функция для создания нового аккаунта
func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, fmt.Errorf("логин и URL обязательны")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("неверный URL: %w", err)
	}
	acc := &account{login: login, url: urlString}
	if password == "" { // Если пароль не введен, генерируем новый
		acc.generatePassword(12)
	} else {
		acc.password = password
	}
	return acc, nil
}

var (
	lowerRunes   = []rune("abcdefghijklmnopqrstuvwxyz")
	upperRunes   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitRunes   = []rune("0123456789")
	specialRunes = []rune("!@#$%^&*()-_=+[]{}|;:,.<>?/`~")
	allRunes     = append(append(append(lowerRunes, upperRunes...), digitRunes...), specialRunes...)
)
