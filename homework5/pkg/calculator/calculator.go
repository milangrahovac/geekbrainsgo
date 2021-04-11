package calculator

import (
	"fmt"
	"math"
)

// Numbers implements calculations for two given numbers
type Numbers struct {
	X float64
	Y float64
}

// Addoton calculates x+y
func (n *Numbers) Addition() float64 {
	return n.X + n.Y
}

// Subtraction calculates x-y

func (n *Numbers) Subtraction() float64 {
	return n.X - n.Y
}

// Multiplication calculates x*y
func (n *Numbers) Multiplication() float64 {
	return n.X * n.Y
}

// Division calculates x/y
func (n *Numbers) Division() float64 {
	return n.X / n.Y
}

// run math add, sub, mult, div function depends of operator and return result of operation float64
func (n *Numbers) Calc(op string) (float64, string) {
	var result float64
	switch op {
	case "+":
		result = n.Addition()
	case "-":
		result = n.Subtraction()
	case "*":
		result = n.Multiplication()
	case "/":
		if n.Y == 0 {
			result = math.Inf(int(n.X))
		} else {
			result = n.Division()
		}
	}
	str := fmt.Sprintf("%g %s %g = %g \n", n.X, op, n.Y, result)
	return result, str
}
