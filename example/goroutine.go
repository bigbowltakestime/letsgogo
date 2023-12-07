package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%c ", char)
	}
}

func main() {
	// 각 함수를 고루틴으로 실행
	go printNumbers()
	go printLetters()

	// main 함수가 종료되지 않도록 대기
	time.Sleep(7 * time.Second)
}
