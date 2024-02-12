package main

import (
	"fmt"
	"strings"
)

func main() {

	// fmt.Println("Введите выражение: ")

	// reader := bufio.NewReader(os.Stdin)
	// var args []string

	// for {

	// }

	const refString = "Mary had a little lamb"
	const refStringTwo = "lamb lamb lamb lamb"

	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)

}
