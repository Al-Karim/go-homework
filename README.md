# Семинар наставника — базовые конструкции Go

10 практических задач на условные операторы, циклы и функции языка Go.
Каждое задание лежит в отдельной папке.

| № | Задание | Папка |
|---|---------|-------|
| 1 | Калькулятор с проверкой деления на ноль | [`task01_calculator`](task01_calculator) |
| 2 | Анализ числа: знак, чётность, делимость на 3 и 5 | [`task02_number_analysis`](task02_number_analysis) |
| 3 | Сумма и среднее от 1 до N | [`task03_sum_average`](task03_sum_average) |
| 4 | Таблица умножения через функцию `printTable(n)` | [`task04_multiplication_table`](task04_multiplication_table) |
| 5 | Максимум и минимум из N чисел | [`task05_max_min`](task05_max_min) |
| 6 | Количество положительных, отрицательных чисел и нулей | [`task06_count_signs`](task06_count_signs) |
| 7 | Сумма чисел от 1 до N, чётных и кратных 3 | [`task07_filtered_sum`](task07_filtered_sum) |
| 8 | Проверка простого числа `isPrime(n int) bool` | [`task08_prime_check`](task08_prime_check) |
| 9 | Мини-банкомат с проверкой достаточности средств | [`task09_mini_atm`](task09_mini_atm) |
| 10 | Статистика последовательности | [`task10_statistics`](task10_statistics) |

## Запуск

Нужен Go 1.21 или новее. Перейдите в папку задания и запустите файл:

```bash
cd task01_calculator
go run main.go
```

Тесты (для заданий 8 и 10), из папки задания:

```bash
go test main.go main_test.go
```
