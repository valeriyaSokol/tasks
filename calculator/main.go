package main

import (
	"fmt"
	"strconv"
)

func getOperator(text string) string {
	var operator string

	fmt.Println(text)
	fmt.Scanln(&operator)

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Неверное значение")
		return getOperator(text)
	}
	return operator
}

func main() {

	aInt := getNumber("Введите первое число")

	operator := getOperator("Введите: + - * /")

	bInt := getNumber("Введите второе число")

	switch operator {
	case "+":
		fmt.Println(aInt + bInt)
	case "-":
		fmt.Println(aInt - bInt)
	case "/":
		fmt.Println(aInt / bInt)
	case "*":
		fmt.Println(aInt * bInt)
	default:
		fmt.Println(operator)
	}
}

func getNumber(text string) int {
	var numberString string

	fmt.Println(text)
	fmt.Scanln(&numberString)

	numberInt, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println("Неверное значение")
		return getNumber(text)
	}

	return numberInt
}
