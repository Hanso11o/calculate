package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var args []string
	var action rune
	var mul int
	var div int
	var result string
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Введите выражение:")
		text, _ := reader.ReadString('\n')

		if strings.Contains(text, "+") {
			args = strings.Split(text, "+")
			action = '+'

		} else if strings.Contains(text, "-") {
			args = strings.Split(text, "-")
			action = '-'

		} else if strings.Contains(text, "*") {
			args = strings.Split(text, "*")
			action = '*'

		} else if strings.Contains(text, "/") {
			args = strings.Split(text, "/")
			action = '/'
		} else {
			panic("некорректный знак действия")
		}

		if !strings.HasPrefix(args[0], "\"") && !strings.HasSuffix(args[0], "\"") {
			panic("первым аргументом выражения, подаваемым на вход, должна быть строка")

		}
		if len(args[0]) == 0 || len(args[0]) > 12 {
			panic("калькулятор может принимать строки длиной не более 10 символов")

		} else if len(args[1]) == 0 || len(args[1]) > 12 {
			panic("вторым аргументом выражения, подаваемым на вход, должна быть строка либо целое число")
		}

		for i := 0; i < len(args); i++ {
			args[0] = strings.ReplaceAll(args[0], "\"", "")
			args[1] = strings.ReplaceAll(args[1], "\"", "")
			args[0] = strings.TrimSpace(args[0])
			args[1] = strings.TrimSpace(args[1])

		}

		switch {
		case action == '+':
			result := args[0] + args[1]
			fmt.Printf("\"%v\"\n", result)

		case action == '-':
			if strings.Contains(args[0], args[1]) {
				args[0] = strings.Replace(args[0], args[1], "", 1)
				fmt.Printf("\"%v\"\n", args[0])

			} else {
				fmt.Printf("\"%v\"\n", args[0])
			}
		case action == '*':
			mul, _ = strconv.Atoi(args[1])
			if mul < 1 || mul > 10 {
				panic("калькулятор может принимать на вход числа от 1 до 10 включительно")
			}
			result = ""
			for i := 0; i < mul; i++ {
				result += args[0]
			}
			fmt.Printf("\"%v\"\n", result)
		case action == '/':
			div, _ = strconv.Atoi(args[1])
			if div < 1 || div > 10 {
				panic("калькулятор может принимать на вход числа от 1 до 10 включительно")
			}
			args[0][:len(args0)/div]

		}

	}

}
