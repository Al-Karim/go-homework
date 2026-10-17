// Задание 2. Анализ числа.
package main

import (
	"bufio"
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

func readInt(prompt string) int {
	for {
		n, err := strconv.Atoi(readLine(prompt))
		if err == nil {
			return n
		}
		fmt.Println("Ошибка: введите целое число.")
	}
}

func main() {
	fmt.Println("=== Анализ числа ===")
	n := readInt("Введите целое число N: ")

	if n > 0 {
		fmt.Printf("%d — положительное число\n", n)
	} else if n < 0 {
		fmt.Printf("%d — отрицательное число\n", n)
	} else {
		fmt.Println("0 — ни положительное, ни отрицательное число")
	}

	if n%2 == 0 {
		fmt.Printf("%d — чётное число\n", n)
	} else {
		fmt.Printf("%d — нечётное число\n", n)
	}

	divBy3 := n%3 == 0
	divBy5 := n%5 == 0

	if divBy3 {
		fmt.Printf("%d делится на 3\n", n)
	} else {
		fmt.Printf("%d не делится на 3\n", n)
	}
	if divBy5 {
		fmt.Printf("%d делится на 5\n", n)
	} else {
		fmt.Printf("%d не делится на 5\n", n)
	}

	if divBy3 && divBy5 {
		fmt.Printf("%d делится и на 3, и на 5 (то есть на 15)\n", n)
	} else {
		fmt.Printf("%d не делится одновременно на 3 и на 5\n", n)
	}
}
