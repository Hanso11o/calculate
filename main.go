package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var args []string
var action rune
var result string

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение:")

	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ошибка чтения строки [%v]", err)
			return
		}

		input = strings.ReplaceAll(input, "\"", "")

		input = strings.TrimSpace(input)

		if strings.Contains(input, "-") {
			args = strings.Split(input, "-")
			action = '-'
		}

		if strings.Contains(input, "+") {
			args = strings.Split(input, "+")
			action = '+'

		}

		if strings.Contains(input, "*") {
			args = strings.Split(input, "*")
			action = '*'
		}

		if strings.Contains(input, "/") {
			args = strings.Split(input, "/")
			action = '/'
		}

		if action == '-' {
			//some code
		}
		if action == '+' {

			args[1] = strings.TrimSpace(args[1])
			result = args[0] + args[1]
			result = strings.ReplaceAll(result, " ", "")
			fmt.Printf("\"%v\"", result)
		}
		if action == '*' {
			//some code
		}
		if action == '/' {
			//some code
		}

	}

}
