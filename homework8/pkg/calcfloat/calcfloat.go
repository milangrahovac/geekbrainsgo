package calcfloat

import (
	"math"
)

type CalcFLoat struct {
	X        float64
	Y        float64
	Operator string
	Result   float64
}

// float methods

func (n *CalcFLoat) Add() float64 {
	return n.X + n.Y
}

func (n *CalcFLoat) Subtract() float64 {
	return n.X - n.Y
}

func (n *CalcFLoat) Multiply() float64 {
	return n.X * n.Y
}

func (n *CalcFLoat) Divide() float64 {
	if n.Y == 0 {
		return math.Inf(int(n.X))
	} else {
		return n.X / n.Y
	}
}

func (n *CalcFLoat) Calc() float64 {
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
	return result
}
