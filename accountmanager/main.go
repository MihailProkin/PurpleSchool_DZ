package main

import (
	"fmt"

	"github.com/mihailprokin/PurpleSchool_DZ/accountmanager/account"
	"github.com/mihailprokin/PurpleSchool_DZ/accountmanager/utils"
)

// Функция для ввода данных с консоли

func main() {
	// Ввод данных с консоли
	login := promptData("Введите логин")
	password := promptData("Введите пароль (оставьте пустым для генерации)")
	url := promptData("Введите URL")

	// Создание аккаунта и вывод данных
	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Printf("Ошибка при создании аккаунта: %v\n", err)
		return
	}

	myAccount.outputPassword()
}
