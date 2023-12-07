package main

import "fmt"

func main() {
	// 예제 1: 기본적인 if, else 문
	age := 25
	if age >= 18 {
		fmt.Println("당신은 성인입니다.")
	} else {
		fmt.Println("당신은 미성년자입니다.")
	}
	start_pokemon := "치코리타"
	fmt.Println("당신의 포켓몬 : ", start_pokemon)
	if start_pokemon == "치코리타" {
		fmt.Println("당신의 포켓몬은 구립니다")
	} else {
		fmt.Println("당신의 포켓몬은 좋습니다")
	}

	// 예제 2: 조건문에서 간단한 문장
	if x := 10; x > 5 {
		fmt.Println("x는 5보다 큽니다.")
	} else {
		fmt.Println("x는 5보다 작거나 같습니다.")
	}

	// 예제 1: 기본적인 for 반복문
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for k := 0; k < 5; k++ {
		fmt.Println("에~ 치코리타 골랐대요")
	}
	// 예제 2: 조건식만 사용하는 무한 루프
	j := 0
	for {
		if j >= 5 {
			break
		}
		fmt.Println(j)
		j++
	}

	// 예제 3: 배열 또는 슬라이스를 순회하는 for 반복문
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("인덱스: %d, 값: %d\n", index, value)
	}

	numbers_2 := []int{1, 2, 3, 4}

	// 슬라이스 순회
	for index, value := range numbers_2 {
		fmt.Printf("인덱스: %d, 값: %d\n", index, value)
	}
	myMap := make(map[string]int)

	// 맵에 데이터 추가
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	// 맵 순회
	for key, value := range myMap {
		fmt.Printf("키: %s, 값: %d\n", key, value)
	}
	myMap_2 := map[string]string{"브케인": "불", "리아코": "물", "치코리타": "구리다"}
	for key, value := range myMap_2 {
		fmt.Printf("키: %s, 값: %d\n", key, value)
	}
	arr1 := [3]int{1, 2, 3}
	slice1 := []int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(slice1)

}
