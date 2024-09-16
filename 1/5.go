package main

import "fmt"

func calculateSumAndDifference(a, b float64) (float64, float64) {
	return a + b, a - b
}

func main() {
	var num1, num2 float64

	fmt.Println("Введите первое число с плавающей запятой:")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе число с плавающей запятой:")
	fmt.Scanln(&num2)

	sum, diff := calculateSumAndDifference(num1, num2)

	fmt.Printf("Сумма: %.2f\n", sum)
	fmt.Printf("Разность: %.2f\n", diff)
}
