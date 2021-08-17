package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var b string

	fmt.Println("Введите первое число")
	fmt.Scanln(&a)
	aInt, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", aInt)

	fmt.Println("Ввидите второе число")
	fmt.Scanln(&b)
	bInt, errInt := strconv.Atoi(b)
	if errInt != nil {
		panic(errInt)
	}
	fmt.Printf("Parsed integer: %d\n", bInt)

	fmt.Println(aInt + bInt)
}
