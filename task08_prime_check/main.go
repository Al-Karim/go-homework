// Задание 8. Проверка простого числа.
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

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func smallestDivisor(n int) int {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func main() {
	fmt.Println("=== Проверка простого числа ===")
	n := readInt("Введите целое число: ")

	if isPrime(n) {
		fmt.Printf("%d — простое число\n", n)
		return
	}

	if n < 2 {
		fmt.Printf("%d — не простое число (простые числа начинаются с 2)\n", n)
	} else {
		d := smallestDivisor(n)
		fmt.Printf("%d — составное число: %d = %d x %d\n", n, n, d, n/d)
	}
}
