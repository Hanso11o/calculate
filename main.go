package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var arg1 string
	var operator string
	var arg2 string

	fmt.Println("Введите выражение:")

	fmt.Scanf("%q%s", &arg1, &operator)

	console := bufio.NewScanner(os.Stdin)
	console.Scan()
	arg2 = console.Text()

	wordRune1 := []rune(arg1)
	wordRune2 := []rune(arg2)

	if arg1 == "" {
		panic("первым аргументом дожна быть строка, неправильный ввод")

	}

	if len(wordRune1) <= 10 && len(wordRune2) <= 12 {
		beginWord := rune(arg2[0])
		endWord := rune(arg2[len(arg2)-1])

		switch operator {
		case "+":
			if string(beginWord) == "" && string(endWord) == "" {
				arg2 = arg2[1 : len(arg2)-1]
				fmt.Printf("\"%s%s\"", arg1, arg2)

			} else {
				panic("неправильный ввод второй строки")
			}

		case "-":
			if string(beginWord) == "" && string(endWord) == "" {
				arg2 = arg2[1 : len(arg2)-1]
				arg1 = strings.ReplaceAll(arg1, arg2, "")
				fmt.Printf("%q", arg1)
			} else {
				panic("неправильный ввод второй строки")
			}

		}

	}

}
