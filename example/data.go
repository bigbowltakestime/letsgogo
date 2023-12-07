package main

import "fmt"

func main() {
	// 예제 1: 변수 선언과 초기화
	var name string
	name = "Jin"
	fmt.Println("이름:", name)

	// 예제 2: 변수 선언과 초기화 간소화
	var age_normal int = 29
	age := 29
	fmt.Println("나이:", age)
	fmt.Println("나이:", age_normal)

	// 예제 3: 다수의 변수 선언
	var (
		height float64 = 175.5
		weight int     = 70
	)
	fmt.Printf("키: %.2fcm, 몸무게: %dkg\n", height, weight)

	// 예제 4: 자료형 변환
	var num1 int = 10
	var num2 float64 = float64(num1)
	fmt.Printf("num1: %d, num2: %.2f\n", num1, num2)

	// 예제 5: 상수 선언
	const pi float64 = 3.14
	const pi2 float32 = 3.14
	fmt.Println("원주율:", pi)
	fmt.Println("원주율:", pi2)
}
