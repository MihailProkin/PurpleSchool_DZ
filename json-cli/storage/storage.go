package storage

import (
	"encoding/json"
	"os"

	"json-cli/bins"
)

// SaveBins сохраняет список Bin в файл в формате JSON
func SaveBins(filename string, binList bins.BinList) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(binList)
}

// LoadBins загружает список Bin из файла в формате JSON
func LoadBins(filename string) (bins.BinList, error) {
	var binList bins.BinList

	file, err := os.Open(filename)
	if err != nil {
		return binList, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&binList)
	if err != nil {
		return binList, err
	}
	return binList, nil
}
