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
		parts[0] = strings.Trim(parts[0], "\"")
		parts[1] = strings.Trim(parts[1], "\"")
		// parts[2] = strings.Trim(parts[2], "\"")
		// if len(parts) != 3 {
		// 	fmt.Print("не достаточно аргументов")
		// 	return
		// }

		fmt.Println(len(parts))
		fmt.Println(parts[0])
		fmt.Println(parts[1])
		// fmt.Println(parts[2])

	}
}
