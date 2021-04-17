package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/milangrahovac/geekbrainsgo/homework5/pkg/calculator"
	"github.com/milangrahovac/geekbrainsgo/homework5/pkg/fibonacci"
)

func f(i int64) int64 {
	switch i {
	case 0:
		return 0
	case 1:
		return 1
	}

	return f(i-1) + f(i-2)
}

func getPosiviveNumber() int64 {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter number between 0 and 30")

		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil || num < 0 || num > 30 {
				fmt.Println("Error: You must type a number between 0 and 30")
				continue
			}
			return num
		}
	}
}

// task 2
func getNumber(s string) float64 {
	var isNumber bool
	var n string
	for !isNumber {
		fmt.Printf("type number %s = ", s)
		fmt.Scan(&n)
		isNumber = isNumberic(n)
	}
	fn, _ := strconv.ParseFloat(n, 8)
	return fn
}

func isNumberic(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func getOperator() (string, error) {
	var op string
	var isOperator bool
	for !isOperator {
		fmt.Print("type operator +, -, * or /: ")
		fmt.Scan(&op)
		isOperator = operatorChecker(op)
	}
	return op, nil
}

func operatorChecker(s string) bool {
	switch s {
	case
		"+",
		"-",
		"*",
		"/":
		return true
	}
	return false
}

func main() {

	fmt.Println("taks 1 - fibonacci")
	// 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
	fmt.Println("Get n-th fibonacci number")

	number := getPosiviveNumber()
	fmt.Printf("%d - th fibonacci number is %d\n", number, f(number))
	fmt.Println("")

	// 2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
	fmt.Println("taks 2")
	fibNumber := fibonacci.Fib(number)
	fmt.Printf("%d - th fibonacci number is %d\n", number, fibNumber)
	fmt.Println("")

	// task 3 & 4
	// 3. Переписать калькутятор c использованием struct
	// 4. *Переписать калькутятор c использованием struct и сделать отдельные пакет
	fmt.Println("Task 3 & 4 - Calculator")

	var v1 = getNumber("x")
	var op, _ = getOperator()

	var v2 = getNumber("y")
	n := calculator.Numbers{X: v1, Y: v2}

	_, resString := n.Calc(op)
	fmt.Println("result:", resString)

}
