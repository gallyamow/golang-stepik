package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
Упражнение 1.7. Вызов функции i o . C o p y ( d s t , s r c ) выполняет чтение s r c и
запись в d s t . Воспользуйтесь ею вместо i o u t i l . R e a d A l l для копирования тела
ответа в поток o s . S t d o u t без необходимости выделения достаточно большого для
хранения всего ответа буфера. Не забудьте проверить, не произошла ли ошибка при
вызове i o . Сору.
Упражнение 1.8. Измените программу f e t c h так, чтобы к каждому аргументу
URL автоматически добавлялся префикс h t t p : / / в случае отсутствия в нем таково­
го. Можете воспользоваться функцией s t r i n g s . H a s P r e f i x .
Упражнение 1.9. Измените программу f e t c h так, чтобы она выводила код состо­
яния HTTP, содержащийся в r e s p . S t a t u s .
*/

func main() {
	fmt.Println(fetch("http://yandex.ru"))

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		fmt.Println(fetch(url))
	}
}

func fetch(url string) (body string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// условие можно писать в ту же строку
	if err = resp.Body.Close(); err != nil {
		return
	}

	body = "--HEADERS--"
	for k := range resp.Header {
		body += k + ": " + resp.Header.Get(k)
	}

	body += "--BODY--" + string(b)
	return
}
