package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	//calc()
	fmt.Println(simpleDimple())

}

/* 1. Доработать калькулятор из методички: больше операций и валидации данных (проверка деления на 0, возведение в степень, факториал) + форматирование строки при выводе дробного числа ( 2 знака после точки)*/
func calc() {

	var a, b, res float64
	var resN float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow, fact): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Printf("Не будем рассматривать здесь пределы и интегральное вычисление...Делить на ноль не будем:)")
			return
		} else {
			res = a / b
		}
	case "pow":
		if b < 0 {
			res = 1 / math.Pow(a, b)
		} else {
			res = math.Pow(a, b)
		}
	case "fact":
		if a > 0 {
			res = factorial(a)
		} else {
			fmt.Printf("Результат операции невозможен, значение меньше нуля")
		}
		if a > 0 {
			resN = factorial(b)
		} else {
			fmt.Printf("Результат операции невозможен, значение меньше нуля")
		}
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
	if resN != 0 {
		fmt.Printf("Результат выполнения операции: %.2f\n", resN)
	}
	more := "y"
	fmt.Println("Ещё (Y/N)? ")
	fmt.Scanln(&more)
	if more == "y" || more == "Y" {
		fmt.Printf("\x1bc")
		calc()
	} else {
		return
	}

}

func factorial(n float64) float64 {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/*2. Задание для продвинутых (необязательное). Написать приложение, которое ищет все простые числа от 0 до N включительно. Число N должно быть задано из стандартного потока ввода.*/

func simpleDimple() []int {
	var n int
	var A = []int{1}
	fmt.Scanln(&n)
	if n < 1 {
		return nil
	}

	for i := 2; i <= n; i += 2 {
		m := int(math.Floor(math.Sqrt(float64(i))))
		for j := 2; j <= m; j++ {
			if i%j == 0 {
				continue
			}
		}
		A = append(A, i)
	}
	return A
}
