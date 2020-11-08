// https://stepik.org/lesson/353243/step/9?unit=337227
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type record struct {
	SystemObjectID string `json:"system_object_id"`
	Kod            string
	IsDeleted      string `json:"is_deleted"`
	SignatureDate  string `json:"signature_date"`
	Nomdescr       string
	GlobalID       int `json:"global_id"`
	Idx            string
	Razdel         string
	Name           string
}

func main() {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var records []record

	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Println(err)
		panic(err)
	}

	sum := 0
	for _, item := range records {
		sum += item.GlobalID
	}

	fmt.Println(sum)
}
