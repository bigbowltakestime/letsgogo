package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "Hello"
	time.Sleep(1 * time.Second)
	ch <- "Go"
	time.Sleep(1 * time.Second)
	ch <- "Routine"
	time.Sleep(1 * time.Second)
	ch <- "이거 진짜 대박이다"
	close(ch) // 채널 닫기
}

func receiveData(ch chan string) {
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

func main() {
	// 채널 생성
	messageChannel := make(chan string)

	// sendData 함수를 고루틴으로 실행
	go sendData(messageChannel)

	// receiveData 함수를 고루틴으로 실행
	go receiveData(messageChannel)

	// main 함수가 종료되지 않도록 대기
	time.Sleep(5 * time.Second)
}
