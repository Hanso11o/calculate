package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	LOW = "Вывод ошибки, так как строка " +
		"не является математической операцией."
	HIGH = "Вывод ошибки, так как формат математической операции " +
		"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Вывод ошибки, так как используются " +
		"одновременно разные системы счисления."
	DIV = "Вывод ошибки, так как в римской системе " +
		"нет отрицательных чисел."
	ZERO  = "Вывод ошибки, так как в римской системе нет числа 0."
	RANGE = "Калькулятор умеет работать только с арабскими целыми " +
		"числами или римскими цифрами от 1 до 10 включительно"
)

func main() {

	fmt.Println("Введите выражение")
	reader := bufio.NewReader(os.Stdin)

	for {
		console, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		s := strings.ReplaceAll(console, " ", "")

		args := []string
		var arg1 string
		var operand string
		var arg2 string

		switch {
		case operand == "+":
			fmt.Println(arg1 + arg2)
		case operand == "-":
			//some code
		case operand == "*":
			//some code
		case operand == "/":
			//some code
		}
		fmt.Println(s)

	}

}
