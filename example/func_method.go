package main

import "fmt"

// add 함수 정의
func add(a, b int) int {
	return a + b
}
func add_float(a int, b float64) float64 {
	return float64(a) + b
}

// Person 구조체 정의
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type PokemonWorld struct {
	first_pokemon  string
	second_pokemon string
}

func (pokemon_name PokemonWorld) get_battle() string {
	return pokemon_name.first_pokemon + "와 " + pokemon_name.second_pokemon + "가 결투를 시작했다."
}

// 메소드 정의: Person 구조체에 대한 메소드
func (p Person) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	// 함수 호출
	result := add(3, 4)

	// 결과 출력
	fmt.Println("3 + 4 =", result)
	result_float := add_float(3, 3.5)
	fmt.Println("3 + 3.5 = ", result_float)

	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// 메소드 호출
	fullName := person.GetFullName()

	// 결과 출력
	fmt.Println("Full Name:", fullName)

	first_battle := PokemonWorld{
		first_pokemon:  "치코리타",
		second_pokemon: "리아코",
	}
	start_battle := first_battle.get_battle()
	fmt.Printf(start_battle)
}
