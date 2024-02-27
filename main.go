package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		case "*":
			if arg2 == "1" || arg2 == "2" || arg2 == "3" || arg2 == "4" || arg2 == "5" || arg2 == "6" ||
				arg2 == "7" || arg2 == "8" || arg2 == "9" || arg2 == "10" {
				argToInt, err := strconv.Atoi(arg2)
				if err != nil {
					panic("ошибка конвертации строки в число")
				}

				wordRune1 = []rune(strings.Repeat(arg1, argToInt))
				if len(wordRune1) > 40 {
					wordRune1 := wordRune1[:40]
					fmt.Printf("\"%s...\"", string(wordRune1))

				} else {
					fmt.Printf("%q", string(wordRune1))
				}

			} else {
				panic("некорректное число")

			}

		case "/":
			if arg2 == "1" || arg2 == "2" || arg2 == "3" || arg2 == "4" || arg2 == "5" || arg2 == "6" ||
				arg2 == "7" || arg2 == "8" || arg2 == "9" || arg2 == "10" {
				argToInt, err := strconv.Atoi(arg2)
				if err != nil {
					panic("ошибка конвертации строки в число")
				}
				wordRune1 := wordRune1[:len(wordRune1)/argToInt]
				fmt.Printf("%q", string(wordRune1))

			} else {
				panic("некорректное число")

			}
		default:
			panic("некорректная арифметическая операция")

		}

	} else {
		panic("некорректная длина введенных строк")

	}

}
