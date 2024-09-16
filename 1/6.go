package main

import "fmt"

func calculateAverage(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	var num1, num2, num3 float64

	fmt.Println("Введите три числа для вычисления среднего значения:")
	fmt.Scanln(&num1, &num2, &num3)

	average := calculateAverage(num1, num2, num3)

	fmt.Printf("Среднее значение: %.2f\n", average)
}
