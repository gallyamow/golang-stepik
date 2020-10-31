// https://stepik.org/lesson/351892/step/13?unit=335849
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type noCsvError struct {
	Path string
}

func (e *noCsvError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Path)
}

const dir = "./task"

func walkFunc(path string, info os.FileInfo, err error) error {-
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	data, err := csv.NewReader(file).ReadAll()
	if err != nil {
		// 1) это создание пустой структуры (ошибка это структура)
		// 2) нужно создать именно ссылку на структуру, так как интерфейс Error работает именно с ней.
		return &noCsvError{Path: path}
	}

	if len(data) == 1 && len(data[0]) == 1 {
		return nil
	}

	fmt.Println(data[4][2])

	return nil
}

func main() {
	if err := filepath.Walk(dir, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}
