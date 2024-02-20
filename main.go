package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func inputReplaceAll(input string) string {
// 	return strings.ReplaceAll(input, "\"", "")

// }

// func inputSplitByAction(input string) (args []string, action rune) {

// 	if strings.Contains(input, "-") {
// 		args = strings.Split(input, "-")
// 		action = '-'
// 	}

// 	if strings.Contains(input, "+") {
// 		args = strings.Split(input, "\\+")
// 		action = '+'
// 	}

// 	if strings.Contains(input, "*") {
// 		args = strings.Split(input, "\\*")
// 		action = '*'
// 	}

// 	if strings.Contains(input, "/") {
// 		args = strings.Split(input, "/")
// 		action = '/'
// 	}

// 	return

// }

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
			result = fmt.Sprintf("\"%s\"", args[0]+args[1])
			fmt.Println(result)
		}
		if action == '*' {
			//some code
		}
		if action == '/' {
			//some code
		}

	}

}
