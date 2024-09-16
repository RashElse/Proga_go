package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Введите первое целое число:")
	fmt.Scanln(&num1)

	fmt.Println("Введите второе целое число:")
	fmt.Scanln(&num2)

	fmt.Println("Результаты арифметических операций:")
	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	if num2 != 0 {
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	} else {
		fmt.Println("Деление на ноль невозможно!")
	}
}
