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

func (n *Numbers) Add() float64 {
	return n.X + n.Y
}

func (n *Numbers) Subtract() float64 {
	return n.X - n.Y
}

func (n *Numbers) Multiply() float64 {
	return n.X * n.Y
}

func (n *Numbers) Divide() float64 {
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
		result = n.Add()
	case "-":
		result = n.Subtract()
	case "*":
		result = n.Multiply()
	case "/":
		result = n.Divide()
	}
	str := fmt.Sprintf("%g %s %g = %g \n", n.X, n.Operator, n.Y, result)
	return result, str
}
