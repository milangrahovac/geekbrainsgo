package calcuint

import (
	"fmt"
	"math"
)

type CalcUint struct {
	X        uint64
	Y        uint64
	Operator string
}

// Uint methods

func (n *CalcUint) Add() uint64 {
	return n.X + n.Y
}

func (n *CalcUint) Subtract() uint64 {
	return n.X - n.Y
}

func (n *CalcUint) Multiply() uint64 {
	return n.X * n.Y
}

func (n *CalcUint) Divide() uint64 {
	if n.Y == 0 {
		return uint64(math.Inf(int(n.X)))
	} else {
		return uint64(n.X / n.Y)
	}
}

func (u *CalcUint) Calc() (uint64, string) {
	var result uint64
	switch u.Operator {
	case "+":
		result = u.Add()
	case "-":
		result = u.Subtract()
	case "*":
		result = u.Multiply()
	case "/":
		result = u.Divide()
	}
	str := fmt.Sprintf("%v %s %v = %v \n", u.X, u.Operator, u.Y, result)
	return result, str
}
