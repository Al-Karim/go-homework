// Задание 10. Статистика последовательности.
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

func readPositiveInt(prompt string) int {
	for {
		n := readInt(prompt)
		if n >= 1 {
			return n
		}
		fmt.Println("Ошибка: количество должно быть не меньше 1.")
	}
}

type Stats struct {
	Sum     int
	Average float64
	Max     int
	Min     int
	Even    int
	Odd     int
}

func computeStats(numbers []int) Stats {
	s := Stats{Max: numbers[0], Min: numbers[0]}
	for _, x := range numbers {
		s.Sum += x
		if x > s.Max {
			s.Max = x
		}
		if x < s.Min {
			s.Min = x
		}
		if x%2 == 0 {
			s.Even++
		} else {
			s.Odd++
		}
	}
	s.Average = float64(s.Sum) / float64(len(numbers))
	return s
}

func main() {
	fmt.Println("=== Статистика последовательности ===")
	n := readPositiveInt("Сколько чисел будет введено? N = ")

	numbers := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, readInt(fmt.Sprintf("Число %d: ", i)))
	}

	s := computeStats(numbers)

	fmt.Println()
	fmt.Println("Введённые числа:", numbers)
	fmt.Println("----------------------------")
	fmt.Printf("Сумма:              %d\n", s.Sum)
	fmt.Printf("Среднее значение:   %.2f\n", s.Average)
	fmt.Printf("Максимум:           %d\n", s.Max)
	fmt.Printf("Минимум:            %d\n", s.Min)
	fmt.Printf("Чётных чисел:       %d\n", s.Even)
	fmt.Printf("Нечётных чисел:     %d\n", s.Odd)
}
