// Задание 5. Максимум из N чисел.
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
	fmt.Println("=== Максимум и минимум из N чисел ===")
	n := readPositiveInt("Сколько чисел будет введено? N = ")

	first := readFloat("Число 1: ")
	maxVal, minVal := first, first
	maxPos, minPos := 1, 1

	for i := 2; i <= n; i++ {
		x := readFloat(fmt.Sprintf("Число %d: ", i))
		if x > maxVal {
			maxVal, maxPos = x, i
		}
		if x < minVal {
			minVal, minPos = x, i
		}
	}

	fmt.Printf("Максимальное значение: %g (введено %d-м)\n", maxVal, maxPos)
	fmt.Printf("Минимальное значение: %g (введено %d-м)\n", minVal, minPos)
}
