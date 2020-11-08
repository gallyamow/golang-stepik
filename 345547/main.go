package main

import (
	"fmt"  // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
    ch1 := make([]chan int, n)
    ch2 := make([]chan int, n)
    
    for i := 0; i < n; i++ {
        ch1[i] = make(chan int)
        ch2[i] = make(chan int)
    }
    
    go func () {
        i1 := 0
        i2 := 0
        
        for {
            select {
                case x1 := <- in1:
                go func(i int, x int) {
                    ch1[i] <- fn(x)
                }(i1, x1)
                i1++
                case x2 := <- in2:
                go func(i int, x int) {
                    ch2[i] <- fn(x)
                }(i2, x2)
                i2++
            }

            if i1 == n && i2 == n {
                break
            }
        }
    }()
    
    go func () {
        for i := 0; i < n; i++ {
            out <- (<-ch1[i]) + (<-ch2[i])
        }
    }()
}