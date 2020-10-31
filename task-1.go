package main

import (
	"fmt"
	"math/cmplx"
	"os"
	"strings"
	"time"
)

/*
Упражнение 1.1. Измените программу e c h o так, чтобы она выводила также
o s . A r g s [ 0 ] , имя выполняемой команды.
Упражнение 1.2. Измените программу e c h o так, чтобы она выводила индекс и
значение каждого аргумента по одному аргументу в строке.
Упражнение 1.3. Поэкспериментируйте с измерением разницы времени выполне­
ния потенциально неэффективных версий и версии с применением s t r i n g s . Doin.
(В разделе 1.6 демонстрируется часть пакета tim e , а в разделе 11.4 — как написать
тест производительности для ее систематической оценки.)
*/
func main() {
	var start = time.Now()
	fmt.Println("For-version")
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	// или без var
	elapsed := start.Sub(time.Now())
	fmt.Println("elapsed: ", elapsed)

	start = time.Now()
	fmt.Println("Join-version")
	fmt.Println(strings.Join(os.Args, " "))
	elapsed = start.Sub(time.Now())
	fmt.Println("elapsed: ", elapsed)

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
