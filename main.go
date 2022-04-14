package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//fmt.Println("S =", task1())
	//task2()
	task3()
}

/* 1. Напишите программу для вычисления площади прямоугольника. Длины сторон
прямоугольника должны вводиться пользователем с клавиатуры */
func task1() uint {
	var a, b uint

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	return a * b
}

/* 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.*/
func task2() {
	var A, D, C float64
	fmt.Scanln(&A)
	D = 2 * math.Sqrt(A/math.Pi)
	C = D * math.Pi
	fmt.Println("D=", D, "; C=", C)
}

/* 3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.*/
func task3() {
	d := 0
	fmt.Scanf("%d", &d)
	str := strconv.Itoa(d)
	fmt.Printf("%c\n", str[0]) //сотни
	fmt.Printf("%c\n", str[1]) //десятки
	fmt.Printf("%c\n", str[2]) //ед.
}
