package main

import "fmt"

func main() {
	fmt.Println("google.com/api", test)

	for i := range 200 {
		fmt.Println(i)
	}
}

func test() {
	// api слой
	// []byte -> struct
	// Бизнес-логика
	// struct -> []byte
}
