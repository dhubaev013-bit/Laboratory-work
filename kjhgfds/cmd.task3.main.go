package main

import (
	"fmt"
	"time"
)

func main() {
	const goroutines = 1000 // попробуйте увеличить до 10000
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d says hello\n", id)
		}(i) // передаём i параметром — иначе все выведут одно значение
	}

	time.Sleep(500 * time.Millisecond)
}
