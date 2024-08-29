package main

import (
	"fmt"

	"json-cli/bins"
	"json-cli/file"
	"json-cli/storage"
)

func main() {
	bin := bins.NewBin("123", true, "example.json")

	binList := bins.NewBinList()
	binList.Bins = append(binList.Bins, bin)

	// Сохранение списка Bin в файл
	err := storage.SaveBins("binlist.json", binList)
	if err != nil {
		fmt.Println("Ошибка сохранения bin:", err)
		return
	}

	// Чтение списка Bin из файла
	loadedBinList, err := storage.LoadBins("binlist.json")
	if err != nil {
		fmt.Println("Ошибка чтения bin:", err)
		return
	}
	fmt.Printf("Загруженные Bins: %+v\n", loadedBinList)

	// Проверка, является ли файл JSON
	isJSON := file.IsJSONFile("binlist.json")
	fmt.Printf("Файл binlist.json является JSON: %t\n", isJSON)

	// Чтение содержимого файла
	content, err := file.ReadFile("binlist.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Printf("Содержимое файла binlist.json: %s\n", content)
}
