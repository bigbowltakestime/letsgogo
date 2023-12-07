package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	// 채널에 데이터 전송
	ch <- "Hello"
	time.Sleep(1 * time.Second)
	ch <- "Go"
	time.Sleep(1 * time.Second)
	ch <- "Routine"
}

func main() {
	// 채널 생성
	messageChannel := make(chan string)

	// 고루틴에서 데이터 전송
	go sendData(messageChannel)

	// 메인 고루틴에서 데이터 수신
	for i := 0; i < 3; i++ {
		msg := <-messageChannel
		fmt.Println(msg)
	}
}
