// Задание 4. Таблица умножения через функцию.
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

func printTable(n int) {
	fmt.Printf("Таблица умножения для %d:\n", n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %2d = %d\n", n, i, n*i)
	}
}

func main() {
	fmt.Println("=== Таблица умножения ===")
	n := readInt("Введите число: ")
	printTable(n)
}
