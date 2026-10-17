// Задание 3. Сумма и среднее от 1 до N.
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

func main() {
	fmt.Println("=== Сумма и среднее от 1 до N ===")
	n := readPositiveInt("Введите N: ")

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	average := float64(sum) / float64(n)

	fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", average)

	if sum == n*(n+1)/2 {
		fmt.Println("Проверка по формуле n(n+1)/2: совпадает ✔")
	}
}
