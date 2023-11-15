package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Создадим функцию main, в которой будет осуществляться ввод данных и вызов функций для их обработки:

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение например, 2 + 2 или II + II ")

	expression, _ := reader.ReadString('\n')

	expression = strings.TrimSpace(expression)

	result, err := evaluateExpression(expression)

	if err != nil {

		fmt.Println("Ошибка:", err)

	} else {

		fmt.Println("Результат:", result)

	}

}

// Внутри функции main считываем выражение с помощью функции ReadString из пакета bufio.
// Мы также используем функции TrimSpace для удаления лишних пробелов и других символов пунктуации.
// Затем вызываем функцию evaluateExpression, которая будет обрабатывать введенное выражение и возвращать результат.
// Теперь создадим функцию evaluateExpression для обработки выражения:

func evaluateExpression(expression string) (int, error) {

	isArabic := isArabicExpression(expression)

	if isArabic {

		return evaluateArabicExpression(expression)

	}

	return evaluateRomanExpression(expression)

}

// Эта функция сначала вызывает функцию isArabicExpression, которая определяет, является ли выражение арабским числом или римским.
// В зависимости от результата, функция вызывает соответствующую функцию для обработки выражения: evaluateArabicExpression для арабских чисел и evaluateRomanExpression для римских чисел.
// Реализуем функцию isArabicExpression, которая будет проверять, содержит ли выражение только арабские числа:

func isArabicExpression(expression string) bool {

	_, err := strconv.Atoi(expression)

	return err == nil

}

// Мы используем функцию Atoi из пакета strconv, чтобы попытаться преобразовать выражение в число.
// Если преобразование прошло успешно, то это арабское число.
// Теперь создадим функцию evaluateArabicExpression для обработки арабского выражения:

func evaluateArabicExpression(expression string) (int, error) {

	parts := strings.Split(expression, " ")

	if len(parts) != 3 {

		return 0, fmt.Errorf("некорректное выражение")

	}

	firstOperand, err := strconv.Atoi(parts[0])

	if err != nil {

		return 0, fmt.Errorf("некорректное выражение")

	}

	secondOperand, err := strconv.Atoi(parts[2])

	if err != nil {

		return 0, fmt.Errorf("некорректное выражение")

	}

	switch operator := parts[1]; operator {

	case "+":

		return firstOperand + secondOperand, nil

	case "-":

		return firstOperand - secondOperand, nil

	case "*":

		return firstOperand * secondOperand, nil

	case "/":

		return firstOperand / secondOperand, nil

	default:

		return 0, fmt.Errorf("неизвестный оператор")

	}

}

func evaluateRomanExpression(expression string) (int, error) {

	parts := strings.Split(expression, " ")

	if len(parts) != 3 {

		return 0, fmt.Errorf("некорректное выражение")

	}

	firstOperand, err := romanToArabic(parts[0])

	if err != nil {

		return 0, fmt.Errorf("некорректное выражение")

	}

	// В этой функции мы сначала разбиваем выражение на три части, используя пробел в качестве разделителя, и проверяем, что количество частей равно трем.
	// Если нет, то возвращаем ошибку.
	// Затем преобразуем первый и второй операнды в числа, используя функцию Atoi. Если преобразование не удалось, то возвращаем ошибку.
	// Затем проверяем оператор с помощью конструкции switch и выполняем соответствующую операцию.
	// Наконец, создадим функцию evaluateRomanExpression для обработки римского выражения:

	secondOperand, err := romanToArabic(parts[2])

	if err != nil {

		return 0, fmt.Errorf("некорректное выражение")

	}

	switch operator := parts[1]; operator {

	case "+":

		return firstOperand + secondOperand, nil

	case "-":

		return firstOperand - secondOperand, nil

	case "*":

		return firstOperand * secondOperand, nil

	case "/":

		return firstOperand / secondOperand, nil

	default:

		return 0, fmt.Errorf("неизвестный оператор")

	}

}

// В этой функции мы используем функцию romanToArabic, которая преобразует римское число в арабское.
// Затем мы выполняем аналогичные действия для операндов и оператора, что и в функции evaluateArabicExpression.
// Наконец, реализуем функцию romanToArabic:

func romanToArabic(roman string) (int, error) {

	romanNumerals := map[rune]int{

		'V': 5,
		'I': 1,
		'X': 10,
	}
	// TO DO..
	//Разобраться или найти библеотеку для работы с римскими числами..

	var result int

	for i := 0; i < len(roman); i++ {

		if i > 0 && romanNumerals[rune(roman[i])] > romanNumerals[rune(roman[i-1])] {

			result += romanNumerals[rune(roman[i])] - 2*romanNumerals[rune(roman[i-1])]

		} else {

			result += romanNumerals[rune(roman[i])]

		}

	}

	if romanToArabic(result) != roman {

		return 0, fmt.Errorf("некорректное римское число")

	}

	return result, nil

}

// В этой функции мы используем карту romanNumerals, которая содержит соответствия между римскими и арабскими числами.
// Затем мы итерируем по символам в строке roman и складываем соответствующее значение из карты romanNumerals.
// Если результат преобразования из арабского числа в римское не соответствует исходному римскому числу, то возвращаем ошибку.
// TO DO..
