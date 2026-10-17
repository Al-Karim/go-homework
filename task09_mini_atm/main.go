// Задание 9. Мини-банкомат.
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

func readAmount(prompt string, allowZero bool) float64 {
	for {
		s := strings.Replace(readLine(prompt), ",", ".", 1)
		x, err := strconv.ParseFloat(s, 64)
		switch {
		case err != nil:
			fmt.Println("Ошибка: введите число.")
		case x < 0:
			fmt.Println("Ошибка: сумма не может быть отрицательной.")
		case x == 0 && !allowZero:
			fmt.Println("Ошибка: сумма должна быть больше нуля.")
		default:
			return x
		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("Выберите действие:")
	fmt.Println("  1 — Посмотреть баланс")
	fmt.Println("  2 — Пополнить баланс")
	fmt.Println("  3 — Снять средства")
	fmt.Println("  0 — Выход")
}

func main() {
	fmt.Println("=== Мини-банкомат ===")
	balance := readAmount("Введите начальный баланс: ", true)

	for {
		printMenu()
		choice := readLine("Ваш выбор: ")

		switch choice {
		case "1":
			fmt.Printf("Текущий баланс: %.2f\n", balance)
		case "2":
			amount := readAmount("Сумма пополнения: ", false)
			balance += amount
			fmt.Printf("Баланс пополнен на %.2f. Новый баланс: %.2f\n", amount, balance)
		case "3":
			amount := readAmount("Сумма снятия: ", false)
			if amount > balance {
				fmt.Printf("Недостаточно средств: запрошено %.2f, на счёте %.2f\n", amount, balance)
			} else {
				balance -= amount
				fmt.Printf("Выдано %.2f. Остаток на счёте: %.2f\n", amount, balance)
			}
		case "0":
			fmt.Printf("Итоговый баланс: %.2f. До свидания!\n", balance)
			return
		default:
			fmt.Println("Неизвестная команда. Введите 0, 1, 2 или 3.")
		}
	}
}
