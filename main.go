package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inputReplaceAll(input string) string {
	return strings.ReplaceAll(input, "\"", "")

}

func inputSplitByAction(input string) (args []string, action rune) {

	if strings.Contains(input, "-") {
		args = strings.Split(input, "-")
		action = '-'
	}

	if strings.Contains(input, "+") {
		args = strings.Split(input, "\\+")
		action = '+'
	}

	if strings.Contains(input, "*") {
		args = strings.Split(input, "\\*")
		action = '*'
	}

	if strings.Contains(input, "/") {
		args = strings.Split(input, "/")
		action = '/'
	}

	return

}

var args []string
var action rune

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение:")

	for {

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ошибка чтения строки [%v]", err)
			return
		}

		input = inputReplaceAll(input)

		inputSplitByAction(input)

		if action == '-' {
			//some code
		}
		if action == '+' {
			fmt.Println(args[0] + args[1])
		}
		if action == '*' {
			//some code
		}
		if action == '/' {
			//some code
		}

	}

}
