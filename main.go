package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Введите выражение: ")

	reader := bufio.NewReader(os.Stdin)
	var args []string
	// var operand string

	for {
		console, err := reader.ReadString('\n')
		if err != nil {
			panic(err)

		}
		console = strings.TrimSpace(console)

		args = strings.Split(console, " ")

	}

}
