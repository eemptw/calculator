package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArabicNumber(numStr string) (int, error) {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("Ошибка: калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно")
	}
	return num, nil
}

func parseRomanNumber(numStr string) (int, error) {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	prevValue := 0
	for i := len(numStr) - 1; i >= 0; i-- {
		value := romanMap[numStr[i]]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	return result, nil
}

func toRomanNumber(num int) (string, error) {
	if num < 1 {
		return "", fmt.Errorf("Ошибка: в римской системе нет отрицательных чисел.")
	}
	romanMap := []struct {
		Value int
		Symbol string
	}{
		{1000, "M"},
        	{900, "CM"},
        	{500, "D"},
        	{400, "CD"},
        	{100, "C"},
        	{90, "XC"},
        	{50, "L"},
        	{40, "XL"},
        	{10, "X"},
        	{9, "IX"},
        	{5, "V"},
        	{4, "IV"},
        	{1, "I"},
	}
	var romanNumberBuilder strings.Builder
	for _, rm := range romanMap {
		for num >= rm.Value {
			romanNumberBuilder.WriteString(rm.Symbol)
			num -= rm.Value
		}
	}
	return romanNumberBuilder.String(), nil
}

func calculate(a, b int, operator string) (int, error) {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Ошибка: деление на ноль")
		}
		result = a / b
	default:
		return 0, fmt.Errorf("Ошибка: недопустимая операция: %s", operator)
	}
	return result, nil
}

func main() {
	fmt.Println("                       ▄█   ▄█▄    ▄████████     ███        ▄████████                                               ")
    fmt.Println("                      ███ ▄███▀   ███    ███ ▀█████████▄   ███    ███                                               ")
    fmt.Println("                      ███▐██▀     ███    ███    ▀███▀▀██   ███    ███                                               ")
    fmt.Println("                     ▄█████▀      ███    ███     ███   ▀   ███    ███                                               ")
    fmt.Println("                     ▀▀█████▄    ▀███████████     ███     ▀███████████                                              ")
    fmt.Println("                      ███▐██▄     ███    ███     ███       ███    ███                                               ")
    fmt.Println("                      ███ ▀███▄   ███    ███     ███       ███    ███                                               ")
    fmt.Println("                      ███   ▀█▀   ███    █▀     ▄████▀     ███    █▀                                                ")
    fmt.Println("                      ▀                                                                                             ")
    fmt.Println(" ▄████████    ▄████████  ▄█        ▄████████ ███    █▄   ▄█          ▄████████     ███      ▄██████▄     ▄████████  ")
    fmt.Println("███    ███   ███    ███ ███       ███    ███ ███    ███ ███         ███    ███ ▀█████████▄ ███    ███   ███    ███  ")
    fmt.Println("███    █▀    ███    ███ ███       ███    █▀  ███    ███ ███         ███    ███    ▀███▀▀██ ███    ███   ███    ███  ")
    fmt.Println("███          ███    ███ ███       ███        ███    ███ ███         ███    ███     ███   ▀ ███    ███  ▄███▄▄▄▄██▀  ")
    fmt.Println("███        ▀███████████ ███       ███        ███    ███ ███       ▀███████████     ███     ███    ███ ▀▀███▀▀▀▀▀    ")
    fmt.Println("███    █▄    ███    ███ ███       ███    █▄  ███    ███ ███         ███    ███     ███     ███    ███ ▀███████████  ")
    fmt.Println("███    ███   ███    ███ ███▌    ▄ ███    ███ ███    ███ ███▌    ▄   ███    ███     ███     ███    ███   ███    ███  ")
    fmt.Println("████████▀    ███    █▀  █████▄▄██ ████████▀  ████████▀  █████▄▄██   ███    █▀     ▄████▀    ▀██████▀    ███    ███  ")
    fmt.Println("                        ▀                               ▀                                               ███    ███  ")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный формат выражения")
		return
	}

	var a, b int
	var err error

	if (!strings.ContainsAny(parts[0], "IVXLCDM") && strings.ContainsAny(parts[2], "IVXLCDM")) || (strings.ContainsAny(parts[0], "IVXLCDM") && !strings.ContainsAny(parts[2], "IVXLCDM")) {
		fmt.Println("Ошибка: используются одновременно разные системы счисления")
        return
	}

	isRoman := strings.ContainsAny(parts[0], "IVXLCDM") && strings.ContainsAny(parts[2], "IVXLCDM")
	if isRoman {
		a, err = parseRomanNumber(parts[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err = parseRomanNumber(parts[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		a, err = parseArabicNumber(parts[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if (a < 0 || a > 10) || (b < 0 || b > 10) {
		fmt.Println("Ошибка: числа должны быть от 0 до 10 включительно")
        return
	}

	result, err := calculate(a, b, parts[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if isRoman {
		romanResult, err := toRomanNumber(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Резултат:", result)
	}
}
