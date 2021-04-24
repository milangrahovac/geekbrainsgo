package main

import (
	"flag"
	"fmt"

	"github.com/milangrahovac/geekbrainsgo/homework8/pkg/calcfloat"
)

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

func printEquation(a float64, b float64, op string, r string) string {
	return fmt.Sprintf("%v %s %v = %s", a, op, b, r)
}

func flagCheker(m map[string]bool) bool {
	if _, ok := m["x"]; !ok {
		return false
	}
	if _, ok := m["y"]; !ok {
		return false
	}
	if _, ok := m["o"]; !ok {
		return false
	}
	return true
}

func main() {
	var (
		x          float64
		y          float64
		op         string
		floatRound string
		p          bool //print equation eg. 2 + 3 = 5
	)

	fmt.Println("taks 1 - calculator")

	flag.Float64Var(&x, "x", 0, "Float64 X")
	flag.Float64Var(&y, "y", 0, "Float64 Y")
	flag.StringVar(&op, "o", "", "Math Operator: add, substract, multiply or divide.")
	flag.StringVar(&floatRound, "d", "2", "Rounding it to x decimal places")
	flag.BoolVar(&p, "p", false, "Print equation")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })

	if !flagCheker(flagset) {
		fmt.Println("You must min add flags: x, y, and o.")
	} else {
		if !operatorChecker(op) {
			fmt.Println("Wrong operator! Put operator in quotes. You can use only +, -, * or /")
		} else {
			f := calcfloat.CalcFLoat{X: x, Y: y, Operator: op}
			result := f.Calc()
			s := "%." + floatRound + "f"

			roundedResult := fmt.Sprintf(s, result)
			fmt.Println("result:", roundedResult)

			if p {
				fmt.Print("equation: ")
				resultString := printEquation(x, y, op, roundedResult)
				fmt.Println(resultString)
			}
		}
	}
}

// ➜  homework8 go run homework8.go -x 234.998989 -y 56.789 -o / -d 8 -p
// taks 1 - calculator
// result: 4.13810754
// all equation: 234.998989 / 56.789 = 4.13810754
// ➜  homework8 go run homework8.go -x 234.998989 -y 56.789 -o / -d 2 -p
// taks 1 - calculator
// result: 4.14
// all equation: 234.998989 / 56.789 = 4.14
// ➜  homework8
