package calcint

import (
	"fmt"
	"math"
)

type CalcInt struct {
	X        int64
	Y        int64
	Operator string
}

// int methods

func (i *CalcInt) Add() int64 {
	return i.X + i.Y
}

func (i *CalcInt) Subtract() int64 {
	return i.X - i.Y
}

func (i *CalcInt) Multiply() int64 {
	return i.X * i.Y
}

func (i *CalcInt) Divide() int64 {
	if i.Y == 0 {
		return int64(math.Inf(int(i.X)))
	} else {
		return i.X / i.Y
	}
}

func (i *CalcInt) Calc() (int64, string) {
	var result int64
	switch i.Operator {
	case "+":
		result = i.Add()
	case "-":
		result = i.Subtract()
	case "*":
		result = i.Multiply()
	case "/":
		result = i.Divide()
	}
	str := fmt.Sprintf("%v %s %v = %v \n", i.X, i.Operator, i.Y, result)
	return result, str
}
