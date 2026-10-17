// Задание 7. Сумма только подходящих чисел.
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
	fmt.Println("=== Сумма чисел от 1 до N, которые чётные и делятся на 3 ===")
	n := readPositiveInt("Введите N: ")

	sum := 0
	var matched []string
	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
			matched = append(matched, strconv.Itoa(i))
		}
	}

	if len(matched) == 0 {
		fmt.Printf("В диапазоне от 1 до %d нет подходящих чисел. Сумма: 0\n", n)
		return
	}

	if len(matched) <= 30 {
		fmt.Println("Подходящие числа:", strings.Join(matched, ", "))
	} else {
		fmt.Printf("Подходящих чисел: %d (первые: %s, ...)\n", len(matched), strings.Join(matched[:10], ", "))
	}
	fmt.Printf("Сумма: %d\n", sum)
}
