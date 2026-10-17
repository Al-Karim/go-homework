// Задание 1. Калькулятор с проверкой.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		fmt.Println("\nВвод завершён.")
		os.Exit(0)
	}
	return strings.TrimSpace(line)
}

func readFloat(prompt string) float64 {
	for {
		s := strings.Replace(readLine(prompt), ",", ".", 1)
		x, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return x
		}
		fmt.Println("Ошибка: введите число (например, 12 или 3.5).")
	}
}

func readOperator() string {
	for {
		op := readLine("Введите операцию (+, -, *, /): ")
		switch op {
		case "+", "-", "*", "/":
			return op
		}
		fmt.Println("Ошибка: неизвестная операция. Допустимы только +, -, *, /.")
	}
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неизвестная операция %q", op)
	}
}

func main() {
	fmt.Println("=== Калькулятор ===")
	for {
		a := readFloat("Введите первое число: ")
		op := readOperator()
		b := readFloat("Введите второе число: ")

		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("Результат: %g %s %g = %g\n", a, op, b, result)
		}

		answer := strings.ToLower(readLine("Ещё одно вычисление? (y/n): "))
		if answer != "y" && answer != "д" && answer != "да" && answer != "yes" {
			fmt.Println("До свидания!")
			return
		}
		fmt.Println()
	}
}
