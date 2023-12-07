package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Job 구조체는 작업을 나타냅니다.
type Job struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

// Worker 함수는 워커의 역할을 하는 함수입니다.
func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// 각 작업을 처리하는 시뮬레이션
		time.Sleep(2 * time.Second)

		// 결과 채널에 결과를 전송
		results <- fmt.Sprintf("Worker %d processed Job %d: %s", id, job.ID, job.Msg)
	}
}

func main() {
	const numWorkers = 3

	// 작업을 보관하는 채널과 결과를 받는 채널 생성
	jobs := make(chan Job, 100)
	results := make(chan string, 100)

	// WaitGroup 생성
	var wg sync.WaitGroup

	// 워커들 시작
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Fiber 애플리케이션 생성
	app := fiber.New()

	// 작업을 큐에 추가하는 엔드포인트
	app.Post("/enqueue", func(c *fiber.Ctx) error {
		var job Job
		if err := c.BodyParser(&job); err != nil {
			return err
		}

		// 작업을 큐에 추가
		jobs <- job
		return c.SendStatus(fiber.StatusAccepted)
	})

	// 작업 결과를 확인하는 엔드포인트
	app.Get("/results", func(c *fiber.Ctx) error {
		var resultStrings []string

		// 모든 결과를 수집
		close(results)
		for result := range results {
			resultStrings = append(resultStrings, result)
		}

		return c.JSON(resultStrings)
	})

	// 서버 시작
	go func() {
		// 모든 워커가 작업을 처리할 때까지 대기
		wg.Wait()
		close(jobs)
	}()

	// 30000번 포트에서 웹 서버 시작
	err := app.Listen(":30000")
	if err != nil {
		panic(err)
	}
}
