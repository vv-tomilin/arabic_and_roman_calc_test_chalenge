package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isArabic(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func isRoman(input string) bool {
	romanNumerals := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}

	_, ok := romanNumerals[input]
	return ok
}

// * функция перевода результата математической операции из арабского в римское число
func arabicToRoman(number int) string {
	romanSymbols := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	//* Упорядоченный слайс арабских символов
	arabicSymbols := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	roman := ""
	for _, arabic := range arabicSymbols {
		for number >= arabic {
			roman += romanSymbols[arabic]
			number -= arabic
		}
	}

	return roman
}

func returnResultForCalculate(result int, typeExpression string) string {
	if typeExpression == "arabic" {
		return strconv.Itoa(result)
	} else if typeExpression == "roman" {
		if result <= 0 {
			panic(fmt.Errorf("Вывод ошибки, так как в римской системе нет отрицательных чисел или 0."))
		}
		return arabicToRoman(result)
	}

	return ""
}

func calculate(operandOne, operandTwo int, operator string, typeExpression string) (string, error) {
	switch operator {
	case "+":
		result := operandOne + operandTwo
		return returnResultForCalculate(result, typeExpression), nil
	case "-":
		result := operandOne - operandTwo
		return returnResultForCalculate(result, typeExpression), nil
	case "*":
		result := operandOne * operandTwo
		return returnResultForCalculate(result, typeExpression), nil
	case "/":
		result := operandOne / operandTwo
		return returnResultForCalculate(result, typeExpression), nil
	default:
		return "", fmt.Errorf("Строка не является математической операцией: %s", operator)
	}
}

func isNumeralInRange(arabicNumeral int) bool {
	return arabicNumeral >= 1 && arabicNumeral <= 10
}

func romanToArabicForToken(romanNumeral string) int {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	arabicNumeral, _ := romanNumerals[romanNumeral]

	return arabicNumeral
}

func parseExpression(expression string) (int, int, string, string) {
	pattern := `[+\-*/]`
	reg := regexp.MustCompile(pattern)
	matches := reg.FindAllString(expression, -1)

	//* проверяем один ли оператор в выражении с помощью регулярки
	if len(matches) == 1 {
		parts := strings.Split(expression, "")

		var operandOne, operandTwo int
		var operator string
		var typeExpression string

		for _, part := range parts {
			if part == "+" || part == "-" || part == "*" || part == "/" {
				operator = part
			}
		}

		if operator != "" {
			splitOperands := strings.Split(expression, operator)

			tokenOne := splitOperands[0]
			tokenTwo := splitOperands[1]

			//* проверяем в выражении только арабские либо римские
			if (isArabic(tokenOne) && isArabic(tokenTwo)) || (isRoman(tokenOne) && isRoman(tokenTwo)) {

				if isArabic(tokenOne) && isArabic(tokenTwo) {
					numOne, errOne := strconv.Atoi(tokenOne)
					numTwo, errTwo := strconv.Atoi(tokenTwo)

					if isNumeralInRange(numOne) && isNumeralInRange(numTwo) {

						if errOne != nil && errTwo != nil {
							operandOne = 0
							operandTwo = 0
						} else {
							operandOne = numOne
							operandTwo = numTwo
							typeExpression = "arabic"
						}
					} else {
						panic(fmt.Errorf("Вывод ошибки, операнды должны быть в диапазоне от 1 - 10."))
					}
				}

				if isRoman(tokenOne) && isRoman(tokenTwo) {
					numOne := romanToArabicForToken(tokenOne)
					numTwo := romanToArabicForToken(tokenTwo)

					if isNumeralInRange(numOne) && isNumeralInRange(numTwo) {
						operandOne = numOne
						operandTwo = numTwo
						typeExpression = "roman"
					}
				}
			} else {
				//* если выражение несорректное проверяем есть ли в выражении арабское либо римское число
				if isArabic(tokenOne) || isArabic(tokenTwo) || isRoman(tokenOne) || isRoman(tokenTwo) {

					if isArabic(tokenOne) {
						if isRoman(tokenTwo) {
							panic(fmt.Errorf("Вывод ошибки, так как используются одновременно разные системы счисления."))
						} else {
							panic(fmt.Errorf("Вывод ошибки, так как в выражении некорректный операнд."))
						}
					} else if isArabic(tokenTwo) {
						if isRoman(tokenOne) {
							panic(fmt.Errorf("Вывод ошибки, так как используются одновременно разные системы счисления."))
						} else {
							panic(fmt.Errorf("Вывод ошибки, так как в выражении некорректный операнд."))
						}
					} else {
						panic(fmt.Errorf("Вывод ошибки, так как в выражении некорректный операнд."))
					}

				} else {
					panic(fmt.Errorf("Вывод ошибки, так как в выражении некорректное выражение (можно использовать только арабские числа [1-10] либо римские [I-X])."))
				}
			}

		} else {
			panic(fmt.Errorf("Вывод ошибки, так как строка не является математической операцией."))
		}

		return operandOne, operandTwo, operator, typeExpression
	} else {
		panic(fmt.Errorf("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	experssion, _ := reader.ReadString('\n')
	experssion = strings.TrimSpace(experssion)

	numOne, numTwo, operator, typeExp := parseExpression(experssion)
	result, _ := calculate(numOne, numTwo, operator, typeExp)
	fmt.Println(result)
}
