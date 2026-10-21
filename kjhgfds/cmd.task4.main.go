package main

import (
	"fmt"
	"time"
)

func printNumbers(id int, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Worker %d: number %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start lab4")

	go printNumbers(1, 5)
	go printNumbers(2, 5)
	go printNumbers(3, 5)
	go printNumbers(4, 5)
	go printNumbers(5, 5)

	// Здесь мы используем time.Sleep, чтобы "подождать" горутины.
	// Это ненадёжно, потому что мы не знаем, сколько точно времени им нужно.
	// Если они будут работать дольше, программа завершится слишком рано.
	// При 100 мс вместо 1 секунды видны только 0–1 итерация каждого воркера:
	// программа заканчивается раньше, чем горутины успевают доработать.
	time.Sleep(100 * time.Millisecond)

	// Правильный интервал возвращаем обратно — 1 секунда:
	// time.Sleep(1 * time.Second)

	fmt.Println("End lab4")
}
