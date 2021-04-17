package calculator

import (
	"fmt"
	"math"
)

type Numbers struct {
	X        float64
	Y        float64
	Operator string
}

func (n *Numbers) Addition() float64 {
	return n.X + n.Y
}

func (n *Numbers) Subtraction() float64 {
	return n.X - n.Y
}

func (n *Numbers) Multiplication() float64 {
	return n.X * n.Y
}

func (n *Numbers) Division() float64 {
	if n.Y == 0 {
		return math.Inf(int(n.X))
	} else {
		return n.X / n.Y
	}
}

func (n *Numbers) Calc() (float64, string) {
	var result float64
	switch n.Operator {
	case "+":
		result = n.Addition()
	case "-":
		result = n.Subtraction()
	case "*":
		result = n.Multiplication()
	case "/":
		result = n.Division()
	}
	str := fmt.Sprintf("%g %s %g = %g \n", n.X, n.Operator, n.Y, result)
	return result, str
}
