package main

import "fmt"

func main() {
	var n1, n2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&n1)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&n2)

	fmt.Print("Введите знак операции (+ - / *): ")
	fmt.Scanln(&operator)

	switch operator {

	case string('+'):
		fmt.Printf("%f %s %f = %f", n1, operator, n2, n1+n2)

	case string('-'):
		fmt.Printf("%f %s %f = %f", n1, operator, n2, n1-n2)

	case string('*'):
		fmt.Printf("%f %s %f = %f", n1, operator, n2, n1*n2)

	case string('/'):
		if n2 == 0.0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Printf("%f %s %f = %f", n1, operator, n2, n1/n2)
		}

	default:
		fmt.Println("Неверный знак")
	}

}
