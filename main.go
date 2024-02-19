package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение:")

	var parts []string
	for {
		var input string
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("ошибка чтения строки [%v]", err)
			return

		}

		parts = strings.Split(input, "-")
		if !strings.HasPrefix(parts[0], `"`) || !strings.HasSuffix(parts[0], `"`) {

			panic("первый аргумент должен быть в кавычках")

		}
		parts[0] = strings.Trim(parts[0], `"`)

		fmt.Println(parts[0])
		fmt.Println(parts[1])

	}
}
