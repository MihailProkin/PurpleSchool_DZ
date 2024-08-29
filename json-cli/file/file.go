package file

import (
	"os"
	"strings"
)

// ReadFile читает содержимое файла и возвращает его в виде строки
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// IsJSONFile проверяет, является ли файл JSON файлом по расширению
func IsJSONFile(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".json")
}
