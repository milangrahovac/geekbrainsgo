package main

import (
	"fmt"
	"strconv"

	"github.com/milangrahovac/geekbrainsgo/homework6/pkg/calculator"
)

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
	fmt.Println("taks 1 - calculator")

	var v1 = getNumber("x")
	var op, _ = getOperator()
	var v2 = getNumber("y")

	n := calculator.Numbers{X: v1, Y: v2, Operator: op}
	_, resString := n.Calc()
	fmt.Println("result:", resString)

}
