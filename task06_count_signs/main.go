// Задание 6. Количество положительных и отрицательных чисел.
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

func readPositiveInt(prompt string) int {
	for {
		n, err := strconv.Atoi(readLine(prompt))
		if err == nil && n >= 1 {
			return n
		}
		fmt.Println("Ошибка: введите целое число не меньше 1.")
	}
}

func readFloat(prompt string) float64 {
	for {
		s := strings.Replace(readLine(prompt), ",", ".", 1)
		x, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return x
		}
		fmt.Println("Ошибка: введите число.")
	}
}

func main() {
	fmt.Println("=== Подсчёт положительных, отрицательных чисел и нулей ===")
	n := readPositiveInt("Сколько чисел будет введено? N = ")

	positive, negative, zeros := 0, 0, 0
	for i := 1; i <= n; i++ {
		x := readFloat(fmt.Sprintf("Число %d: ", i))
		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println("Результат:")
	fmt.Printf("  положительных: %d\n", positive)
	fmt.Printf("  отрицательных: %d\n", negative)
	fmt.Printf("  нулей:         %d\n", zeros)
}
