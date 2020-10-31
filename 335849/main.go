package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./task.data.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл.")
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	buf := make([]byte, 10)

	index := 0
	found := false

	word := ""

	for {
		n, err := rd.Read(buf)
		// word, err := rd.ReadString(';')

		if err != nil && err != io.EOF {
			panic(err)
		}

		// интересный момент, иногда возвращает 6 вместо 10
		// наверно это когда переполняется буфер
		// а buf остается предзаполнен предыдущими данными
		if n != 10 {
			fmt.Println(n)
		}
		// fmt.Println(word)

		// index++

		// if word == "0;" {
		// 	break
		// }

		for i := 0; i < n; i++ {
			// так будет работать неправильно
			// for _, v := range buf {
			s := string(buf[i])

			if s == ";" {
				index++

				//fmt.Println(word, index, n)
				if word == "0" {
					found = true
					break
				}
				word = ""
			} else {
				word += s
			}
		}

		if err == io.EOF || found {
			break
		}
	}

	fmt.Println(index)

}
