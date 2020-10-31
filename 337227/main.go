package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type student struct {
	LastName  string
	FirstName string
	Birthday  string
	Address   string
	Phone     string
	Rating    []int
}

type group struct {
	ID       int
	Number   string
	Year     int
	Students []student
}

type result struct {
	Average float64
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)

	// file, err := os.Open("text.json")
	// data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	var g group

	if err := json.Unmarshal(data, &g); err != nil {
		fmt.Println(err)
		panic(err)
	}

	sum := 0
	cnt := 0
	for _, sItem := range g.Students {
		cnt++
		sum += len(sItem.Rating)
	}

	res := result{Average: float64(sum) / float64(cnt)}

	js, _ := json.MarshalIndent(res, "", "    ")

	io.WriteString(os.Stdout, string(js))
}
